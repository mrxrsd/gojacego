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

func TestBasicInterpreter1(test *testing.T) {
	interpreter := &Interpreter{}

	// 6 + (2 * 4)

	ret, _ := interpreter.Execute(
		NewAddOperation(
			Integer,
			NewConstantOperation(Integer, 6),
			NewMultiplicationOperation(
				Integer,
				NewConstantOperation(Integer, 2),
				NewConstantOperation(Integer, 4))), nil)

	if ret != 14.0 {
		test.Errorf("Expected: 14.0, got: %f", ret)
	}
}

func TestBasicInterpreterWithVariables(test *testing.T) {
	interpreter := &Interpreter{}

	// var1 + 2 * (3 * age)

	parameters := make(map[string]interface{}, 4)
	parameters["var1"] = 2.0
	parameters["age"] = 4.0

	ret, _ := interpreter.Execute(
		NewAddOperation(FloatingPoint,
			NewVariableOperation("var1"),
			NewMultiplicationOperation(
				FloatingPoint,
				NewConstantOperation(Integer, 2),
				NewMultiplicationOperation(
					FloatingPoint,
					NewConstantOperation(Integer, 3),
					NewVariableOperation("age")))), parameters)

	if ret != 26.0 {
		test.Errorf("Expected: 26.0, got: %f", ret)
	}
}
