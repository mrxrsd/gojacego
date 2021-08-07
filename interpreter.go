package gojacego

import (
	"errors"
	"fmt"
	"math"
)

type interpreter struct {
}

type Formula func(vars map[string]interface{}) (float64, error)

func (*interpreter) execute(op operation, vars formulaVariables, functionRegistry *functionRegistry, constantRegistry *constantRegistry) (ret float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()

	ret = execute(op, vars, functionRegistry, constantRegistry)
	return ret, err
}

func (*interpreter) buildFormula(op operation, functionRegistry *functionRegistry, constantRegistry *constantRegistry) Formula {
	return func(vars map[string]interface{}) (ret float64, err error) {

		defer func() {
			if r := recover(); r != nil {
				err = errors.New(r.(string))
			}
		}()

		ret = execute(op, vars, functionRegistry, constantRegistry)
		return ret, err
	}
}

func execute(op operation, vars formulaVariables, functionRegistry *functionRegistry, constantRegistry *constantRegistry) float64 {

	if op == nil {
		panic("operation cannot be nil")
	}

	if cop, ok := op.(*constantOperation); ok {
		if cop.Metadata.DataType == integer {
			return toFloat64Panic(cop.Value)
		} else {
			return cop.Value.(float64)
		}

	} else if cop, ok := op.(*variableOperation); ok {

		variableValue, err := vars.Get(cop.Name)
		if err == nil {
			return toFloat64Panic(variableValue)
		} else {
			panic("The variable '" + cop.Name + "' used is not defined.")
		}

	} else if cop, ok := op.(*multiplicationOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left * right
	} else if cop, ok := op.(*addOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left + right
	} else if cop, ok := op.(*subtractionOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left - right
	} else if cop, ok := op.(*divisorOperation); ok {
		left := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return left / right
	} else if cop, ok := op.(*moduloOperation); ok {
		left := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return math.Mod(left, right)
	} else if cop, ok := op.(*exponentiationOperation); ok {
		left := execute(cop.Base, vars, functionRegistry, constantRegistry)
		right := execute(cop.Exponent, vars, functionRegistry, constantRegistry)

		return math.Pow(left, right)
	} else if cop, ok := op.(*unaryMinusOperation); ok {
		arg := execute(cop.Operation, vars, functionRegistry, constantRegistry)
		return -arg
	} else if cop, ok := op.(*andOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 && right != 0 {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*orOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 || right != 0 {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*lessThanOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left < right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*lessOrEqualThanOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left <= right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*greaterThanOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left > right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*greaterOrEqualThanOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left >= right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*equalOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left == right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*notEqualOperation); ok {
		left := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != right {
			return 1.0
		}
		return 0.0
	} else if cop, ok := op.(*functionOperation); ok {

		fn, _ := functionRegistry.get(cop.Name)
		arguments := make([]float64, len(cop.Arguments))

		for idx, fnParam := range cop.Arguments {
			arg := execute(fnParam, vars, functionRegistry, constantRegistry)
			arguments[idx] = arg
		}
		ret, err := runDelegate(fn, arguments)
		if err != nil {
			panic(err.Error())
		}
		return ret
	}

	panic(fmt.Sprintf("not implemented %T", op))
}

func runDelegate(fn *functionInfo, arguments []float64) (ret float64, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("function '%s': runtime error (%T)", fn.name, r)
		}
	}()

	ret = fn.function(arguments...)
	return ret, err
}
