package gojacego

import (
	"errors"
	"strings"
)

type variableCollection interface {
	Get(name string) (float64, error)
}

type formulaVariables map[string]float64

func createFormulaVariables(vars map[string]interface{}, caseSensitive bool) (formulaVariables, error) {

	ret := make(map[string]float64, len(vars))

	for k, v := range vars {
		name := k
		if caseSensitive {
			name = strings.ToLower(k)
		}
		if retFloat, err := toFloat64(v); err == nil {
			ret[name] = retFloat
		} else {
			return nil, err
		}

	}

	return ret, nil
}

func (p formulaVariables) Get(name string) (float64, error) {

	value, found := p[name]

	if !found {
		return 0, errors.New("Variable '" + name + "' not found.")
	}

	return value, nil
}
