package bin

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strconv"
	"strings"
)

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {

	switch t := node.(type) {

	case *ast.File:
		{
			// Change the package name
			t.Name.Name = "main"

		}

	case *ast.FuncDecl:
		{
			fn, ok := node.(*ast.FuncDecl)
			if ok {

				if fn.Name.String() == "main" {
					fmt.Println("Function main is not allowed")
					os.Exit(1)
				}

				if fn.Recv == nil && fn.Name.IsExported() {

					if len(targetFuncList) > 0 {

						for _, v := range targetFuncList {
							if v == strings.ToUpper(fn.Name.String()) {
								taskFunctions = append(taskFunctions, fn.Name.String())
							}
						}

					} else {

						taskFunctions = append(taskFunctions, fn.Name.String())

					}

					if fn.Type.Params != nil {

						if len(fn.Type.Params.List) > 0 {
							fmt.Println("Params not allowed in exported functions")
							os.Exit(1)
						}

					}

					if fn.Type.Results != nil {

						if len(fn.Type.Results.List) > 1 {
							fmt.Println("Only one result in exported functions")
							os.Exit(1)
						}

						res, ok := fn.Type.Results.List[0].Type.(*ast.Ident)
						if ok {

							if res.Name != "error" && res.Name != "bool" {
								fmt.Println("Only error and bollean result in exported functions")
								os.Exit(1)
							}

							tp := KError
							if res.Name == "bool" {
								tp = KBool
							}

							res.Name = "*goTaskScript.GoCallError"

							checkFunctionReturn(tp, fn.Body.List...)

						}

					}

				}

			}

		}

	case *ast.GenDecl:
		{
			if t.Tok == token.IMPORT {
				iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("github.com/leandroveronezi/go-task/goTaskScript")}}
				t.Specs = append(t.Specs, iSpec)
			}
		}

	}

	return v
}
