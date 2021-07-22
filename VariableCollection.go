package gojacego

import "errors"

type VariableCollection interface {
	Get(name string) (interface{}, error)
}

type FormulaVariables map[string]interface{}

func (p FormulaVariables) Get(name string) (interface{}, error) {

	value, found := p[name]

	if !found {
		return nil, errors.New("Variable '" + name + "' not found.")
	}

	return value, nil
}
