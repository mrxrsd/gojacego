package gojacego

import "errors"

func toFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case int8:
		return float64(value.(int8)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case int:
		return float64(value.(int)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return float64(value.(float64)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	}
	return 0, errors.New("cannot convert parameter to float64")
}

func toFloat64Panic(value interface{}) float64 {
	switch value.(type) {
	case int8:
		return float64(value.(int8))
	case int16:
		return float64(value.(int16))
	case int32:
		return float64(value.(int32))
	case int64:
		return float64(value.(int64))
	case int:
		return float64(value.(int))
	case float32:
		return float64(value.(float32))
	case float64:
		return float64(value.(float64))
	case uint8:
		return float64(value.(uint8))
	case uint16:
		return float64(value.(uint16))
	case uint32:
		return float64(value.(uint32))
	case uint64:
		return float64(value.(uint64))
	}

	panic("cannot convert parameter to float64")
}
