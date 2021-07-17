package gojacego

import (
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
	_, err := astBuilder.Build(params)

	if !errorContains(err, "formula cannot be empty") {
		test.Errorf("unexpected error: %v", err)
	}
}
