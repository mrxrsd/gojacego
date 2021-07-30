package gojacego

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mrxrsd/gojacego/stack"
)

var precedences = map[rune]int{
	'(': 0,
	'&': 1,
	'|': 1,
	'<': 2,
	'>': 2,
	'≤': 2,
	'≥': 2,
	'≠': 2,
	'=': 2,
	'+': 3,
	'-': 3,
	'*': 4,
	'/': 4,
	'%': 4,
	'_': 5,
	'^': 6,
}

type AstBuilder struct {
	caseSensitive            bool
	resultStack              *stack.Stack
	operatorStack            *stack.Stack
	parameterCount           *stack.Stack
	constantRegistry         *constantRegistry
	compiledConstantRegistry *constantRegistry
	functionRegistry         *functionRegistry
}

func NewAstBuilder(caseSensitive bool, functionRegistry *functionRegistry, constantRegistry *constantRegistry, compiledConstantRegistry *constantRegistry) *AstBuilder {

	resultStack := stack.New()
	operatorStack := stack.New()
	parameterCount := stack.New()

	return &AstBuilder{
		caseSensitive:            caseSensitive,
		resultStack:              resultStack,
		operatorStack:            operatorStack,
		parameterCount:           parameterCount,
		constantRegistry:         constantRegistry,
		functionRegistry:         functionRegistry,
		compiledConstantRegistry: compiledConstantRegistry,
	}
}

func (this AstBuilder) popOperations(untilLeftBracket bool, currentToken *Token) error {

	if untilLeftBracket && currentToken == nil {
		return errors.New("If the parameter \"untillLeftBracket\" is set to true, " +
			"the parameter \"currentToken\" cannot be null.")
	}

	for this.operatorStack.Len() > 0 && this.operatorStack.Peek().(Token).Type != LEFT_BRACKET {

		token := this.operatorStack.Pop().(Token)

		switch token.Type {
		case OPERATION:
			t, err := this.convertOperation(token)
			if err == nil {
				this.resultStack.Push(t)
			}
			break
		case TEXT:
			f, err := this.convertFunction(token)
			if err == nil {
				this.resultStack.Push(f)
			}
			break
		}
	}

	if untilLeftBracket {
		if this.operatorStack.Len() > 0 && this.operatorStack.Peek().(Token).Type == LEFT_BRACKET {
			this.operatorStack.Pop()
		} else {
			return errors.New(fmt.Sprintf("No matching left bracket found for the right "+
				"bracket at position %d.", currentToken.StartPosition))
		}
	} else {
		if this.operatorStack.Len() > 0 && this.operatorStack.Peek().(Token).Type == LEFT_BRACKET && !(currentToken != nil && currentToken.Type == ARGUMENT_SEPARATOR) {
			return errors.New(fmt.Sprintf("No matching right bracket found for the left "+
				"bracket at position %d.", this.operatorStack.Peek().(Token).StartPosition))
		}
	}

	return nil
}

func (this AstBuilder) verifyResult() error {
	if this.resultStack.Len() > 1 {
		return errors.New("The syntax of the provided formula is not valid.")
	}
	return nil
}

func (this AstBuilder) convertFunction(operationToken Token) (Operation, error) {

	functionName := operationToken.Value.(string)

	if item, found := this.functionRegistry.get(functionName); found {

		var numberOfParameters int
		if true {
			numberOfParameters = this.parameterCount.Pop().(int)
		} else {
			// fixed parameter
		}

		operations := make([]Operation, numberOfParameters)
		for i := 0; i < numberOfParameters; i++ {
			operations[i] = this.resultStack.Pop().(Operation)
		}

		// vscode reverse
		for i, j := 0, len(operations)-1; i < j; i, j = i+1, j-1 {
			operations[i], operations[j] = operations[j], operations[i]
		}

		return NewFunctionOperation(FloatingPoint, functionName, operations, item.isIdempotent), nil
	}

	return nil, nil
}

func (this AstBuilder) convertOperation(operationToken Token) (Operation, error) {

	var dataType OperationDataType
	var argument1 Operation
	var argument2 Operation
	var divisor Operation
	var divident Operation

	switch rune(operationToken.Value.(int32)) {
	case '+':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewAddOperation(dataType, argument1, argument2), nil
	case '-':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewSubtractionOperation(dataType, argument1, argument2), nil
	case '*':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewMultiplicationOperation(dataType, argument1, argument2), nil
	case '/':
		divisor = this.resultStack.Pop().(Operation)
		divident = this.resultStack.Pop().(Operation)
		return NewDivisorOperation(FloatingPoint, divident, divisor), nil
	case '%':
		divisor = this.resultStack.Pop().(Operation)
		divident = this.resultStack.Pop().(Operation)
		return NewModuloOperation(FloatingPoint, divident, divisor), nil
	case '_':
		argument1 = this.resultStack.Pop().(Operation)
		return NewUnaryMinusOperation(argument1.OperationMetadata().DataType, argument1), nil
	case '^':
		exponent := this.resultStack.Pop().(Operation)
		base := this.resultStack.Pop().(Operation)
		return NewExponentiationOperation(FloatingPoint, base, exponent), nil
	case '&':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewAndOperation(dataType, argument1, argument2), nil
	case '|':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewOrOperation(dataType, argument1, argument2), nil
	case '<':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewLessThanOperation(dataType, argument1, argument2), nil
	case '≤':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewLessOrEqualThanOperation(dataType, argument1, argument2), nil
	case '>':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewGreaterThanOperation(dataType, argument1, argument2), nil
	case '≥':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewGreaterOrEqualThanOperation(dataType, argument1, argument2), nil
	case '=':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewEqualOperation(dataType, argument1, argument2), nil
	case '≠':
		argument2 = this.resultStack.Pop().(Operation)
		argument1 = this.resultStack.Pop().(Operation)
		dataType = requiredDataType(argument1, argument2)
		return NewNotEqualOperation(dataType, argument1, argument2), nil
	default:
		return nil, fmt.Errorf("unknown operation %s", operationToken.Value)
	}
}

func (this AstBuilder) Build(tokens []Token) (Operation, error) {

	for _, token := range tokens {
		val := token.Value

		switch token.Type {
		case INTEGER:
			this.resultStack.Push(NewConstantOperation(Integer, val))
			break
		case FLOATING_POINT:
			this.resultStack.Push(NewConstantOperation(FloatingPoint, val))
			break
		case TEXT:
			tokenText := token.Value.(string)
			if _, found := this.functionRegistry.get(tokenText); found {
				this.operatorStack.Push(token)
				this.parameterCount.Push(1)
			} else {

				if this.compiledConstantRegistry != nil {
					if val, found := this.compiledConstantRegistry.get(tokenText); found {
						// constant registry
						this.resultStack.Push(NewConstantOperation(FloatingPoint, val))
						break
					}
				}

				if val, found := this.constantRegistry.get(tokenText); found {
					// constant registry
					this.resultStack.Push(NewConstantOperation(FloatingPoint, val))
				} else {
					if !this.caseSensitive {
						tokenText = strings.ToLower(tokenText)
					}
					this.resultStack.Push(NewVariableOperation(tokenText))
				}
			}
			break
		case LEFT_BRACKET:
			this.operatorStack.Push(token)
			break
		case RIGHT_BRACKET:
			this.popOperations(true, &token)
			break
		case ARGUMENT_SEPARATOR:
			this.popOperations(false, &token)
			this.parameterCount.Push(this.parameterCount.Pop().(int) + 1)
			break
		case OPERATION:
			operation1Token := token
			// operation1 := []rune(operation1Token.Value.(string))[0]
			operation1 := rune(operation1Token.Value.(int32))

			for this.operatorStack.Len() > 0 && (this.operatorStack.Peek().(Token).Type == OPERATION || this.operatorStack.Peek().(Token).Type == TEXT) {

				var operation2Token Token
				operation2Token = this.operatorStack.Peek().(Token)

				isFunctionOnTopOfStack := false
				if operation2Token.Type == TEXT {
					isFunctionOnTopOfStack = true
				}

				if !isFunctionOnTopOfStack {
					operation2 := rune(operation2Token.Value.(int32))
					// operation2 = []rune(operation2Token.Value.(string))[0]

					if (isLeftAssociativeOperation(operation1) && precedences[operation1] <= precedences[operation2]) || (precedences[operation1] < precedences[operation2]) {
						this.operatorStack.Pop()
						t, err := this.convertOperation(operation2Token)
						if err == nil {
							this.resultStack.Push(t)
						}
					} else {
						break
					}
				} else {
					this.operatorStack.Pop()
					t, err := this.convertFunction(operation2Token)
					if err == nil {
						this.resultStack.Push(t)
					}
				}
			}

			this.operatorStack.Push(operation1Token)
			break
		}
	}

	this.popOperations(false, nil)

	err := this.verifyResult()

	if err != nil {
		return nil, err
	} else {
		resultOperation := this.resultStack.Pop().(Operation)
		return resultOperation, nil
	}
}

func isLeftAssociativeOperation(character rune) bool {
	return character == '*' || character == '+' || character == '-' || character == '/'
}

func requiredDataType(argument1 Operation, argument2 Operation) OperationDataType {
	if argument1.OperationMetadata().DataType == FloatingPoint || argument2.OperationMetadata().DataType == FloatingPoint {
		return FloatingPoint
	}
	return Integer
}
