package goTaskScript

import (
	"github.com/leandroveronezi/go-terminal"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*
Task function representation
*/
type TaskFunction struct {
	Fun  interface{}
	Name string
}

var startGeral time.Time

/*
Call Task Functions
*/
func CallFunctions(functions map[int]TaskFunction, silent bool, continueOnErrors bool) {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)

	printScrenBegin(silent)

	for i := 0; i < len(functions); i++ {

		printTaskTitle(silent, functions, i)

		startTask := time.Now()

		goTerminal.ColorRGBForeground(255, 213, 79)
		goTerminal.Print("[" + startTask.Format("15:04:05") + "] ")

		goTerminal.ColorRGBForeground(79, 195, 247)
		//goTerminal.ColorRGBForeground(129, 199, 132)
		goTerminal.Print("Starting")

		goTerminal.CursorNextLine()

		callResult, err := callFunc(functions, i)

		endTask := time.Now()

		goTerminal.ColorRGBForeground(255, 213, 79)
		goTerminal.Print("[" + endTask.Format("15:04:05") + "] ")

		strTaskElapsed := time.Since(startTask).String()

		err, status, line := processCallFunction(silent, callResult, err)

		goTerminal.ColorRGBForeground(129, 199, 132)
		if err != nil {
			goTerminal.ColorRGBForeground(229, 115, 115)
		}

		goTerminal.Print("Finished after " + strTaskElapsed)
		goTerminal.CursorNextLine()

		goTerminal.Println("Return Line: " + line)
		goTerminal.Println("Return Status: " + status)

		if err != nil && !continueOnErrors {
			printScrenEnd(silent)
			os.Exit(1)
		}

	}

	printScrenEnd(silent)

}

func processCallFunction(Silent bool, CallResult []reflect.Value, CallErr error) (err error, status string, line string) {

	if CallErr != nil {

		err = CallErr
		status = "Error on call function"
		line = ""

		return err, status, line
	}

	if len(CallResult) == 0 {

		err = nil
		status = "Success"
		line = "non-return function"

		return err, status, line

	}

	for _, value := range CallResult {

		if foo, ok := value.Interface().(*GoTaskReturn); ok {

			status = "Success"

			if foo.Err != nil {
				status = "Error: " + foo.Err.Error()
			}

			line = strconv.Itoa(foo.Line)

			err = foo.Err

			return err, status, line

		}

	}

	return err, status, line

}

func printScrenBegin(Silent bool) {

	if Silent {
		return
	}

	startGeral = time.Now()

	goTerminal.ColorRGBForeground(255, 213, 79)
	goTerminal.Print("┤       ├" + strings.Repeat("─", 70))

	goTerminal.ColorRGBForeground(79, 195, 247)
	goTerminal.CursorColumn(3)
	goTerminal.Print("BEGIN")

	goTerminal.CursorNextLine()

}

func printTaskTitle(Silent bool, Functions map[int]TaskFunction, Idx int) {

	if Silent {
		return
	}

	strTotTask := strconv.Itoa(len(Functions))
	strTotTask = strings.Repeat("0", 4-len(strTotTask)) + strTotTask

	goTerminal.CursorNextLine()
	goTerminal.ColorRGBForeground(255, 213, 79)

	taskOf := strconv.Itoa(Idx + 1)
	taskOf = strings.Repeat("0", 4-len(taskOf)) + taskOf
	taskOf = taskOf + " of " + strTotTask

	goTerminal.Print("┤" + strings.Repeat(" ", 7+len(Functions[Idx].Name)) + " ├" + strings.Repeat("─", 55-len(Functions[Idx].Name)) + "┤" + strings.Repeat(" ", len(taskOf)) + "├")

	goTerminal.CursorColumn(3)
	goTerminal.ColorRGBForeground(79, 195, 247)
	goTerminal.Print("TASK: " + Functions[Idx].Name)

	goTerminal.CursorColumn(67)
	goTerminal.Print(taskOf)

	goTerminal.CursorNextLine()

}

func printScrenEnd(Silent bool) {

	if Silent {
		return
	}

	strElapsed := time.Since(startGeral).String()

	goTerminal.CursorNextLine()
	goTerminal.SetSGR(goTerminal.Reset)

	goTerminal.ColorRGBForeground(255, 213, 79)
	goTerminal.Print("┤     ├" + strings.Repeat("─", 53-len(strElapsed)) + "┤ " + strings.Repeat(" ", 15+len(strElapsed)) + " ├")

	goTerminal.ColorRGBForeground(79, 195, 247)
	goTerminal.CursorColumn(3)
	goTerminal.Print("END")

	goTerminal.CursorColumn(62 - len(strElapsed))
	goTerminal.Print("Duration Total: " + strElapsed)

	goTerminal.CursorNextLine()
	goTerminal.CursorNextLine()

	goTerminal.SetSGR(goTerminal.Reset)

}
