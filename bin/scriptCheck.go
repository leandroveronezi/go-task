package bin

import (
	"go/ast"
	"os"
)

func fileExists(FileName string) bool {
	file, err := os.Stat(FileName)
	return (err == nil) && !file.IsDir()
}

func expToStr(Expr ast.Expr) string {

	lito, ok := Expr.(*ast.BasicLit)

	if ok {
		return lito.Value
	}

	dente, ok := Expr.(*ast.Ident)

	if ok {
		return dente.Name
	}

	return "to-do"

}

func checkFunctionReturn(tp KindType, List ...ast.Stmt) {

	for _, aa := range List {

		ccase, ok := aa.(*ast.CaseClause)
		if ok {
			checkFunctionReturn(tp, ccase.Body...)
			continue
		}

		block, ok := aa.(*ast.BlockStmt)
		if ok {
			checkFunctionReturn(tp, block.List...)
			continue
		}

		forst, ok := aa.(*ast.ForStmt)
		if ok {
			checkFunctionReturn(tp, forst.Body.List...)
			continue
		}

		ifst, ok := aa.(*ast.IfStmt)
		if ok {
			checkFunctionReturn(tp, ifst.Body.List...)
			continue
		}

		ret, ok := aa.(*ast.ReturnStmt)
		if ok {

			for idx, ll := range ret.Results {

				call, ok := ll.(*ast.CallExpr)

				if ok {

					var newcall ast.CallExpr
					var oldcall ast.CallExpr

					oldcall = *call

					var Args []ast.Expr
					Args = make([]ast.Expr, 0)
					Args = append(Args, &oldcall)

					fname := "ProcessReturn"
					if tp == KBool {
						fname = "ProcessBool"
					}

					newcall.Fun = &ast.SelectorExpr{

						X: &ast.Ident{
							Name: "goTaskScript",
						},

						Sel: &ast.Ident{
							Name: fname,
						},
					}

					call.Fun = newcall.Fun
					call.Args = Args

				}

				dente, ok := ll.(*ast.Ident)

				if ok {

					fname := "ProcessReturn"
					if tp == KBool {
						fname = "ProcessBool"
					}

					dente.Name = "goTaskScript." + fname + "(" + dente.Name + ")"

				}

				bin, ok := ll.(*ast.BinaryExpr)

				if ok {

					var lito ast.BasicLit

					lito.Value = "goTaskScript.ProcessBool(" + expToStr(bin.X) + " " + bin.Op.String() + " " + expToStr(bin.Y) + ")"

					ret.Results[idx] = &lito

				}

			}

		}
	}

}
