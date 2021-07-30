package gojacego

import (
	"errors"
	"fmt"
	"math"
)

type interpreter struct {
	caseSensitive bool
}

type Formula func(vars formulaVariables) float64

func (*interpreter) execute(op operation, vars formulaVariables, functionRegistry *functionRegistry, constantRegistry *constantRegistry) (float64, error) {
	return execute(op, vars, functionRegistry, constantRegistry)
}

func (*interpreter) buildFormula(op operation, functionRegistry *functionRegistry, constantRegistry *constantRegistry) Formula {
	return func(vars formulaVariables) float64 {
		ret, _ := execute(op, vars, functionRegistry, constantRegistry)
		return ret
	}
}

func execute(op operation, vars formulaVariables, functionRegistry *functionRegistry, constantRegistry *constantRegistry) (float64, error) {

	if op == nil {
		return 0, errors.New("operation cannot be nil")
	}

	if cop, ok := op.(*constantOperation); ok {
		if cop.Metadata.DataType == integer {
			return toFloat64(cop.Value), nil
		} else {
			return cop.Value.(float64), nil
		}

	} else if cop, ok := op.(*variableOperation); ok {

		variableValue, err := vars.Get(cop.Name)
		if err == nil {
			return variableValue, nil
		} else {
			return 0, errors.New("The variable '" + cop.Name + "' used is not defined.")
		}

	} else if cop, ok := op.(*multiplicationOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left * right, nil
	} else if cop, ok := op.(*addOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left + right, nil
	} else if cop, ok := op.(*subtractionOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left - right, nil
	} else if cop, ok := op.(*divisorOperation); ok {
		left, _ := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return left / right, nil
	} else if cop, ok := op.(*moduloOperation); ok {
		left, _ := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return math.Mod(left, right), nil
	} else if cop, ok := op.(*exponentiationOperation); ok {
		left, _ := execute(cop.Base, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Exponent, vars, functionRegistry, constantRegistry)

		return math.Pow(left, right), nil
	} else if cop, ok := op.(*unaryMinusOperation); ok {
		arg, _ := execute(cop.Operation, vars, functionRegistry, constantRegistry)
		return -arg, nil
	} else if cop, ok := op.(*andOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 && right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*orOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 || right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*lessThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left < right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*lessOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left <= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*greaterThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left > right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*greaterOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left >= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*equalOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left == right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*notEqualOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*functionOperation); ok {

		fn, _ := functionRegistry.get(cop.Name)
		arguments := make([]float64, len(cop.Arguments))

		for idx, fnParam := range cop.Arguments {
			arg, _ := execute(fnParam, vars, functionRegistry, constantRegistry)
			arguments[idx] = arg
		}
		return fn.function(arguments...)
	}

	return 0.0, fmt.Errorf("not implemented %T", op)
}
