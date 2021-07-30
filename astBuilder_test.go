package gojacego

import (
	"reflect"
	"testing"
)

func getConstantRegistry() *constantRegistry {
	return &constantRegistry{
		caseSensitive: false,
	}
}

func getFunctionRegistry() *functionRegistry {
	return &functionRegistry{
		caseSensitive: false,
	}
}

func TestBuildFormula1(test *testing.T) {
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)
	params := []token{
		{Value: '(', Type: tt_LEFT_BRACKET},
		{Value: 42, Type: tt_INTEGER},
		{Value: '+', Type: tt_OPERATION},
		{Value: 8, Type: tt_INTEGER},
		{Value: ')', Type: tt_RIGHT_BRACKET},
		{Value: '*', Type: tt_OPERATION},
		{Value: 2, Type: tt_INTEGER},
	}
	op, _ := astBuilder.build(params)

	if reflect.TypeOf(op).String() != "*gojacego.multiplicationOperation" {
		test.Errorf("expected: multiplicationOperation, got: %s", reflect.TypeOf(op).String())
	}

	multiplication := op.(*multiplicationOperation)
	addition := multiplication.OperationOne.(*addOperation)
	add_one := addition.OperationOne.(*constantOperation).Value
	if add_one != 42 {
		test.Errorf("exptected: 42, got: %d", add_one)
	}

	add_two := addition.OperationTwo.(*constantOperation).Value
	if add_two != 8 {
		test.Errorf("exptected: 8, got: %d", add_one)
	}

	multi_two := multiplication.OperationTwo.(*constantOperation).Value
	if multi_two != 2 {
		test.Errorf("exptected: 2, got: %d", add_one)
	}
}

func TestBuildFormula2(test *testing.T) {
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)
	params := []token{
		{Value: 2, Type: tt_INTEGER},
		{Value: '+', Type: tt_OPERATION},
		{Value: 8, Type: tt_INTEGER},
		{Value: '*', Type: tt_OPERATION},
		{Value: 3, Type: tt_INTEGER},
	}
	op, _ := astBuilder.build(params)

	if reflect.TypeOf(op).String() != "*gojacego.addOperation" {
		test.Errorf("expected: AddOperation, got: %s", reflect.TypeOf(op).String())
	}

	addition := op.(*addOperation)
	multiplication := addition.OperationTwo.(*multiplicationOperation)

	add_one := addition.OperationOne.(*constantOperation).Value
	if add_one != 2 {
		test.Errorf("exptected: 2, got: %d", add_one)
	}

	multi_one := multiplication.OperationOne.(*constantOperation).Value
	if multi_one != 8 {
		test.Errorf("exptected: 8, got: %d", multi_one)
	}

	multi_two := multiplication.OperationTwo.(*constantOperation).Value
	if multi_two != 3 {
		test.Errorf("exptected: 3, got: %d", add_one)
	}

}

func TestUnaryMinus(test *testing.T) {
	astBuilder := newAstBuilder(false, getFunctionRegistry(), getConstantRegistry(), nil)
	params := []token{
		{Value: 5.3, Type: tt_FLOATING_POINT},
		{Value: '*', Type: tt_OPERATION},
		{Value: '_', Type: tt_OPERATION},
		{Value: '(', Type: tt_LEFT_BRACKET},
		{Value: 5, Type: tt_INTEGER},
		{Value: '+', Type: tt_OPERATION},
		{Value: 42, Type: tt_INTEGER},
		{Value: ')', Type: tt_RIGHT_BRACKET},
	}

	op, _ := astBuilder.build(params)

	if reflect.TypeOf(op).String() != "*gojacego.multiplicationOperation" {
		test.Errorf("expected: multiplicationOperation, got: %s", reflect.TypeOf(op).String())
	}

	multiplication := op.(*multiplicationOperation)

	multi_one := multiplication.OperationOne.(*constantOperation).Value
	if multi_one != 5.3 {
		test.Errorf("exptected: 5.3, got: %d", multi_one)
	}

	unaryMinus := multiplication.OperationTwo.(*unaryMinusOperation)

	addition := unaryMinus.Operation.(*addOperation)

	add_one := addition.OperationOne.(*constantOperation).Value
	if add_one != 5 {
		test.Errorf("exptected: 5, got: %d", add_one)
	}

	add_two := addition.OperationTwo.(*constantOperation).Value
	if add_two != 42 {
		test.Errorf("exptected: 42, got: %d", add_one)
	}

}
