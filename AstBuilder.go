package gojacego

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mrxrsd/gojacego/stack"
)

type AstBuilder struct {
	caseSensitive  bool
	resultStack    stack.Stack
	operatorStack  stack.Stack
	parameterCount stack.Stack
}

func NewAstBuilder(caseSensitive bool) *AstBuilder {

	resultStack := stack.New()
	operatorStack := stack.New()
	parameterCount := stack.New()

	return &AstBuilder{
		caseSensitive:  caseSensitive,
		resultStack:    *resultStack,
		operatorStack:  *operatorStack,
		parameterCount: *parameterCount,
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
			if err != nil {
				this.resultStack.Push(t)
			}
			break
		case TEXT:
			f, err := this.convertFunction(token)
			if err != nil {
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

func (this AstBuilder) convertFunction(operationToken Token) (Operation, error) {

	return nil, nil
}

func (this AstBuilder) convertOperation(operationToken Token) (Operation, error) {

	return nil, nil
}

func (this AstBuilder) Build(tokens []Token) (*Expression, error) {

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
			if false {
				// FUNCTION REGISTRY
			} else {

				if false {
					// constant registry
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
		}
	}

	return nil, nil
}
