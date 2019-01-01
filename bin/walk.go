package bin

import (
	"go/ast"
	"go/token"
	"os"
	"strconv"
	"strings"
)

type FuncVisitor struct {
}

// ast Walk Visit for node
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
					printError("Function main is not allowed")
					os.Exit(1)
				}

				if fn.Recv == nil && fn.Name.IsExported() {

					if *flagGroupFunc != "" {

						auxDoc := strings.ToUpper(fn.Doc.Text())
						auxDoc = strings.Trim(auxDoc, " ")
						auxDoc = strings.TrimSuffix(auxDoc, "\n")

						if strings.Contains(auxDoc, "GROUP:") {

							idx := strings.Index(auxDoc, "GROUP:")
							auxDoc = auxDoc[idx+6:]

							docGroup := strings.Split(auxDoc, ",")

							for _, fname := range docGroup {

								if fname == *flagGroupFunc {
									taskFunctions = append(taskFunctions, fn.Name.String())
									break
								}
							}

						}

					} else if len(targetFuncList) > 0 {

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
							printError("Params not allowed in exported functions")
							os.Exit(1)
						}

					}

					if fn.Type.Results != nil {

						if len(fn.Type.Results.List) > 1 {
							printError("Only one result in exported functions")
							os.Exit(1)
						}

						res, ok := fn.Type.Results.List[0].Type.(*ast.Ident)
						if ok {

							if res.Name != "error" && res.Name != "bool" {
								printError("Only error and boolean result in exported functions")
								os.Exit(1)
							}

							tp := KError
							if res.Name == "bool" {
								tp = KBool
							}

							res.Name = "*goTaskScript.GoTaskReturn"

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
