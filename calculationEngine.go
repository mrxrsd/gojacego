package gojacego

import (
	"errors"
	"fmt"
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

func NewCalculationEngine() *CalculationEngine {
	return NewCalculationEngineWithOptions(&JaceOptions{
		decimalSeparator:  '.',
		argumentSeparador: ',',
		caseSensitive:     false,
		optimizeEnabled:   true,
		defaultConstants:  true,
		defaultFunctions:  true,
	})
}

func NewCalculationEngineWithOptions(options *JaceOptions) *CalculationEngine {
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

func (this *CalculationEngine) Calculate(formulaText string, vars map[string]interface{}) (float64, error) {

	if len(strings.TrimSpace(formulaText)) == 0 {
		return 0, errors.New("the parameter 'formula' is required")
	}

	formulaVariables := CreateFormulaVariables(vars, this.options.caseSensitive)

	item, found := this.cache.Get(formulaText)

	if found {
		formula := item.(Formula)
		return formula(formulaVariables), nil
	}

	op, err := this.buildAbstractSyntaxTree(formulaText, nil)
	if err != nil {
		return 0, err
	}

	formula := this.buildFormula(formulaText, nil, op)

	return formula(formulaVariables), nil
}

func (this *CalculationEngine) generateFormulaCacheKey(formulaText string, compiledConstantsRegistry *ConstantRegistry) string {
	if compiledConstantsRegistry != nil {
		var data []byte

		data = append(data, formulaText...)
		for k, p := range compiledConstantsRegistry.constants {
			data = append(data, "@"...)
			data = append(data, k...)
			data = append(data, ":"...)
			data = append(data, (fmt.Sprint(p.value))...)
			data = append(data, "@"...)
		}
		return string(data)
	}
	return formulaText
}

func (this *CalculationEngine) getFormula(formulaText string) Formula {

	item, found := this.cache.Get(formulaText)
	if found {
		return item.(Formula)
	}
	return nil
}

func (this *CalculationEngine) buildFormula(formulaText string, compiledConstants *ConstantRegistry, operation Operation) Formula {
	key := this.generateFormulaCacheKey(formulaText, compiledConstants)
	formula := this.executor.BuildFormula(operation, this.functionRegistry, this.constantRegistry)
	this.cache.Add(key, formula)

	return formula

}

func (this *CalculationEngine) Build(formulaText string) (Formula, error) {
	return this.BuildWithConstants(formulaText, nil)
}

func (this *CalculationEngine) BuildWithConstants(formulaText string, vars map[string]interface{}) (Formula, error) {

	if len(strings.TrimSpace(formulaText)) == 0 {
		return nil, errors.New("the parameter 'formula' is required")
	}

	compiledConstantsRegistry := NewConstantRegistry(this.options.caseSensitive)

	for k, p := range vars {
		compiledConstantsRegistry.RegisterConstant(k, ToFloat64(p), true)
	}

	item, found := this.cache.Get(this.generateFormulaCacheKey(formulaText, compiledConstantsRegistry))

	if found {
		return item.(Formula), nil
	}

	op, err := this.buildAbstractSyntaxTree(formulaText, compiledConstantsRegistry)
	if err != nil {
		return nil, err
	}

	return this.buildFormula(formulaText, compiledConstantsRegistry, op), nil
}

func (this *CalculationEngine) AddFunction(name string, body Delegate, isIdempotent bool) {
	this.functionRegistry.RegisterFunction(name, body, true, isIdempotent)
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string, compiledConstants *ConstantRegistry) (Operation, error) {

	tokenReader := NewTokenReader(this.options.decimalSeparator, this.options.argumentSeparador)
	astBuilder := NewAstBuilder(this.options.caseSensitive, this.functionRegistry, this.constantRegistry, compiledConstants)

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
