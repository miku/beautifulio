package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"io"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var x io.Writer

var Analyzer = &analysis.Analyzer{
	Name: "addlint",
	Doc:  "reports integer additions",
	Run:  run,
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

// func run(pass *analysis.Pass) (interface{}, error) {
// 	return nil, errors.New("not implemented yet")
// }

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// check whether the call expression matches time.Now().Sub()
			be, ok := n.(*ast.BinaryExpr)
			if !ok {
				return true
			}

			if be.Op != token.ADD {
				return true
			}

			if _, ok := be.X.(*ast.BasicLit); !ok {
				return true
			}

			if _, ok := be.Y.(*ast.BasicLit); !ok {
				return true
			}

			isInteger := func(expr ast.Expr) bool {
				t := pass.TypesInfo.TypeOf(expr)
				if t == nil {
					return false
				}

				bt, ok := t.Underlying().(*types.Basic)
				if !ok {
					return false
				}

				if (bt.Info() & types.IsInteger) == 0 {
					return false
				}

				return true
			}

			// check that both left and right hand side are integers
			if !isInteger(be.X) || !isInteger(be.Y) {
				return true
			}

			pass.Reportf(be.Pos(), "integer addition found %q",
				render(pass.Fset, be))
			return true
		})
	}

	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
