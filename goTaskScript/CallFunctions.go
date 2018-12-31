package goTaskScript

import (
	"github.com/leandroveronezi/go-terminal"
	"os"
	"strconv"
	"strings"
)

type RepFunctions struct {
	Fun  interface{}
	Name string
}

func CallFunctions(functions map[int]RepFunctions, silent bool, continueOnErrors bool) {

	goTerminal.Clean()
	goTerminal.CursorLineColumn(1, 1)

	totColumn := 80

	if !silent {
		goTerminal.SetSGR(goTerminal.ForegroundLightGray)
		goTerminal.Println(strings.Repeat("┈", totColumn))
	}

	for i := 0; i < len(functions); i++ {

		if !silent {
			goTerminal.SetSGR(goTerminal.ForegroundLightBlue)
			goTerminal.CursorNextLine()
			goTerminal.Println("Function: " + functions[i].Name)
		}

		callResult, err := callFunc(functions, i)

		if err != nil {

			if !silent {
				goTerminal.SetSGR(goTerminal.ForegroundLightRed)
				goTerminal.Println("Error on call function " + functions[i].Name)
			}

			goTerminal.Println(err.Error())
			os.Exit(1)
		}

		if len(callResult) == 0 {

			if !silent {
				goTerminal.SetSGR(goTerminal.ForegroundLightGreen)
				goTerminal.Println("Status: Success")
				goTerminal.Println("Line: non-return function")
			}

		} else {

			for _, value := range callResult {

				if foo, ok := value.Interface().(*GoCallError); ok {

					if foo.Err != nil {

						if !silent {
							goTerminal.SetSGR(goTerminal.ForegroundLightRed)
							goTerminal.Println("Status: Error")
						}

						goTerminal.Println(foo.Err.Error())

					} else {

						if !silent {
							goTerminal.SetSGR(goTerminal.ForegroundLightGreen)
							goTerminal.Println("Status: Success")
						}

					}

					if !silent {
						goTerminal.Println("Line: " + strconv.Itoa(foo.Line))
					}

					if foo.Err != nil {
						if !continueOnErrors {
							goTerminal.SetSGR(goTerminal.Reset)
							os.Exit(1)
						}
					}

				}

			}

		}

	}

	goTerminal.SetSGR(goTerminal.ForegroundLightGray)
	goTerminal.Println(strings.Repeat("┈", totColumn))

	goTerminal.SetSGR(goTerminal.Reset)

}
