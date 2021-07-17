package gojacego

type AstBuilder struct {
	caseSensitive bool
}

func NewAstBuilder(caseSensitive bool) *AstBuilder {
	return &AstBuilder{
		caseSensitive: caseSensitive,
	}
}

func (this AstBuilder) Build(tokens []Token) (*Expression, error) {

	return nil, nil
}
