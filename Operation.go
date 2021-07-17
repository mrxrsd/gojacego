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

func (varOp *VariableOperation) OperationMetadata() OperationMetadata { return varOp.Metadata }

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

func (addOp *AddOperation) OperationMetadata() OperationMetadata { return addOp.Metadata }

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

func (andOp *AndOperation) OperationMetadata() OperationMetadata { return andOp.Metadata }

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
