package gojacego

type Optimizer struct {
	executor Interpreter
}

type IOptimizer interface {
	Optimize(op Operation) Operation
}

func (this *Optimizer) Optimize(op Operation) Operation {
	return optimize(this.executor, op)
}
func optimize(executor Interpreter, op Operation) Operation {

	if !op.OperationMetadata().DependsOnVariables && op.OperationMetadata().IsIdempotent && op.(*ConstantOperation) != nil {
		result, _ := executor.Execute(op, nil)
		return NewConstantOperation(FloatingPoint, result)
	} else {

		if cop, ok := op.(*AddOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne)
			cop.OperationTwo = optimize(executor, cop.OperationTwo)

		} else if cop, ok := op.(*SubtractionOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne)
			cop.OperationTwo = optimize(executor, cop.OperationTwo)

		} else if cop, ok := op.(*MultiplicationOperation); ok {
			cop.OperationOne = optimize(executor, cop.OperationOne)
			cop.OperationTwo = optimize(executor, cop.OperationTwo)

			cop1, ok1 := cop.OperationOne.(*ConstantOperation)
			cop2, ok2 := cop.OperationTwo.(*ConstantOperation)

			if ok1 && cop1.Value == 0.0 || ok2 && cop2.Value == 0.0 {
				return NewConstantOperation(FloatingPoint, 0.0)
			}

		} else if cop, ok := op.(*DivisorOperation); ok {
			cop.Dividend = optimize(executor, cop.Dividend)
			cop.Divisor = optimize(executor, cop.Divisor)

		} else if cop, ok := op.(*ExponentiationOperation); ok {
			cop.Base = optimize(executor, cop.Base)
			cop.Exponent = optimize(executor, cop.Exponent)

		} else if _, ok := op.(*FunctionOperation); ok {
			panic("Not implemented.")
		}

		return op
	}
}
