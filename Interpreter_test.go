package gojacego

import (
	"testing"
)

func TestBasicInterpreterSubstraction(test *testing.T) {
	interpreter := &Interpreter{}

	ret, _ := interpreter.Execute(NewSubtractionOperation(Integer,
		NewConstantOperation(Integer, 6),
		NewConstantOperation(Integer, 9)), nil)

	if ret != -3.0 {
		test.Errorf("Expected: -3.0, got: %f", ret)
	}
}
