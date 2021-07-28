package gojacego

import (
	"reflect"
	"testing"
)

func TestOptimizerMultiplicationByZero(test *testing.T) {
	interpreter := &Interpreter{}
	optimizer := &Optimizer{executor: *interpreter}
	reader := NewTokenReader('.', ',')
	astBuilder := NewAstBuilder(false, getFunctionRegistry(), getConstantRegistry())

	tokens, _ := reader.Read("var1 * 0.0")
	operation, _ := astBuilder.Build(tokens)
	optimizedOperation := optimizer.Optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.ConstantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*ConstantOperation).Value != 0.0 {
		test.Errorf("Expected: 0.0, got: %f", optimizedOperation.(*ConstantOperation).Value)
	}
}
