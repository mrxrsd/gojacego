package gojacego

import (
	"math"
	"strings"
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

}
