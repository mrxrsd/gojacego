package gojacego

type Expression struct {
	Operation []Operation
}
type OperationDataType int

const (
	Integer OperationDataType = iota
	FloatingPoint
)

type OperationMetadata struct {
	DataType           OperationDataType
	DependsOnVariables bool
	IsIdempotent       bool
}

type Operation interface {
	OperationMetadata() OperationMetadata
}

// Operations

// Variable
type VariableOperation struct {
	Name     string
	Metadata OperationMetadata
}

func (op *VariableOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewVariableOperation(name string) *VariableOperation {
	meta := OperationMetadata{
		DataType:           FloatingPoint,
		DependsOnVariables: true,
		IsIdempotent:       false,
	}

	return &VariableOperation{
		Name:     name,
		Metadata: meta,
	}
}

// Add
type AddOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *AddOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewAddOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *AddOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &AddOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// And
type AndOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *AndOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewAndOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *AndOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &AndOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// Constant
type ConstantOperation struct {
	Value    interface{}
	Metadata OperationMetadata
}

func (op *ConstantOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewConstantOperation(dataType OperationDataType, value interface{}) *ConstantOperation {
	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: false,
		IsIdempotent:       true,
	}

	return &ConstantOperation{
		Value:    value,
		Metadata: meta,
	}
}

// Divisor
type DivisorOperation struct {
	Dividend Operation
	Divisor  Operation
	Metadata OperationMetadata
}

func (op *DivisorOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewDivisorOperation(dataType OperationDataType, dividend Operation, divisor Operation) *DivisorOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: dividend.OperationMetadata().DependsOnVariables || divisor.OperationMetadata().DependsOnVariables,
		IsIdempotent:       dividend.OperationMetadata().IsIdempotent && divisor.OperationMetadata().DependsOnVariables,
	}

	return &DivisorOperation{
		Dividend: dividend,
		Divisor:  divisor,
		Metadata: meta,
	}
}

// Equal
type EqualOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *EqualOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewEqualOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *EqualOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &EqualOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// Exponentiation
type ExponentiationOperation struct {
	Base     Operation
	Exponent Operation
	Metadata OperationMetadata
}

func (op *ExponentiationOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewExponentiationOperation(dataType OperationDataType, base Operation, exponent Operation) *ExponentiationOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: base.OperationMetadata().DependsOnVariables || exponent.OperationMetadata().DependsOnVariables,
		IsIdempotent:       base.OperationMetadata().IsIdempotent && exponent.OperationMetadata().DependsOnVariables,
	}

	return &ExponentiationOperation{
		Base:     base,
		Exponent: exponent,
		Metadata: meta,
	}
}

// Function
type FunctionOperation struct {
	Name      string
	Arguments []Operation
	Metadata  OperationMetadata
}

func (op *FunctionOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewFunctionOperation(datatype OperationDataType, name string, arguments []Operation, isIdempotent bool) *FunctionOperation {

	anyDependesOnVars := false
	allIsIdempotent := isIdempotent
	for _, v := range arguments {
		anyDependesOnVars = anyDependesOnVars || v.OperationMetadata().DependsOnVariables
		allIsIdempotent = allIsIdempotent && v.OperationMetadata().IsIdempotent
	}

	meta := OperationMetadata{
		DataType:           datatype,
		DependsOnVariables: anyDependesOnVars,
		IsIdempotent:       allIsIdempotent,
	}

	return &FunctionOperation{
		Name:      name,
		Arguments: arguments,
		Metadata:  meta,
	}
}

// GreaterOrEqualThan
type GreaterOrEqualThanOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *GreaterOrEqualThanOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewGreaterOrEqualThanOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *GreaterOrEqualThanOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &GreaterOrEqualThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// GreaterThanOperation
type GreaterThanOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *GreaterThanOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewGreaterThanOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *GreaterThanOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &GreaterThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// LessOrEqualThan
type LessOrEqualThanOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *LessOrEqualThanOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewLessOrEqualThanOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *LessOrEqualThanOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &LessOrEqualThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

// LessThan
type LessThanOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *LessThanOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewLessThanOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *LessThanOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &LessThanOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Modulo
type ModuloOperation struct {
	Dividend Operation
	Divisor  Operation
	Metadata OperationMetadata
}

func (op *ModuloOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewModuloOperation(dataType OperationDataType, dividend Operation, divisor Operation) *ModuloOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: dividend.OperationMetadata().DependsOnVariables || divisor.OperationMetadata().DependsOnVariables,
		IsIdempotent:       dividend.OperationMetadata().IsIdempotent && divisor.OperationMetadata().DependsOnVariables,
	}

	return &ModuloOperation{
		Dividend: dividend,
		Divisor:  divisor,
		Metadata: meta,
	}
}

//Mutiplication
type MultiplicationOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *MultiplicationOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewMultiplicationOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *MultiplicationOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &MultiplicationOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Not Equal
type NotEqualOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *NotEqualOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewNotEqualOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *NotEqualOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &NotEqualOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Or
type OrOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *OrOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewOrOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *OrOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &OrOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//Subtraction
type SubtractionOperation struct {
	OperationOne Operation
	OperationTwo Operation
	Metadata     OperationMetadata
}

func (op *SubtractionOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewSubtractionOperation(dataType OperationDataType, operationOne Operation, operationTwo Operation) *SubtractionOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operationOne.OperationMetadata().DependsOnVariables || operationTwo.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operationOne.OperationMetadata().IsIdempotent && operationTwo.OperationMetadata().DependsOnVariables,
	}

	return &SubtractionOperation{
		OperationOne: operationOne,
		OperationTwo: operationTwo,
		Metadata:     meta,
	}
}

//UnaryMinus
type UnaryMinusOperation struct {
	Operation Operation
	Metadata  OperationMetadata
}

func (op *UnaryMinusOperation) OperationMetadata() OperationMetadata { return op.Metadata }

func NewUnaryMinusOperation(dataType OperationDataType, operation Operation) *UnaryMinusOperation {

	meta := OperationMetadata{
		DataType:           dataType,
		DependsOnVariables: operation.OperationMetadata().DependsOnVariables,
		IsIdempotent:       operation.OperationMetadata().IsIdempotent,
	}

	return &UnaryMinusOperation{
		Operation: operation,
		Metadata:  meta,
	}
}
