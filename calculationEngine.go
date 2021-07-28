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

	op, err := this.buildAbstractSyntaxTree(formulaText)
	if err != nil {
		return 0, err
	}

	formula := this.buildFormula(formulaText, op)

	return formula(formulaVariables), nil
}

func (this *CalculationEngine) generateFormulaCacheKey(formulaText string) string {
	return formulaText
}

func (this *CalculationEngine) getFormula(formulaText string) Formula {

	item, found := this.cache.Get(formulaText)
	if found {
		return item.(Formula)
	}
	return nil
}

func (this *CalculationEngine) buildFormula(formulaText string, operation Operation) Formula {
	key := this.generateFormulaCacheKey(formulaText)
	formula := this.executor.BuildFormula(operation, this.functionRegistry, this.constantRegistry)
	this.cache.Add(key, formula)

	return formula

}

func (this *CalculationEngine) Build(formulaText string) (Formula, error) {

	if len(strings.TrimSpace(formulaText)) == 0 {
		return nil, errors.New("the parameter 'formula' is required")
	}

	item, found := this.cache.Get(formulaText)

	if found {
		return item.(Formula), nil
	}

	op, err := this.buildAbstractSyntaxTree(formulaText)
	if err != nil {
		return nil, err
	}

	return this.buildFormula(formulaText, op), nil
}

func (this *CalculationEngine) AddFunction(name string, body Delegate, isIdempotent bool) {
	this.functionRegistry.RegisterFunction(name, body, true, isIdempotent)
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
