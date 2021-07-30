package gojacego

import (
	"math"
	"strings"
)

type constantRegistry struct {
	caseSensitive bool
	constants     map[string]constantInfo
}

type constantInfo struct {
	name           string
	value          float64
	isOverWritable bool
}

func newConstantRegistry(caseSensitive bool) *constantRegistry {
	return &constantRegistry{
		caseSensitive: caseSensitive,
		constants:     map[string]constantInfo{},
	}
}

func (this *constantRegistry) get(name string) (float64, bool) {
	if item, found := this.constants[this.convertConstantName(name)]; found {
		return item.value, true
	}
	return 0, false
}

func (this *constantRegistry) registerConstant(name string, value float64, isOverWritable bool) {
	handledConstantName := this.convertConstantName(name)

	if item, found := this.constants[handledConstantName]; found {
		if !item.isOverWritable {
			panic("the constant '" + item.name + "' cannot be overwritten")
		}
	}

	constantInfo := &constantInfo{
		name:           handledConstantName,
		value:          value,
		isOverWritable: isOverWritable,
	}

	this.constants[handledConstantName] = *constantInfo
}

func (this *constantRegistry) convertConstantName(name string) string {
	if this.caseSensitive {
		return name
	}

	return strings.ToLower(name)
}

func registryDefaultConstants(registry *constantRegistry) {
	registry.registerConstant("e", math.E, false)
	registry.registerConstant("pi", math.Pi, false)
}
