package gojacego

import (
	"errors"
	"strconv"
)

type TokenReader struct {
	decimalSeparator  rune
	argumentSeparator rune
}

func NewTokenReader(decimalSeparator rune) *TokenReader {
	return &TokenReader{
		decimalSeparator: decimalSeparator,
	}
}

func (this TokenReader) Read(formula string) ([]Token, error) {
	ret := make([]Token, 0)

	if formula == "" {
		return nil, errors.New("formula cannot be empty")
	}

	runes := []rune(formula)
	runesLength := len(runes)
	isFormulaSubPart := true
	isScientific := false

	for i := 0; i < runesLength; i++ {
		if this.isPartOfNumeric(runes[i], true, isFormulaSubPart) {
			buffer := make([]rune, 0)
			buffer = append(buffer, runes[i])
			startPosition := i

			i++
			for i < runesLength {
				if !this.isPartOfNumeric(runes[i], false, isFormulaSubPart) {
					break
				}

				if isScientific && this.isScientificNotation(runes[i]) {
					isScientific = true
					if runes[i+1] == '-' {
						i++
						buffer = append(buffer, runes[i])
					}
				}

				buffer = append(buffer, runes[i])
				i++
			}

			if intVal, err := strconv.ParseInt(string(buffer), 10, 64); err == nil {
				ret = append(ret, Token{Type: INTEGER,
					Value:         intVal,
					StartPosition: startPosition,
					Length:        i - startPosition})
			} else {

				if floatVal, err := strconv.ParseFloat(string(buffer), 64); err == nil {
					ret = append(ret, Token{Type: FLOATING_POINT,
						Value:         floatVal,
						StartPosition: startPosition,
						Length:        i - startPosition})

					isScientific = false
					isFormulaSubPart = false
				} else if string(buffer) == "-" {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '_',
						StartPosition: startPosition,
						Length:        i - startPosition})
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

			ret = append(ret, Token{Type: TEXT,
				Value:         string(buffer),
				StartPosition: startPosition,
				Length:        i - startPosition})

			isFormulaSubPart = false

			if i == runesLength {
				continue
			}
		}

		if runes[i] == this.argumentSeparator {
			ret = append(ret, Token{Type: ARGUMENT_SEPARATOR,
				Value:         runes[i],
				StartPosition: i,
				Length:        1})
		} else {

			switch runes[i] {
			case ' ':
				continue
			case '+', '-', '*', '/', '^', '%', '≤', '≥', '≠':
				if this.isUnaryMinus(runes[i], ret) {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '_',
						StartPosition: i,
						Length:        1})
				} else {
					ret = append(ret, Token{Type: OPERATION,
						Value:         runes[i],
						StartPosition: i,
						Length:        1})

				}
				isFormulaSubPart = true
				break
			case '(':

				ret = append(ret, Token{Type: LEFT_BRACKET,
					Value:         runes[i],
					StartPosition: i,
					Length:        1})
				isFormulaSubPart = true
				break
			case ')':
				ret = append(ret, Token{Type: RIGHT_BRACKET,
					Value:         runes[i],
					StartPosition: i,
					Length:        1})
				isFormulaSubPart = false
				break
			case '<':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '≤',
						StartPosition: i,
						Length:        1})
				} else {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '<',
						StartPosition: i,
						Length:        1})
				}
				isFormulaSubPart = false
				break
			case '>':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '≥',
						StartPosition: i,
						Length:        1})
				} else {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '>',
						StartPosition: i,
						Length:        1})
				}
				isFormulaSubPart = false
				break
			case '!':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '≠',
						StartPosition: i,
						Length:        2})
					i++

					isFormulaSubPart = false
				} else {
					//TODO PARSE EXCEPTION
				}
				break
			case '&':
				if i+1 < runesLength && runes[i+1] == '&' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '&',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					//TODO PARSE EXCEPTION
				}
				break
			case '|':
				if i+1 < runesLength && runes[i+1] == '|' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '|',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					//TODO PARSE EXCEPTION
				}
				break
			case '=':
				if i+1 < runesLength && runes[i+1] == '=' {
					ret = append(ret, Token{Type: OPERATION,
						Value:         '=',
						StartPosition: i,
						Length:        2})
					i++
					isFormulaSubPart = false
				} else {
					//TODO PARSE EXCEPTION
				}
				break
			default:
				//TODO PARSE EXCEPTION
				//throw new ParseException(string.Format("Invalid token \"{0}\" detected at position {1}.", characters[i], i));
			}
		}
	}
	return ret, nil
}

func (this TokenReader) isUnaryMinus(currentToken rune, tokens []Token) bool {

	if currentToken == '-' {
		previousToken := tokens[len(tokens)-1]

		return !(previousToken.Type == FLOATING_POINT ||
			previousToken.Type == INTEGER ||
			previousToken.Type == TEXT ||
			previousToken.Type == RIGHT_BRACKET)
	} else {
		return false
	}
}

func (this TokenReader) isPartOfNumeric(character rune, isFirstCharacter bool, isFormulaSubPart bool) bool {
	return character == this.decimalSeparator || (character >= '0' && character <= '9') || (isFormulaSubPart && isFirstCharacter && character == '-') || (!isFirstCharacter && character == 'e') || (!isFirstCharacter && character == 'E')
}

func (this TokenReader) isPartOfVariable(character rune, isFirstCharacter bool) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (!isFirstCharacter && character >= '0' && character <= '9') || (!isFirstCharacter && character == '_')
}

func (this TokenReader) isScientificNotation(char rune) bool {
	return char == 'e' || char == 'E'
}
