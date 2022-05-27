package calledge

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestGetCallEdges(t *testing.T) {
	expect := []CallEdge{
		CallEdge{
			Caller: "Concat",
			Callee: "make",
		},
		CallEdge{
			Caller: "Concat",
			Callee: "len",
		},
		CallEdge{
			Caller: "Concat",
			Callee: "append",
		},
		CallEdge{
			Caller: "DynamicRows",
			Callee: "newDynamicRowGroupReader",
		}}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tests/concat.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	callEdges := GetCallEdges(f)
	if !reflect.DeepEqual(callEdges, expect) {
		t.Fatalf("test fails: %v", callEdges)
	}
}
