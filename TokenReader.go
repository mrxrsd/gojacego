package gojacego

import "errors"

type TokenReader struct {
	decimalSeparator rune
}

func NewTokenReader(decimalSeparator rune) *TokenReader {
	return &TokenReader{
		decimalSeparator: decimalSeparator,
	}
}

func (this TokenReader) Read(formula string) ([]Token, error) {
	ret := make([]Token, 0)

	if formula == "" {
		return nil, errors.New("formula cannot be empty.")
	}

	return ret, nil
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
