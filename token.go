package gojacego

/*
	Represents an input token
*/
type token struct {
	StartPosition int
	Length        int
	Type          tokenType
	Value         interface{}
}
