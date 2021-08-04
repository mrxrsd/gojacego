package gojacego

type optimizer struct {
	executor interpreter
}

func (this *optimizer) optimize(op operation, functionRegistry *functionRegistry, constantRegistry *constantRegistry) operation {
	return optimize(this.executor, op, functionRegistry, constantRegistry)
}
func optimize(executor interpreter, op operation, functionRegistry *functionRegistry, constantRegistry *constantRegistry) operation {

	if _, b := op.(*constantOperation); !op.OperationMetadata().DependsOnVariables && op.OperationMetadata().IsIdempotent && !b {
		result, _ := executor.execute(op, nil, functionRegistry, constantRegistry)
		return newConstantOperation(floatingPoint, result)
	} else {

		if cop, ok := op.(*addOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*subtractionOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*multiplicationOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

			cop1, ok1 := cop.OperationOne.(*constantOperation)
			cop2, ok2 := cop.OperationTwo.(*constantOperation)

			if ok1 && cop1.Value == 0.0 || ok2 && cop2.Value == 0.0 {
				return newConstantOperation(floatingPoint, 0.0)
			}

		} else if cop, ok := op.(*divisorOperation); ok {
			cop.Dividend = optimize(executor, cop.Dividend, functionRegistry, constantRegistry)
			cop.Divisor = optimize(executor, cop.Divisor, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*exponentiationOperation); ok {
			cop.Base = optimize(executor, cop.Base, functionRegistry, constantRegistry)
			cop.Exponent = optimize(executor, cop.Exponent, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*greaterThanOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*greaterOrEqualThanOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*andOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*orOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*lessThanOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*lessOrEqualThanOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*functionOperation); ok {
			optimizedArguments := make([]operation, len(cop.Arguments))

			for idx, arg := range cop.Arguments {
				ret := optimize(executor, arg, functionRegistry, constantRegistry)
				optimizedArguments[idx] = ret
			}

			cop.Arguments = optimizedArguments

		}

		return op
	}
}
