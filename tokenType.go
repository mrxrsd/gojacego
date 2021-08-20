package gojacego

type tokenType int

const (
	tt_INTEGER tokenType = iota
	tt_FLOATING_POINT
	tt_TEXT
	tt_OPERATION
	tt_LEFT_BRACKET
	tt_RIGHT_BRACKET
	tt_ARGUMENT_SEPARATOR
	tt_LITERAL
	tt_DATE
)
