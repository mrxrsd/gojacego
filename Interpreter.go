package gojacego

import (
	"errors"
)

type Interpreter struct {
	caseSensitive bool
}

func (*Interpreter) Execute(op Operation, vars FormulaVariables) (float64, error) {
	return execute(op, vars)
}

func execute(op Operation, vars FormulaVariables) (float64, error) {

	if op == nil {
		return 0, errors.New("Operation cannot be nil.")
	}

	if cop, ok := op.(*ConstantOperation); ok {
		if cop.Metadata.DataType == Integer {
			return float64(cop.Value.(int)), nil
		} else {
			return cop.Value.(float64), nil
		}

	} else if cop, ok := op.(*VariableOperation); ok {

		variableValue, err := vars.Get(cop.Name)
		if err != nil {
			return variableValue.(float64), nil
		} else {
			return 0, errors.New("The variable '" + cop.Name + "' used is not defined.")
		}

	} else if cop, ok := op.(*MultiplicationOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		return left * right, nil
	} else if cop, ok := op.(*SubtractionOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		return left - right, nil
	}

	return 1.0, nil
}
