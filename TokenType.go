package gojacego

type TokenType int

const (
	INTEGER TokenType = iota
	FLOATING_POINT
	TEXT
	OPERATION
	LEFT_BRACKET
	RIGHT_BRACKET
	ARGUMENT_SEPARATOR
)
