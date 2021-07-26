package gojacego

import (
	"math"
	"strings"
)

type ConstantRegistry struct {
	caseSensitive bool
	constants     map[string]constantInfo
}

type constantInfo struct {
	name           string
	value          float64
	isOverWritable bool
}

func NewConstantRegistry(caseSensitive bool) *ConstantRegistry {
	return &ConstantRegistry{
		caseSensitive: caseSensitive,
		constants:     map[string]constantInfo{},
	}
}

func (this *ConstantRegistry) Get(name string) (float64, bool) {
	if item, found := this.constants[this.convertConstantName(name)]; found {
		return item.value, true
	}
	return 0, false
}

func (this *ConstantRegistry) RegisterConstant(name string, value float64, isOverWritable bool) {
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

func (this *ConstantRegistry) convertConstantName(name string) string {
	if this.caseSensitive {
		return name
	}

	return strings.ToLower(name)
}

func RegistryDefaultConstants(registry *ConstantRegistry) {
	registry.RegisterConstant("e", math.E, false)
	registry.RegisterConstant("pi", math.Pi, false)
}
