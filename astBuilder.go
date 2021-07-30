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

type astBuilder struct {
	caseSensitive            bool
	resultStack              *stack.Stack
	operatorStack            *stack.Stack
	parameterCount           *stack.Stack
	constantRegistry         *constantRegistry
	compiledConstantRegistry *constantRegistry
	functionRegistry         *functionRegistry
}

func newAstBuilder(caseSensitive bool, functionRegistry *functionRegistry, constantRegistry *constantRegistry, compiledConstantRegistry *constantRegistry) *astBuilder {

	resultStack := stack.New()
	operatorStack := stack.New()
	parameterCount := stack.New()

	return &astBuilder{
		caseSensitive:            caseSensitive,
		resultStack:              resultStack,
		operatorStack:            operatorStack,
		parameterCount:           parameterCount,
		constantRegistry:         constantRegistry,
		functionRegistry:         functionRegistry,
		compiledConstantRegistry: compiledConstantRegistry,
	}
}

func (this astBuilder) popOperations(untilLeftBracket bool, currentToken *Token) error {

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

func (this astBuilder) verifyResult() error {
	if this.resultStack.Len() > 1 {
		return errors.New("The syntax of the provided formula is not valid.")
	}
	return nil
}

func (this astBuilder) convertFunction(operationToken Token) (operation, error) {

	functionName := operationToken.Value.(string)

	if item, found := this.functionRegistry.get(functionName); found {

		var numberOfParameters int
		if true {
			numberOfParameters = this.parameterCount.Pop().(int)
		} else {
			// fixed parameter
		}

		operations := make([]operation, numberOfParameters)
		for i := 0; i < numberOfParameters; i++ {
			operations[i] = this.resultStack.Pop().(operation)
		}

		// vscode reverse
		for i, j := 0, len(operations)-1; i < j; i, j = i+1, j-1 {
			operations[i], operations[j] = operations[j], operations[i]
		}

		return newFunctionOperation(floatingPoint, functionName, operations, item.isIdempotent), nil
	}

	return nil, nil
}

func (this astBuilder) convertOperation(operationToken Token) (operation, error) {

	var dataType operationDataType
	var argument1 operation
	var argument2 operation
	var divisor operation
	var divident operation

	switch rune(operationToken.Value.(int32)) {
	case '+':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newAddOperation(dataType, argument1, argument2), nil
	case '-':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newSubtractionOperation(dataType, argument1, argument2), nil
	case '*':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newMultiplicationOperation(dataType, argument1, argument2), nil
	case '/':
		divisor = this.resultStack.Pop().(operation)
		divident = this.resultStack.Pop().(operation)
		return newDivisorOperation(floatingPoint, divident, divisor), nil
	case '%':
		divisor = this.resultStack.Pop().(operation)
		divident = this.resultStack.Pop().(operation)
		return newModuloOperation(floatingPoint, divident, divisor), nil
	case '_':
		argument1 = this.resultStack.Pop().(operation)
		return NewUnaryMinusOperation(argument1.OperationMetadata().DataType, argument1), nil
	case '^':
		exponent := this.resultStack.Pop().(operation)
		base := this.resultStack.Pop().(operation)
		return newExponentiationOperation(floatingPoint, base, exponent), nil
	case '&':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return NewAndOperation(dataType, argument1, argument2), nil
	case '|':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newOrOperation(dataType, argument1, argument2), nil
	case '<':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newLessThanOperation(dataType, argument1, argument2), nil
	case '≤':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newLessOrEqualThanOperation(dataType, argument1, argument2), nil
	case '>':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newGreaterThanOperation(dataType, argument1, argument2), nil
	case '≥':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newGreaterOrEqualThanOperation(dataType, argument1, argument2), nil
	case '=':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newEqualOperation(dataType, argument1, argument2), nil
	case '≠':
		argument2 = this.resultStack.Pop().(operation)
		argument1 = this.resultStack.Pop().(operation)
		dataType = requiredDataType(argument1, argument2)
		return newNotEqualOperation(dataType, argument1, argument2), nil
	default:
		return nil, fmt.Errorf("unknown operation %s", operationToken.Value)
	}
}

func (this astBuilder) build(tokens []Token) (operation, error) {

	for _, token := range tokens {
		val := token.Value

		switch token.Type {
		case INTEGER:
			this.resultStack.Push(newConstantOperation(integer, val))
			break
		case FLOATING_POINT:
			this.resultStack.Push(newConstantOperation(floatingPoint, val))
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
						this.resultStack.Push(newConstantOperation(floatingPoint, val))
						break
					}
				}

				if val, found := this.constantRegistry.get(tokenText); found {
					// constant registry
					this.resultStack.Push(newConstantOperation(floatingPoint, val))
				} else {
					if !this.caseSensitive {
						tokenText = strings.ToLower(tokenText)
					}
					this.resultStack.Push(newVariableOperation(tokenText))
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
		resultOperation := this.resultStack.Pop().(operation)
		return resultOperation, nil
	}
}

func isLeftAssociativeOperation(character rune) bool {
	return character == '*' || character == '+' || character == '-' || character == '/'
}

func requiredDataType(argument1 operation, argument2 operation) operationDataType {
	if argument1.OperationMetadata().DataType == floatingPoint || argument2.OperationMetadata().DataType == floatingPoint {
		return floatingPoint
	}
	return integer
}
