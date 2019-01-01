package bin

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/leandroveronezi/go-terminal"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// go-task
func Main() {

	flagFileName = flag.String("f", "", "File")
	flagSilent = flag.Bool("silent", false, "Silent mode")
	flagKeepFile = flag.Bool("k", false, "Keep generated file")
	flagViewGenSource = flag.Bool("w", false, "View generated source")
	flagSortFunctions = flag.Bool("s", false, "Sort orders of functions by name before run")
	flagContinueOnErrors = flag.Bool("c", false, "Skip errors and continue")
	flagTargetFunc = flag.String("t", "", "Target functions")
	flagGroupFunc = flag.String("g", "", "Run function by group")

	flag.Parse()

	if !processFlags() {
		return
	}

	if err := createAstFile(); err != nil {
		printError(err)
		return
	}

	if len(targetFuncList) > 0 && len(targetFuncList) != len(taskFunctions) {

	FORMAIN:
		for _, tar := range targetFuncList {

			for _, tas := range taskFunctions {

				if tar == strings.ToUpper(tas) {
					continue FORMAIN
				}

			}

			printError("Function " + tar + " not found")

		}

		os.Exit(1)
	}

	if *flagSortFunctions {
		sort.Strings(taskFunctions)
	}

	if err := generateFileAndRun(); err != nil {
		printError(err)
		return
	}

}

func processFlags() bool {

	if *flagFileName == "" {
		printError("File required")
		return false
	}

	if !fileExists(*flagFileName) {
		printError("File " + *flagFileName + " not exist")
		return false
	}

	*flagGroupFunc = strings.ToUpper(strings.Trim(*flagGroupFunc, " "))

	if *flagGroupFunc != "" {
		return true
	}

	if *flagTargetFunc == "" {
		return true
	}

	temTarget := strings.Split(*flagTargetFunc, ",")

	for _, fname := range temTarget {
		targetFuncList = append(targetFuncList, strings.ToUpper(fname))
	}

	return true

}

func createAstFile() error {

	var err error

	fset = token.NewFileSet() // positions are relative to fset
	nodeFile, err = parser.ParseFile(fset, *flagFileName, nil, parser.ParseComments)

	if err != nil {
		return err
	}

	ast.Walk(new(FuncVisitor), nodeFile)

	return nil

}

func generateFileAndRun() error {

	vmode := "false"
	if *flagSilent {
		vmode = "true"
	}

	cOnError := "false"
	if *flagContinueOnErrors {
		cOnError = "true"
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, nodeFile)

	fmt.Fprintln(&buf, "//CODE AUTO GENERATED")
	fmt.Fprintln(&buf, "func main(){")
	fmt.Fprintln(&buf, "\t var taskFunctions = map[int]goTaskScript.TaskFunction{")

	for idx, name := range taskFunctions {
		fmt.Fprintln(&buf, "\t\t "+strconv.Itoa(idx)+" : goTaskScript.TaskFunction{Name:"+strconv.Quote(name)+",Fun:"+name+"},")
	}

	fmt.Fprintln(&buf, "\t }")
	fmt.Fprintln(&buf, "\t goTaskScript.CallFunctions(taskFunctions,"+vmode+","+cOnError+")")
	fmt.Fprintln(&buf, "}")

	s := buf.String()

	if *flagViewGenSource {
		fmt.Println(s)
		return nil
	}

	dir, _ := filepath.Split(*flagFileName)

	tempFile, err := ioutil.TempFile(dir, "gotask.*.go")

	if err != nil {
		return err
	}

	if !*flagKeepFile {
		defer os.Remove(tempFile.Name())
	}

	_, err = fmt.Fprint(tempFile, s)

	if err != nil {
		return err
	}

	tempFile.Close()

	return goRunFile(tempFile.Name())

}

func printError(Err interface{}) {

	goTerminal.ColorRGBForeground(229, 115, 115)
	fmt.Println(Err)
	goTerminal.SetSGR(goTerminal.Reset)

}
