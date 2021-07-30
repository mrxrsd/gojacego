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
	options          JaceOptions
	optimizer        *optimizer
	executor         *interpreter
	constantRegistry *constantRegistry
	functionRegistry *functionRegistry
}

func NewCalculationEngine() *CalculationEngine {
	return NewCalculationEngineWithOptions(JaceOptions{
		decimalSeparator:  '.',
		argumentSeparador: ',',
		caseSensitive:     false,
		optimizeEnabled:   true,
		defaultConstants:  true,
		defaultFunctions:  true,
	})
}

func NewCalculationEngineWithOptions(options JaceOptions) *CalculationEngine {
	cache := cache.NewCache()

	if options == (JaceOptions{}) {
		options = JaceOptions{
			decimalSeparator:  '.',
			argumentSeparador: ',',
			caseSensitive:     false,
			optimizeEnabled:   true,
			defaultConstants:  true,
			defaultFunctions:  true,
		}
	}

	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	constantRegistry := newConstantRegistry(options.caseSensitive)
	functionRegistry := newFunctionRegistry(options.caseSensitive)

	if options.defaultConstants {
		registryDefaultConstants(constantRegistry)
	}

	if options.defaultFunctions {
		registryDefaultFunctions(functionRegistry)
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

	formulaVariables := createFormulaVariables(vars, this.options.caseSensitive)

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

func (this *CalculationEngine) generateFormulaCacheKey(formulaText string, compiledConstantsRegistry *constantRegistry) string {
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

func (this *CalculationEngine) buildFormula(formulaText string, compiledConstants *constantRegistry, operation operation) Formula {
	key := this.generateFormulaCacheKey(formulaText, compiledConstants)
	formula := this.executor.buildFormula(operation, this.functionRegistry, this.constantRegistry)
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

	compiledConstantsRegistry := newConstantRegistry(this.options.caseSensitive)

	for k, p := range vars {
		compiledConstantsRegistry.registerConstant(k, toFloat64(p), true)
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
	this.functionRegistry.registerFunction(name, body, true, isIdempotent)
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string, compiledConstants *constantRegistry) (operation, error) {

	tokenReader := newTokenReader(this.options.decimalSeparator, this.options.argumentSeparador)
	astBuilder := newAstBuilder(this.options.caseSensitive, this.functionRegistry, this.constantRegistry, compiledConstants)

	tokens, err := tokenReader.read(formula)
	if err != nil {
		return nil, err
	}

	operation, err := astBuilder.build(tokens)
	if err != nil {
		return nil, err
	}

	if this.options.optimizeEnabled {
		optimizedOperation := this.optimizer.optimize(operation, this.functionRegistry, this.constantRegistry)
		return optimizedOperation, nil
	}

	return operation, nil
}
