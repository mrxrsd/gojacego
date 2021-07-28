package gojacego

import (
	"errors"
	"strings"

	"github.com/mrxrsd/gojacego/cache"
)

type JaceOptions struct {
	decimalSeparator  rune
	argumentSeparador rune
	caseSensitive     bool
	optimizeEnabled   bool
	defaultConstants  bool
	defaultFunctions  bool
}

type CalculationEngine struct {
	cache            *cache.Memorycache
	options          *JaceOptions
	optimizer        *Optimizer
	executor         *Interpreter
	constantRegistry *ConstantRegistry
	functionRegistry *FunctionRegistry
}

func NewCalculationEngine(options *JaceOptions) *CalculationEngine {
	cache := cache.NewCache()

	if options == nil {
		options = &JaceOptions{
			decimalSeparator:  '.',
			argumentSeparador: ',',
			caseSensitive:     false,
			optimizeEnabled:   true,
			defaultConstants:  true,
			defaultFunctions:  true,
		}
	}

	interpreter := &Interpreter{}
	optimizer := &Optimizer{executor: *interpreter}
	constantRegistry := NewConstantRegistry(options.caseSensitive)
	functionRegistry := NewFunctionRegistry(options.caseSensitive)

	if options.defaultConstants {
		RegistryDefaultConstants(constantRegistry)
	}

	if options.defaultFunctions {
		RegistryDefaultFunctions(functionRegistry)
	}

	return &CalculationEngine{
		cache:            cache,
		options:          options,
		optimizer:        optimizer,
		executor:         interpreter,
		constantRegistry: constantRegistry,
		functionRegistry: functionRegistry,
	}
}

func (this *CalculationEngine) Calculate(formula string, vars map[string]interface{}) (float64, error) {

	if len(strings.TrimSpace(formula)) == 0 {
		return 0, errors.New("the parameter 'formula' is required")
	}

	formulaVariables := CreateFormulaVariables(vars, this.options.caseSensitive)

	trimmedFormula := strings.TrimSpace(formula)
	item, found := this.cache.Get(trimmedFormula)

	if found {
		ret, err := this.executor.Execute(item.(Operation), formulaVariables, this.functionRegistry, this.constantRegistry)
		if err != nil {
			return 0, nil
		}
		return ret, nil
	}

	op, err := this.buildAbstractSyntaxTree(trimmedFormula)
	if err != nil {
		return 0, err
	}

	this.cache.Add(trimmedFormula, op)

	ret, err := this.executor.Execute(op, formulaVariables, this.functionRegistry, this.constantRegistry)
	if err != nil {
		return 0, nil
	}

	return ret, nil
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string) (Operation, error) {

	tokenReader := NewTokenReader(this.options.decimalSeparator, this.options.argumentSeparador)
	astBuilder := NewAstBuilder(this.options.caseSensitive, this.functionRegistry, this.constantRegistry)

	tokens, err := tokenReader.Read(formula)
	if err != nil {
		return nil, err
	}

	operation, err := astBuilder.Build(tokens)
	if err != nil {
		return nil, err
	}

	if this.options.optimizeEnabled {
		optimizedOperation := this.optimizer.Optimize(operation, this.functionRegistry, this.constantRegistry)
		return optimizedOperation, nil
	}

	return operation, nil
}
