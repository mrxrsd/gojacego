package gojacego

import (
	"errors"
	"fmt"
	"math"
)

type Interpreter struct {
	caseSensitive bool
}

type Formula func(vars FormulaVariables) float64

func (*Interpreter) Execute(op Operation, vars FormulaVariables, functionRegistry *FunctionRegistry, constantRegistry *ConstantRegistry) (float64, error) {
	return execute(op, vars, functionRegistry, constantRegistry)
}

func (*Interpreter) BuildFormula(op Operation, functionRegistry *FunctionRegistry, constantRegistry *ConstantRegistry) Formula {
	return func(vars FormulaVariables) float64 {
		ret, _ := execute(op, vars, functionRegistry, constantRegistry)
		return ret
	}
}

func execute(op Operation, vars FormulaVariables, functionRegistry *FunctionRegistry, constantRegistry *ConstantRegistry) (float64, error) {

	if op == nil {
		return 0, errors.New("operation cannot be nil")
	}

	if cop, ok := op.(*ConstantOperation); ok {
		if cop.Metadata.DataType == Integer {
			return ToFloat64(cop.Value), nil
		} else {
			return cop.Value.(float64), nil
		}

	} else if cop, ok := op.(*VariableOperation); ok {

		variableValue, err := vars.Get(cop.Name)
		if err == nil {
			return variableValue, nil
		} else {
			return 0, errors.New("The variable '" + cop.Name + "' used is not defined.")
		}

	} else if cop, ok := op.(*MultiplicationOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left * right, nil
	} else if cop, ok := op.(*AddOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left + right, nil
	} else if cop, ok := op.(*SubtractionOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		return left - right, nil
	} else if cop, ok := op.(*DivisorOperation); ok {
		left, _ := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return left / right, nil
	} else if cop, ok := op.(*ModuloOperation); ok {
		left, _ := execute(cop.Dividend, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Divisor, vars, functionRegistry, constantRegistry)

		return math.Mod(left, right), nil
	} else if cop, ok := op.(*ExponentiationOperation); ok {
		left, _ := execute(cop.Base, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.Exponent, vars, functionRegistry, constantRegistry)

		return math.Pow(left, right), nil
	} else if cop, ok := op.(*UnaryMinusOperation); ok {
		arg, _ := execute(cop.Operation, vars, functionRegistry, constantRegistry)
		return -arg, nil
	} else if cop, ok := op.(*AndOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 && right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*OrOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != 0 || right != 0 {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*LessThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left < right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*LessOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left <= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*GreaterThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left > right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*GreaterOrEqualThanOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left >= right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*EqualOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left == right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*NotEqualOperation); ok {
		left, _ := execute(cop.OperationOne, vars, functionRegistry, constantRegistry)
		right, _ := execute(cop.OperationTwo, vars, functionRegistry, constantRegistry)

		if left != right {
			return 1.0, nil
		}
		return 0.0, nil
	} else if cop, ok := op.(*FunctionOperation); ok {

		fn, _ := functionRegistry.Get(cop.Name)
		arguments := make([]float64, len(cop.Arguments))

		for idx, fnParam := range cop.Arguments {
			arg, _ := execute(fnParam, vars, functionRegistry, constantRegistry)
			arguments[idx] = arg
		}
		return fn.function(arguments...)
	}

	return 0.0, fmt.Errorf("not implemented %T", op)
}
