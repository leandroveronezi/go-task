package bin

import (
	"bytes"
	"flag"
	"fmt"
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

func Main() {

	flagFile := flag.String("f", "", "File")
	silent := flag.Bool("silent", false, "Verbose mode")

	viewGenSource := flag.Bool("w", false, "Show generated source")
	keepFile := flag.Bool("k", false, "Keep generated file")
	sortFunctions := flag.Bool("s", false, "Sort orders of functions by name")

	continueOnErrors := flag.Bool("c", false, "Continue on errors")
	targetFunc := flag.String("t", "", "Target functions")

	flag.Parse()

	if len(os.Args) <= 1 {
		fmt.Println("File required")
		return
	}

	if *flagFile == "" {
		fmt.Println("File required")
		return
	}

	if *targetFunc != "" {

		temTarg := strings.Split(*targetFunc, ",")

		for _, fname := range temTarg {
			targetFuncList = append(targetFuncList, strings.ToUpper(fname))
		}

	}

	fset := token.NewFileSet() // positions are relative to fset
	node, err := parser.ParseFile(fset, *flagFile, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Walk(new(FuncVisitor), node)

	if *sortFunctions {
		sort.Strings(taskFunctions)
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, node)

	fmt.Fprintln(&buf, "\n")

	fmt.Fprintln(&buf, "//CODE AUTO GENERATED")

	fmt.Fprintln(&buf, "func main(){")

	fmt.Fprintln(&buf, "\n")

	vmode := "false"
	if *silent {
		vmode = "true"
	}

	cOnError := "false"
	if *continueOnErrors {
		cOnError = "true"
	}

	fmt.Fprintln(&buf, "	var taskFunctions = map[int]goTaskScript.RepFunctions{")

	for idx, name := range taskFunctions {

		fmt.Fprintln(&buf, "		"+strconv.Itoa(idx)+" : goTaskScript.RepFunctions{Name:"+strconv.Quote(name)+",Fun:"+name+"},")

	}

	fmt.Fprintln(&buf, "	}")
	fmt.Fprintln(&buf, "\n")

	fmt.Fprintln(&buf, "	goTaskScript.CallFunctions(taskFunctions,"+vmode+","+cOnError+")")
	fmt.Fprintln(&buf, "\n")

	fmt.Fprintln(&buf, "}")

	s := buf.String()

	if *viewGenSource {
		fmt.Println(s)
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

			fmt.Println("Function " + tar + " not found")

		}

		os.Exit(1)
	}

	dir, _ := filepath.Split(*flagFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	tmpFile, err := ioutil.TempFile(dir, "gotask.*.go")
	if err != nil {
		panic(err)
	}

	if !*keepFile {
		defer os.Remove(tmpFile.Name())
	}

	_, err = fmt.Fprint(tmpFile, s)

	if err != nil {
		panic(err)
	}

	tmpFile.Close()

	err = goRunFile(tmpFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}

}
