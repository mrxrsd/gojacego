package gojacego

type operationDataType int

const (
	integer operationDataType = iota
	floatingPoint
)

type operationMetadata struct {
	DataType           operationDataType
	DependsOnVariables bool
	IsIdempotent       bool
}

type operation interface {
	OperationMetadata() operationMetadata
}

// Operations

// Variable
type variableOperation struct {
	Name     string
	Metadata operationMetadata
}

func (op *variableOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newVariableOperation(name string) *variableOperation {
	meta := operationMetadata{
		DataType:           floatingPoint,
		DependsOnVariables: true,
		IsIdempotent:       false,
	}

	return &variableOperation{
		Name:     name,
		Metadata: meta,
	}
}

// Add
type addOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *addOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newAddOperation(dataType operationDataType, operationOne operation, operationTwo operation) *addOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &addOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// And
type andOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *andOperation) OperationMetadata() operationMetadata { return op.Metadata }

func NewAndOperation(dataType operationDataType, operationOne operation, operationTwo operation) *andOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &andOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// Constant
type constantOperation struct {
	Value    interface{}
	Metadata operationMetadata
}

func (op *constantOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newConstantOperation(dataType operationDataType, value interface{}) *constantOperation {
	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: false,
		IsIdempotent:       true,
	}

	return &constantOperation{
		Value:    value,
		Metadata: meta,
	}
}

// Divisor
type divisorOperation struct {
	Dividend operation
	Divisor  operation
	Metadata operationMetadata
}

func (op *divisorOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newDivisorOperation(dataType operationDataType, dividend operation, divisor operation) *divisorOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: dividend.OperationMetadata().DependsOnVariables || divisor.OperationMetadata().DependsOnVariables,
		IsIdempotent:       dividend.OperationMetadata().IsIdempotent && divisor.OperationMetadata().DependsOnVariables,
	}

	return &divisorOperation{
		Dividend: dividend,
		Divisor:  divisor,
		Metadata: meta,
	}
}

// Equal
type equalOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *equalOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newEqualOperation(dataType operationDataType, operationOne operation, operationTwo operation) *equalOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &equalOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// Exponentiation
type exponentiationOperation struct {
	Base     operation
	Exponent operation
	Metadata operationMetadata
}

func (op *exponentiationOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newExponentiationOperation(dataType operationDataType, base operation, exponent operation) *exponentiationOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: base.OperationMetadata().DependsOnVariables || exponent.OperationMetadata().DependsOnVariables,
		IsIdempotent:       base.OperationMetadata().IsIdempotent && exponent.OperationMetadata().DependsOnVariables,
	}

	return &exponentiationOperation{
		Base:     base,
		Exponent: exponent,
		Metadata: meta,
	}
}

// Function
type functionOperation struct {
	Name      string
	Arguments []operation
	Metadata  operationMetadata
}

func (op *functionOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newFunctionOperation(dataType operationDataType, name string, arguments []operation, isIdempotent bool) *functionOperation {

	anyDependesOnVars := false
	allIsIdempotent := isIdempotent
	for _, v := range arguments {
		anyDependesOnVars = anyDependesOnVars || v.OperationMetadata().DependsOnVariables
		allIsIdempotent = allIsIdempotent && v.OperationMetadata().IsIdempotent
	}

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: anyDependesOnVars,
		IsIdempotent:       allIsIdempotent,
	}

	return &functionOperation{
		Name:      name,
		Arguments: arguments,
		Metadata:  meta,
	}
}

// GreaterOrEqualThan
type greaterOrEqualThanOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *greaterOrEqualThanOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newGreaterOrEqualThanOperation(dataType operationDataType, operationOne operation, operationTwo operation) *greaterOrEqualThanOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &greaterOrEqualThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// GreaterThanOperation
type greaterThanOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *greaterThanOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newGreaterThanOperation(dataType operationDataType, operationOne operation, operationTwo operation) *greaterThanOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &greaterThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// LessOrEqualThan
type lessOrEqualThanOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *lessOrEqualThanOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newLessOrEqualThanOperation(dataType operationDataType, operationOne operation, operationTwo operation) *lessOrEqualThanOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &lessOrEqualThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// LessThan
type lessThanOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *lessThanOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newLessThanOperation(dataType operationDataType, operationOne operation, operationTwo operation) *lessThanOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &lessThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Modulo
type moduloOperation struct {
	Dividend operation
	Divisor  operation
	Metadata operationMetadata
}

func (op *moduloOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newModuloOperation(dataType operationDataType, dividend operation, divisor operation) *moduloOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: dividend.OperationMetadata().DependsOnVariables || divisor.OperationMetadata().DependsOnVariables,
		IsIdempotent:       dividend.OperationMetadata().IsIdempotent && divisor.OperationMetadata().DependsOnVariables,
	}

	return &moduloOperation{
		Dividend: dividend,
		Divisor:  divisor,
		Metadata: meta,
	}
}

//Mutiplication
type multiplicationOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *multiplicationOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newMultiplicationOperation(dataType operationDataType, operationOne operation, operationTwo operation) *multiplicationOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &multiplicationOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Not Equal
type notEqualOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *notEqualOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newNotEqualOperation(dataType operationDataType, operationOne operation, operationTwo operation) *notEqualOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &notEqualOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Or
type orOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *orOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newOrOperation(dataType operationDataType, operationOne operation, operationTwo operation) *orOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &orOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Subtraction
type subtractionOperation struct {
	OperationOne operation
	OperationTwo operation
	Metadata     operationMetadata
}

func (op *subtractionOperation) OperationMetadata() operationMetadata { return op.Metadata }

func newSubtractionOperation(dataType operationDataType, operationOne operation, operationTwo operation) *subtractionOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &subtractionOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//UnaryMinus
type unaryMinusOperation struct {
	Operation operation
	Metadata  operationMetadata
}

func (op *unaryMinusOperation) OperationMetadata() operationMetadata { return op.Metadata }

func NewUnaryMinusOperation(dataType operationDataType, operation operation) *unaryMinusOperation {

	meta := operationMetadata{
		DataType:           dataType,
		DependsOnVariables: operation.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operation.OperationMetadata().IsIdempotent,
	}

	return &unaryMinusOperation{
		Operation: operation,
		Metadata:  meta,
	}
}
