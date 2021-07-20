package gojacego

import (
	"reflect"
	"testing"
)

func TestBuildFormula1(test *testing.T) {
	astBuilder := NewAstBuilder(false)
	params := []Token{
		{Value: '(', Type: LEFT_BRACKET},
		{Value: 42, Type: INTEGER},
		{Value: '+', Type: OPERATION},
		{Value: 8, Type: INTEGER},
		{Value: ')', Type: RIGHT_BRACKET},
		{Value: '*', Type: OPERATION},
		{Value: 2, Type: INTEGER},
	}
	op, _ := astBuilder.Build(params)

	if reflect.TypeOf(op).String() != "*gojacego.MultiplicationOperation" {
		test.Errorf("expected: MultiplicationOperation, got: %s", reflect.TypeOf(op).String())
	}

	multiplication := op.(*MultiplicationOperation)
	addition := multiplication.OperationOne.(*AddOperation)
	add_one := addition.OperationOne.(*ConstantOperation).Value
	if add_one != 42 {
		test.Errorf("exptected: 42, got: %d", add_one)
	}

	add_two := addition.OperationTwo.(*ConstantOperation).Value
	if add_two != 8 {
		test.Errorf("exptected: 8, got: %d", add_one)
	}

	multi_two := multiplication.OperationTwo.(*ConstantOperation).Value
	if multi_two != 2 {
		test.Errorf("exptected: 2, got: %d", add_one)
	}
}

func TestBuildFormula2(test *testing.T) {
	astBuilder := NewAstBuilder(false)
	params := []Token{
		{Value: 2, Type: INTEGER},
		{Value: '+', Type: OPERATION},
		{Value: 8, Type: INTEGER},
		{Value: '*', Type: OPERATION},
		{Value: 3, Type: INTEGER},
	}
	op, _ := astBuilder.Build(params)

	if reflect.TypeOf(op).String() != "*gojacego.AddOperation" {
		test.Errorf("expected: AddOperation, got: %s", reflect.TypeOf(op).String())
	}

	addition := op.(*AddOperation)
	multiplication := addition.OperationTwo.(*MultiplicationOperation)

	add_one := addition.OperationOne.(*ConstantOperation).Value
	if add_one != 2 {
		test.Errorf("exptected: 2, got: %d", add_one)
	}

	multi_one := multiplication.OperationOne.(*ConstantOperation).Value
	if multi_one != 8 {
		test.Errorf("exptected: 8, got: %d", multi_one)
	}

	multi_two := multiplication.OperationTwo.(*ConstantOperation).Value
	if multi_two != 3 {
		test.Errorf("exptected: 3, got: %d", add_one)
	}

}
