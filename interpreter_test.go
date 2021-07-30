package gojacego

import (
	"testing"
)

func TestBasicInterpreterSubstraction(test *testing.T) {
	interpreter := &interpreter{}

	ret, _ := interpreter.execute(newSubtractionOperation(integer,
		newConstantOperation(integer, 6),
		newConstantOperation(integer, 9)), nil, nil, nil)

	if ret != -3.0 {
		test.Errorf("Expected: -3.0, got: %f", ret)
	}
}

func TestBasicInterpreter1(test *testing.T) {
	interpreter := &interpreter{}

	// 6 + (2 * 4)

	ret, _ := interpreter.execute(
		newAddOperation(
			integer,
			newConstantOperation(integer, 6),
			newMultiplicationOperation(
				integer,
				newConstantOperation(integer, 2),
				newConstantOperation(integer, 4))), nil, nil, nil)

	if ret != 14.0 {
		test.Errorf("Expected: 14.0, got: %f", ret)
	}
}

func TestBasicInterpreterWithVariables(test *testing.T) {
	interpreter := &interpreter{}

	// var1 + 2 * (3 * age)

	parameters := make(map[string]interface{}, 4)
	parameters["var1"] = 2
	parameters["age"] = 4

	formulaVariables := createFormulaVariables(parameters, false)

	ret, _ := interpreter.execute(
		newAddOperation(floatingPoint,
			newVariableOperation("var1"),
			newMultiplicationOperation(
				floatingPoint,
				newConstantOperation(integer, 2),
				newMultiplicationOperation(
					floatingPoint,
					newConstantOperation(integer, 3),
					newVariableOperation("age")))), formulaVariables, nil, nil)

	if ret != 26.0 {
		test.Errorf("Expected: 26.0, got: %f", ret)
	}
}
