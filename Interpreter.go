package gojacego

import (
	"errors"
	"reflect"
)

func Execute(op Operation) (float64, error) {

	if op == nil {
		return 0, errors.New("Operation cannot be nil.")
	}

	if reflect.TypeOf(op).Name() == "IntegerConstant" {

	}
	return 1.0, nil
}
