package gojacego

import (
	"reflect"
	"testing"
)

func TestOptimizerIdempotentFunction(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')

	fnRegistry := getFunctionRegistry()
	fnRegistry.registerFunction("test", func(arguments ...interface{}) float64 {
		return arguments[0].(float64) + arguments[1].(float64)
	}, false, true)

	astBuilder := newAstBuilder(false, fnRegistry, getConstantRegistry(), nil)

	tokens, _ := reader.read("test(var1, (2+3) * 500)")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, fnRegistry, getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.functionOperation" {
		test.Errorf("expected: functionOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	fnOperation := optimizedOperation.(*functionOperation)
	fnArgument := fnOperation.Arguments[1]

	if reflect.TypeOf(fnArgument).String() != "*gojacego.constantOperation" {
		test.Errorf("Expected: *gojacego.constantOperation, got: %s", reflect.TypeOf(fnArgument).String())
	}
}

func TestOptimizerNonIdempotentFunction(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')

	fnRegistry := getFunctionRegistry()
	fnRegistry.registerFunction("test", func(arguments ...interface{}) float64 {
		return arguments[0].(float64)
	}, false, false)

	astBuilder := newAstBuilder(false, fnRegistry, getConstantRegistry(), nil)

	tokens, _ := reader.read("test(500)")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, fnRegistry, getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.functionOperation" {
		test.Errorf("expected: functionOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	fnOperation := optimizedOperation.(*functionOperation)
	fnArgument := fnOperation.Arguments[0]

	if reflect.TypeOf(fnArgument).String() != "*gojacego.constantOperation" {
		test.Errorf("Expected: *gojacego.constantOperation, got: %s", reflect.TypeOf(fnArgument).String())
	}
}

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

func TestBasicOptimizer(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("2 + 2")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 4.0 {
		test.Errorf("Expected: 4.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestBooleanOperationOptimizer(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("4 > 2")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 1.0 {
		test.Errorf("Expected: 1.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestBooleanOperation2Optimizer(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("(4 > 2) && var1")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.andOperation" {
		test.Errorf("expected: andOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}
	andOperation := optimizedOperation.(*andOperation)

	if andOperation.OperationOne.(*constantOperation).Value != 1.0 {
		test.Errorf("Expected: 1.0, got: %f", andOperation.OperationOne.(*constantOperation).Value)
	}
}

func TestShortCircuitAndOptimizer(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("0 && var_x")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 0.0 {
		test.Errorf("Expected: 0.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestShortCircuitAndOptimizer1(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("var_x && 0")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 0.0 {
		test.Errorf("Expected: 0.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestShortCircuitAndOptimizer3(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("var_x && 1")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.andOperation" {
		test.Errorf("expected: AndOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}
}

func TestShortCircuitOrOptimizer(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("1 || var_x")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 1.0 {
		test.Errorf("Expected: 1.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestShortCircuitOrOptimizer2(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("var_x || 1")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.constantOperation" {
		test.Errorf("expected: ConstantOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}

	if optimizedOperation.(*constantOperation).Value != 1.0 {
		test.Errorf("Expected: 1.0, got: %f", optimizedOperation.(*constantOperation).Value)
	}
}

func TestShortCircuitOrOptimizer3(test *testing.T) {
	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	reader := newTokenReader('.', ',')
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)

	tokens, _ := reader.read("var_x || 0")
	operation, _ := astBuilder.build(tokens)
	optimizedOperation := optimizer.optimize(operation, getFunctionRegistry(), getConstantRegistry())

	if reflect.TypeOf(optimizedOperation).String() != "*gojacego.orOperation" {
		test.Errorf("expected: OrOperation, got: %s", reflect.TypeOf(optimizedOperation).String())
	}
}
