package gojacego

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/mrxrsd/gojacego/cache"
)

type JaceOptions struct {
	DecimalSeparator  rune
	ArgumentSeparador rune
	CaseSensitive     bool
	OptimizeEnabled   bool
	DefaultConstants  bool
	DefaultFunctions  bool
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
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     false,
		OptimizeEnabled:   true,
		DefaultConstants:  true,
		DefaultFunctions:  true,
	})
}

func NewCalculationEngineWithOptions(options JaceOptions) *CalculationEngine {
	cache := cache.NewCache()

	if options == (JaceOptions{}) {
		options = JaceOptions{
			DecimalSeparator:  '.',
			ArgumentSeparador: ',',
			CaseSensitive:     false,
			OptimizeEnabled:   true,
			DefaultConstants:  true,
			DefaultFunctions:  true,
		}
	}

	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	constantRegistry := newConstantRegistry(options.CaseSensitive)
	functionRegistry := newFunctionRegistry(options.CaseSensitive)

	if options.DefaultConstants {
		registryDefaultConstants(constantRegistry)
	}

	if options.DefaultFunctions {
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

	key := this.generateFormulaCacheKey(formulaText, nil)

	item, found := this.cache.Get(key)

	if found {
		formula := item.(Formula)
		return formula(vars)
	}

	op, err := this.buildAbstractSyntaxTree(formulaText, nil)
	if err != nil {
		return 0, err
	}

	formula := this.buildFormula(formulaText, nil, op)

	this.cache.Add(key, formula)

	return formula(vars)
}

func (this *CalculationEngine) generateFormulaCacheKey(formulaText string, compiledConstantsRegistry *constantRegistry) string {
	if compiledConstantsRegistry != nil {
		var data []byte
		var keys []string
		for k := range compiledConstantsRegistry.constants {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		data = append(data, formulaText...)
		data = append(data, "@"...)
		for _, k := range keys {
			data = append(data, k...)
			data = append(data, ":"...)
			data = append(data, (fmt.Sprint(compiledConstantsRegistry.constants[k].value))...)
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
	return this.executor.buildFormula(operation, this.functionRegistry, this.constantRegistry)
}

func (this *CalculationEngine) Build(formulaText string) (Formula, error) {
	return this.BuildWithConstants(formulaText, nil)
}

func (this *CalculationEngine) BuildWithConstants(formulaText string, vars map[string]interface{}) (Formula, error) {

	if len(strings.TrimSpace(formulaText)) == 0 {
		return nil, errors.New("the parameter 'formula' is required")
	}

	compiledConstantsRegistry := newConstantRegistry(this.options.CaseSensitive)

	for k, p := range vars {
		retFloat, err := toFloat64(p)
		if err != nil {
			return nil, fmt.Errorf("the variable '%s' cannot be converted to float", k)
		}
		compiledConstantsRegistry.registerConstant(k, retFloat, true)
	}

	key := this.generateFormulaCacheKey(formulaText, compiledConstantsRegistry)

	item, found := this.cache.Get(key)

	if found {
		return item.(Formula), nil
	}

	op, err := this.buildAbstractSyntaxTree(formulaText, compiledConstantsRegistry)
	if err != nil {
		return nil, err
	}

	formula := this.buildFormula(formulaText, compiledConstantsRegistry, op)

	this.cache.Add(key, formula)

	return formula, nil
}

func (this *CalculationEngine) AddConstant(name string, value float64, isOverwritable bool) {
	this.constantRegistry.registerConstant(name, value, isOverwritable)
	this.cache.Invalidate()
}

func (this *CalculationEngine) AddFunction(name string, body Delegate, isIdempotent bool) {
	this.functionRegistry.registerFunction(name, body, true, isIdempotent)
	this.cache.Invalidate()
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string, compiledConstants *constantRegistry) (operation, error) {

	tokenReader := newTokenReader(this.options.DecimalSeparator, this.options.ArgumentSeparador)
	astBuilder := newAstBuilder(this.options.CaseSensitive, this.functionRegistry, this.constantRegistry, compiledConstants)

	tokens, err := tokenReader.read(formula)
	if err != nil {
		return nil, err
	}

	operation, err := astBuilder.build(tokens)
	if err != nil {
		return nil, err
	}

	if this.options.OptimizeEnabled {
		optimizedOperation := this.optimizer.optimize(operation, this.functionRegistry, this.constantRegistry)
		return optimizedOperation, nil
	}

	return operation, nil
}
