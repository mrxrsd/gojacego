package gojacego

import (
	"math"
	"math/rand"
	"strings"
)

type Delegate func(arguments ...float64) (float64, error)

type functionRegistry struct {
	caseSensitive bool
	functions     map[string]functionInfo
}

type functionInfo struct {
	name           string
	function       Delegate
	isOverWritable bool
	isIdempotent   bool
}

func newFunctionRegistry(caseSensitive bool) *functionRegistry {
	return &functionRegistry{
		caseSensitive: caseSensitive,
		functions:     map[string]functionInfo{},
	}
}

func (this *functionRegistry) get(name string) (*functionInfo, bool) {
	if item, found := this.functions[this.convertFunctionName(name)]; found {
		return &item, true
	}
	return nil, false
}

func (this *functionRegistry) registerFunction(name string, function Delegate, isOverWritable bool, isIdempotent bool) {
	handledFunctionName := this.convertFunctionName(name)

	if item, found := this.functions[handledFunctionName]; found {
		if !item.isOverWritable {
			panic("the function '" + item.name + "' cannot be overwritten")
		}
	}

	functionInfo := &functionInfo{
		name:           handledFunctionName,
		function:       function,
		isOverWritable: isOverWritable,
		isIdempotent:   isIdempotent,
	}

	this.functions[handledFunctionName] = *functionInfo
}

func (this *functionRegistry) convertFunctionName(name string) string {
	if this.caseSensitive {
		return name
	}
	return strings.ToLower(name)
}

func registryDefaultFunctions(registry *functionRegistry) {

	registry.registerFunction("sin", func(arguments ...float64) (float64, error) {
		return math.Sin(arguments[0]), nil
	}, false, true)

	registry.registerFunction("cos", func(arguments ...float64) (float64, error) {
		return math.Cos(arguments[0]), nil
	}, false, true)

	registry.registerFunction("asin", func(arguments ...float64) (float64, error) {
		return math.Asin(arguments[0]), nil
	}, false, true)

	registry.registerFunction("acos", func(arguments ...float64) (float64, error) {
		return math.Acos(arguments[0]), nil
	}, false, true)

	registry.registerFunction("tan", func(arguments ...float64) (float64, error) {
		return math.Tan(arguments[0]), nil
	}, false, true)

	registry.registerFunction("atan", func(arguments ...float64) (float64, error) {
		return math.Atan(arguments[0]), nil
	}, false, true)

	registry.registerFunction("log", func(arguments ...float64) (float64, error) {
		return math.Log(arguments[0]), nil
	}, false, true)

	registry.registerFunction("sqrt", func(arguments ...float64) (float64, error) {
		return math.Sqrt(arguments[0]), nil
	}, false, true)

	registry.registerFunction("trunc", func(arguments ...float64) (float64, error) {
		return math.Trunc(arguments[0]), nil
	}, false, true)

	registry.registerFunction("ceil", func(arguments ...float64) (float64, error) {
		return math.Ceil(arguments[0]), nil
	}, false, true)

	registry.registerFunction("round", func(arguments ...float64) (float64, error) {
		if len(arguments) <= 1 {
			return math.Round(arguments[0]), nil
		} else {
			pow := math.Pow(10, arguments[1])
			return math.Round(arguments[0]*pow) / pow, nil
		}
	}, false, true)

	registry.registerFunction("random", func(arguments ...float64) (float64, error) {
		rand.Seed(int64(arguments[0]))
		return rand.Float64(), nil
	}, false, false)

	registry.registerFunction("floor", func(arguments ...float64) (float64, error) {
		return math.Floor(arguments[0]), nil
	}, false, true)

	registry.registerFunction("max", func(arguments ...float64) (float64, error) {
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

	registry.registerFunction("min", func(arguments ...float64) (float64, error) {
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

	registry.registerFunction("if", func(arguments ...float64) (float64, error) {
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
