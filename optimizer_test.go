package gojacego

import (
	"reflect"
	"testing"
)

func TestOptimizerMultiplicationByZero(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("var1 * 0.0")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 0.0 {
		test.Errorf("Expected: 0.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}
