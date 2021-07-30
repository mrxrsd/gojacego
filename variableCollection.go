package gojacego

import (
	"errors"
	"strings"
)

type variableCollection interface {
	Get(name string) (float64, error)
}

type formulaVariables map[string]float64

func createFormulaVariables(vars map[string]interface{}, caseSensitive bool) formulaVariables {

	ret := make(map[string]float64, len(vars))

	for k, v := range vars {
		name := k
		if caseSensitive {
			name = strings.ToLower(k)
		}
		ret[name] = toFloat64(v)
	}

	return ret
}

func (p formulaVariables) Get(name string) (float64, error) {

	value, found := p[name]

	if !found {
		return 0, errors.New("Variable '" + name + "' not found.")
	}

	return value, nil
}
