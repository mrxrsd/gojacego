package gojacego

/*
	Represents an input token
*/
type Token struct {
	StartPosition int
	Length        int
	Type          TokenType
	Value         interface{}
}
