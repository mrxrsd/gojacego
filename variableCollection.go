package gojacego

import (
	"errors"
	"strings"
)

type VariableCollection interface {
	Get(name string) (float64, error)
}

type FormulaVariables map[string]float64

func CreateFormulaVariables(vars map[string]interface{}, caseSensitive bool) FormulaVariables {

	ret := make(map[string]float64, len(vars))

	for k, v := range vars {
		name := k
		if caseSensitive {
			name = strings.ToLower(k)
		}
		ret[name] = ToFloat64(v)
	}

	return ret
}

func (p FormulaVariables) Get(name string) (float64, error) {

	value, found := p[name]

	if !found {
		return 0, errors.New("Variable '" + name + "' not found.")
	}

	return value, nil
}
