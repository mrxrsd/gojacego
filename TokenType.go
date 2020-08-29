package gojacego

/*
	The type of the token.
*/
type TokenType int

const (
	INTEGER = iota
	FLOATING_POINT
	TEXT
	OPERATION
	LEFT_BRACKET
	RIGHT_BRACKET
	ARGUMENT_SEPARATOR
)
