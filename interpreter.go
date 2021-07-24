package gojacego

import (
	"errors"
	"fmt"
	"math"
)

type Interpreter struct {
	caseSensitive bool
}

type Formula func(vars map[string]interface{}) float64

func (*Interpreter) Execute(op Operation, vars map[string]interface{}) (float64, error) {
	return execute(op, vars)
}

func (*Interpreter) BuildFormula(op Operation) Formula {
	return func(vars map[string]interface{}) float64 {
		ret, _ := execute(op, vars)
		return ret
	}
}

func execute(op Operation, vars FormulaVariables) (float64, error) {

	if op == nil {
		return 0, errors.New("operation cannot be nil")
	}

	if cop, ok := op.(*ConstantOperation); ok {
		if cop.Metadata.DataType == Integer {
			return float64(cop.Value.(int)), nil
		} else {
			return cop.Value.(float64), nil
		}

	} else if cop, ok := op.(*VariableOperation); ok {

		variableValue, err := vars.Get(cop.Name)
		if err == nil {
			return variableValue.(float64), nil
		} else {
			return 0, errors.New("The variable '" + cop.Name + "' used is not defined.")
		}

	} else if cop, ok := op.(*MultiplicationOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		return left * right, nil
	} else if cop, ok := op.(*AddOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		return left + right, nil
	} else if cop, ok := op.(*SubtractionOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		return left - right, nil
	} else if cop, ok := op.(*DivisorOperation); ok {
		left, _ := execute(cop.Dividend, vars)
		right, _ := execute(cop.Divisor, vars)

		return left / right, nil
	} else if cop, ok := op.(*ModuloOperation); ok {
		left, _ := execute(cop.Dividend, vars)
		right, _ := execute(cop.Divisor, vars)

		return math.Mod(left, right), nil
	} else if cop, ok := op.(*ExponentiationOperation); ok {
		left, _ := execute(cop.Base, vars)
		right, _ := execute(cop.Exponent, vars)

		return math.Pow(left, right), nil
	} else if cop, ok := op.(*UnaryMinusOperation); ok {
		arg, _ := execute(cop.Operation, vars)
		return -arg, nil
	} else if cop, ok := op.(*AndOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left != 0 && right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*OrOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left != 0 || right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*LessThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left < right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*LessOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left <= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*GreaterThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left > right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*GreaterOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left >= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*EqualOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left == right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*NotEqualOperation); ok {
		left, _ := execute(cop.OperationOne, vars)
		right, _ := execute(cop.OperationTwo, vars)

		if left != right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*FunctionOperation); ok {

		return 0.0, fmt.Errorf("Not implemented %T", cop)
	}

	return 0.0, fmt.Errorf("Not implemented %T", op)
}
