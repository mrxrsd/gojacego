package gojacego

func toFloat64(value interface{}) float64 {
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

	panic("cannot convert to float64")
}
