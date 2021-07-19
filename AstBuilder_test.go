package gojacego

import (
	"reflect"
	"testing"
)

func TestBuildFormula1(test *testing.T) {
	astBuilder := NewAstBuilder(false)
	params := []Token{
		{Value: 42, Type: INTEGER},
		{Value: '+', Type: OPERATION},
		{Value: 8, Type: INTEGER},
	}
	op, _ := astBuilder.Build(params)

	if reflect.TypeOf(op).Name() != "MultiplicationOperation" {
		test.Errorf("expected: MultiplicationOperation, got: %s", reflect.TypeOf(op).Name())
	}
}

func TestBuildFormula2(test *testing.T) {
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

	if reflect.TypeOf(op).Name() != "MultiplicationOperation" {
		test.Errorf("expected: MultiplicationOperation, got: %s", reflect.TypeOf(op).Name())
	}
}
