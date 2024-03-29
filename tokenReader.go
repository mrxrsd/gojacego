package gojacego

import (
	"errors"
	"fmt"
	"strconv"
)

type tokenReader struct {
	decimalSeparator  rune
	argumentSeparator rune
}

func newTokenReader(decimalSeparator rune, argumentSeparador rune) *tokenReader {
	return &tokenReader{
		decimalSeparator:  decimalSeparator,
		argumentSeparator: argumentSeparador,
	}
}

func (this tokenReader) read(formula string) ([]token, error) {
	ret := make([]token, 0)

	if formula == "" {
		return nil, errors.New("formula cannot be empty")
	}

	runes := []rune(formula)
	runesLength := len(runes)
	isFormulaSubPart := true
	isScientific := false

	for i := 0; i < runesLength; i++ {
		if this.isPartOfNumeric(runes[i], true, false, isFormulaSubPart) {
			buffer := make([]rune, 0)
			buffer = append(buffer, runes[i])
			startPosition := i

			i++
			for i < runesLength {
				if !this.isPartOfNumeric(runes[i], false, runes[i-1] == '-', isFormulaSubPart) {
					break
				}

				if isScientific && this.isScientificNotation(runes[i]) {
					return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
				}

				if this.isScientificNotation(runes[i]) {
					isScientific = true
					if len(runes) > i+1 && runes[i+1] == '-' {
						buffer = append(buffer, runes[i])
						i++
					}
				}

				buffer = append(buffer, runes[i])
				i++
			}

			if intVal, err := strconv.ParseInt(string(buffer), 10, 64); err == nil {
				ret = append(ret, token{Type: tt_INTEGER,
					Value:         intVal,
					StartPosition: startPosition,
					Length:        i - startPosition})
			} else {

				if floatVal, err := strconv.ParseFloat(string(buffer), 64); err == nil {
					ret = append(ret, token{Type: tt_FLOATING_POINT,
						Value:         floatVal,
						StartPosition: startPosition,
						Length:        i - startPosition})

					isScientific = false
					isFormulaSubPart = false
				} else if string(buffer) == "-" {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '_',
						StartPosition: startPosition,
						Length:        i - startPosition})
				} else {
					return nil, fmt.Errorf("invalid floating point number: %s", string(buffer))
				}
			}

			if i == runesLength {
				continue
			}

		}

		if this.isPartOfVariable(runes[i], true) {
			buffer := make([]rune, 0)
			buffer = append(buffer, runes[i])
			startPosition := i

			i++
			for i < runesLength {
				if !this.isPartOfVariable(runes[i], false) {
					break
				}
				buffer = append(buffer, runes[i])
				i++
			}

			ret = append(ret, token{Type: tt_TEXT,
				Value:         string(buffer),
				StartPosition: startPosition,
				Length:        i - startPosition})

			isFormulaSubPart = false

			if i == runesLength {
				continue
			}
		}

		if runes[i] == this.argumentSeparator {
			ret = append(ret, token{Type: tt_ARGUMENT_SEPARATOR,
				Value:         runes[i],
				StartPosition: i,
				Length:        1})
		} else {

			switch runes[i] {
			case ' ':
				continue
			case '+', '-', '*', '/', '^', '%', '≤', '≥', '≠':
				if this.isUnaryMinus(runes[i], ret) {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '_',
						StartPosition: i,
						Length:        1})
				} else {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         runes[i],
						StartPosition: i,
						Length:        1})

				}
				isFormulaSubPart = true
			case '(':

				ret = append(ret, token{Type: tt_LEFT_BRACKET,
					Value:         runes[i],
					StartPosition: i,
					Length:        1})
				isFormulaSubPart = true
			case ')':
				ret = append(ret, token{Type: tt_RIGHT_BRACKET,
					Value:         runes[i],
					StartPosition: i,
					Length:        1})
				isFormulaSubPart = false
			case '<':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '≤',
						StartPosition: i,
						Length:        1})
					i++
				} else {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '<',
						StartPosition: i,
						Length:        1})
				}
				isFormulaSubPart = false
			case '>':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '≥',
						StartPosition: i,
						Length:        1})
					i++
				} else {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '>',
						StartPosition: i,
						Length:        1})
				}
				isFormulaSubPart = false
			case '!':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '≠',
						StartPosition: i,
						Length:        2})
					i++

					isFormulaSubPart = false
				} else {
					return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
				}
			case '&':
				if i+1 < runesLength && runes[i+1] == '&' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '&',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
				}
			case '|':
				if i+1 < runesLength && runes[i+1] == '|' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '|',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
				}
			case '=':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, token{Type: tt_OPERATION,
						Value:         '=',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
				}
			default:
				return nil, fmt.Errorf("invalid token '%s' detected at position '%d'", string(runes[i]), i)
			}
		}
	}
	return ret, nil
}

func (this tokenReader) isUnaryMinus(currentToken rune, tokens []token) bool {

	if currentToken == '-' {
		previousToken := tokens[len(tokens)-1]

		return !(previousToken.Type == tt_FLOATING_POINT ||
			previousToken.Type == tt_INTEGER ||
			previousToken.Type == tt_TEXT ||
			previousToken.Type == tt_RIGHT_BRACKET)
	} else {
		return false
	}
}

func (this tokenReader) isPartOfNumeric(character rune, isFirstCharacter bool, afterMinus bool, isFormulaSubPart bool) bool {
	return character == this.decimalSeparator || (character >= '0' && character <= '9') || (isFormulaSubPart && isFirstCharacter && character == '-') || (!isFirstCharacter && !afterMinus && character == 'e') || (!isFirstCharacter && character == 'E')
}

func (this tokenReader) isPartOfVariable(character rune, isFirstCharacter bool) bool {
	return (character == '$') || (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (!isFirstCharacter && character >= '0' && character <= '9') || (!isFirstCharacter && character == '_')
}

func (this tokenReader) isScientificNotation(char rune) bool {
	return char == 'e' || char == 'E'
}
