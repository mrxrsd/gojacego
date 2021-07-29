package gojacego

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

type Delegate func(arguments ...float64) (float64, error)

type FunctionRegistry struct {
	caseSensitive bool
	functions     map[string]FunctionInfo
}

type FunctionInfo struct {
	name           string
	function       Delegate
	isOverWritable bool
	isIdempotent   bool
}

func NewFunctionRegistry(caseSensitive bool) *FunctionRegistry {
	return &FunctionRegistry{
		caseSensitive: caseSensitive,
		functions:     map[string]FunctionInfo{},
	}
}

func (this *FunctionRegistry) Get(name string) (*FunctionInfo, bool) {
	if item, found := this.functions[this.convertFunctionName(name)]; found {
		return &item, true
	}
	return nil, false
}

func (this *FunctionRegistry) RegisterFunction(name string, function Delegate, isOverWritable bool, isIdempotent bool) {
	handledFunctionName := this.convertFunctionName(name)

	if item, found := this.functions[handledFunctionName]; found {
		if !item.isOverWritable {
			panic("the function '" + item.name + "' cannot be overwritten")
		}
	}

	functionInfo := &FunctionInfo{
		name:           handledFunctionName,
		function:       function,
		isOverWritable: isOverWritable,
		isIdempotent:   isIdempotent,
	}

	this.functions[handledFunctionName] = *functionInfo
}

func (this *FunctionRegistry) convertFunctionName(name string) string {
	if this.caseSensitive {
		return name
	}
	return strings.ToLower(name)
}

func RegistryDefaultFunctions(registry *FunctionRegistry) {

	registry.RegisterFunction("sin", func(arguments ...float64) (float64, error) {
		return math.Sin(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("cos", func(arguments ...float64) (float64, error) {
		return math.Cos(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("asin", func(arguments ...float64) (float64, error) {
		return math.Asin(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("acos", func(arguments ...float64) (float64, error) {
		return math.Acos(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("tan", func(arguments ...float64) (float64, error) {
		return math.Tan(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("atan", func(arguments ...float64) (float64, error) {
		return math.Atan(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("log", func(arguments ...float64) (float64, error) {
		return math.Log(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("sqrt", func(arguments ...float64) (float64, error) {
		return math.Sqrt(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("trunc", func(arguments ...float64) (float64, error) {
		return math.Trunc(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("ceil", func(arguments ...float64) (float64, error) {
		return math.Ceil(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("round", func(arguments ...float64) (float64, error) {
		if len(arguments) <= 1 {
			return math.Round(arguments[0]), nil
		} else {
			pow := math.Pow(10, arguments[1])
			return math.Round(arguments[0]*pow) / pow, nil
		}
	}, false, true)

	registry.RegisterFunction("random", func(arguments ...float64) (float64, error) {
		rand.Seed(time.Now().UnixNano())
		return rand.Float64(), nil
	}, false, false)

	registry.RegisterFunction("floor", func(arguments ...float64) (float64, error) {
		return math.Floor(arguments[0]), nil
	}, false, true)

	registry.RegisterFunction("max", func(arguments ...float64) (float64, error) {
		if len(arguments) > 0 {
			max := arguments[0]
			for _, v := range arguments {
				if v > max {
					max = v
				}
			}
			return max, nil
		} else {
			return 0, nil
		}
	}, false, true)

	registry.RegisterFunction("min", func(arguments ...float64) (float64, error) {
		if len(arguments) > 0 {
			min := arguments[0]
			for _, v := range arguments {
				if v < min {
					min = v
				}
			}
			return min, nil
		} else {
			return 0, nil
		}
	}, false, true)

	registry.RegisterFunction("if", func(arguments ...float64) (float64, error) {
		if len(arguments) == 3 {
			if arguments[0] != 0.0 {
				return arguments[1], nil
			} else {
				return arguments[2], nil
			}

		} else {
			return 0, nil
		}
	}, false, true)

}
