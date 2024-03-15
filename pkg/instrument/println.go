package instrument

import (
	"go/ast"
	"go/token"
)

// AddPrintln 函数开始和结束添加打印
// fmt.Println("func %fnName begin")
// fmt.Println("func %fnName end")
func AddPrintln(fn *ast.FuncDecl) {
	fnName := fn.Name.Name
	begin := "func " + fnName + " begin"
	beginPrintStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{Name: "fmt.Println"},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: "\"" + begin + "\""},
			},
		},
	}

	end := "func " + fnName + " end"
	endPrintStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{Name: "fmt.Println"},
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.STRING, Value: "\"" + end + "\""},
			},
		},
	}

	fn.Body.List = append([]ast.Stmt{beginPrintStmt}, fn.Body.List...)
	fn.Body.List = append(fn.Body.List, endPrintStmt)
}
