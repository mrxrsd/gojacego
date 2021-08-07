package gojacego

import (
	"errors"
)

type variableCollection interface {
	Get(name string) (interface{}, error)
}

type formulaVariables map[string]interface{}

func (p formulaVariables) Get(name string) (interface{}, error) {

	value, found := p[name]

	if !found {
		return 0, errors.New("Variable '" + name + "' not found.")
	}

	return value, nil
}
