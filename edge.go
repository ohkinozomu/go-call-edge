package calledge

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

type CallEdge struct {
	Caller string
	Callee string
}

func GetCallEdges(f *ast.File) []CallEdge {
	callEdges := []CallEdge{}
	for _, d := range f.Decls {
		callEdges = append(callEdges, getCallEdges(d)...)
	}
	return callEdges
}

func getCallEdges(d ast.Decl) []CallEdge {
	callEdges := []CallEdge{}
	astutil.Apply(d, func(cr *astutil.Cursor) bool {
		if _, ok := cr.Node().(*ast.FuncDecl); ok {
			funcDecl := cr.Node().(*ast.FuncDecl)
			caller := funcDecl.Name.Name
			callExprs := getCallExprs(funcDecl)
			for _, callExpr := range callExprs {
				if _, ok := callExpr.Fun.(*ast.Ident); ok {
					callee := callExpr.Fun.(*ast.Ident).Name
					callEdge := CallEdge{
						Caller: caller,
						Callee: callee,
					}
					callEdges = append(callEdges, callEdge)
				}
			}

		}

		return true
	}, nil)
	return callEdges
}

func getCallExprs(d ast.Decl) []*ast.CallExpr {
	callExprs := []*ast.CallExpr{}
	astutil.Apply(d, func(cr *astutil.Cursor) bool {
		if _, ok := cr.Node().(*ast.CallExpr); ok {
			callExprs = append(callExprs, cr.Node().(*ast.CallExpr))
		}

		return true
	}, nil)
	return callExprs
}
