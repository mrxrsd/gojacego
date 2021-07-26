package gojacego

type Optimizer struct {
	executor Interpreter
}

func (this *Optimizer) Optimize(op Operation, functionRegistry *FunctionRegistry, constantRegistry *ConstantRegistry) Operation {
	return optimize(this.executor, op, functionRegistry, constantRegistry)
}
func optimize(executor Interpreter, op Operation, functionRegistry *FunctionRegistry, constantRegistry *ConstantRegistry) Operation {

	if _, b := op.(*ConstantOperation); !op.OperationMetadata().DependsOnVariables && op.OperationMetadata().IsIdempotent && !b {
		result, _ := executor.Execute(op, nil, functionRegistry, constantRegistry)
		return NewConstantOperation(FloatingPoint, result)
	} else {

		if cop, ok := op.(*AddOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*SubtractionOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*MultiplicationOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne, functionRegistry, constantRegistry)
			cop.OperationTwo = optimize(executor, cop.OperationTwo, functionRegistry, constantRegistry)

			cop1, ok1 := cop.OperationOne.(*ConstantOperation)
			cop2, ok2 := cop.OperationTwo.(*ConstantOperation)

			if ok1 && cop1.Value == 0.0 || ok2 && cop2.Value == 0.0 {
				return NewConstantOperation(FloatingPoint, 0.0)
			}

		} else if cop, ok := op.(*DivisorOperation); ok {
			cop.Dividend = optimize(executor, cop.Dividend, functionRegistry, constantRegistry)
			cop.Divisor = optimize(executor, cop.Divisor, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*ExponentiationOperation); ok {
			cop.Base = optimize(executor, cop.Base, functionRegistry, constantRegistry)
			cop.Exponent = optimize(executor, cop.Exponent, functionRegistry, constantRegistry)

		} else if cop, ok := op.(*FunctionOperation); ok {
			optimizedArguments := make([]Operation, len(cop.Arguments))

			for idx, arg := range cop.Arguments {
				ret := optimize(executor, arg, functionRegistry, constantRegistry)
				optimizedArguments[idx] = ret
			}

			cop.Arguments = optimizedArguments

		}

		return op
	}
}
