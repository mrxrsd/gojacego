package gojacego

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/mrxrsd/gojacego/cache"
)

type jaceOptions struct {
	decimalSeparator  *rune
	argumentSeparator *rune
	caseSensitive     *bool
	optimizeEnabled   *bool
	defaultConstants  *bool
	defaultFunctions  *bool
}

type JaceOptions interface {
	apply(*jaceOptions) error
}

type applyOptions struct {
	f func(*jaceOptions) error
}

func (apply *applyOptions) apply(opts *jaceOptions) error {
	return apply.f(opts)
}

/*
	Decimal separator is a symbol used to separate the integer part from the fractional part
	of a number written in decimal form (i.e. '12.45').

	This value should be '.' or ','.
*/
func WithDecimalSeparator(decimalSeparator rune) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			if decimalSeparator != '.' && decimalSeparator != ',' {
				return errors.New("decimal separator should be equals '.' or ','")
			}
			options.decimalSeparator = &decimalSeparator
			return nil
		},
	}
}

/*
	Argument separator is a symbol used to separate the variables of a function (i.e. 'foo(a,b)').

	This value should be ',' or ';'.
*/
func WithArgumentSeparator(argumentSeparator rune) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			if argumentSeparator != ';' && argumentSeparator != ',' {
				return errors.New("argument separator should be equals ';' or ','")
			}
			options.argumentSeparator = &argumentSeparator
			return nil
		},
	}
}

/*
	CaseSensitive defines whether uppercase and lowercase letters are treated as distinct  or equivalent.
*/
func WithCaseSensitive(enabled bool) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			options.caseSensitive = &enabled
			return nil
		},
	}
}

/*
	Enable or disable optimizing of formulas.
*/
func WithOptimizeEnabled(enabled bool) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			options.optimizeEnabled = &enabled
			return nil
		},
	}
}

/*
	Enable or disable the default constants.
*/
func WithDefaultConstants(enabled bool) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			options.defaultConstants = &enabled
			return nil
		},
	}
}

/*
	Enable or disable the default functions.
*/
func WithDefaultFunctions(enabled bool) JaceOptions {
	return &applyOptions{
		f: func(options *jaceOptions) error {
			options.defaultFunctions = &enabled
			return nil
		},
	}
}

/*
	CalculationEngine represents the context of your evaluation engine.
*/
type CalculationEngine struct {
	cache            *cache.Memorycache
	options          *jaceOptions
	optimizer        *optimizer
	executor         *interpreter
	constantRegistry *constantRegistry
	functionRegistry *functionRegistry
}

func buildOptions(options []JaceOptions) (*jaceOptions, error) {

	var opts jaceOptions
	for _, opt := range options {
		err := opt.apply(&opts)
		if err != nil {
			return nil, err
		}
	}

	decimalSeparatorDefault := '.'
	argumentSeparatorDefault := ','
	caseSensitiveDefault := false
	optimizeEnabledDefault := true
	defaultConstantsDefault := true

	if opts.decimalSeparator == nil {
		opts.decimalSeparator = &decimalSeparatorDefault
	}

	if opts.argumentSeparator == nil {
		opts.argumentSeparator = &argumentSeparatorDefault
	}

	if opts.caseSensitive == nil {
		opts.caseSensitive = &caseSensitiveDefault
	}

	if opts.optimizeEnabled == nil {
		opts.optimizeEnabled = &optimizeEnabledDefault
	}

	if opts.defaultConstants == nil {
		opts.defaultConstants = &defaultConstantsDefault
	}

	if opts.defaultFunctions == nil {
		opts.defaultFunctions = &defaultConstantsDefault
	}

	return &opts, nil
}

/*
	Create a new calculation engine with the given options.
*/
func NewCalculationEngine(options ...JaceOptions) (*CalculationEngine, error) {
	cache := cache.NewCache()

	opts, err := buildOptions(options)
	if err != nil {
		return nil, err
	}

	interpreter := &interpreter{}
	optimizer := &optimizer{executor: *interpreter}
	constantRegistry := newConstantRegistry(*opts.caseSensitive)
	functionRegistry := newFunctionRegistry(*opts.caseSensitive)

	if *opts.defaultConstants {
		registryDefaultConstants(constantRegistry)
	}

	if *opts.defaultFunctions {
		registryDefaultFunctions(functionRegistry)
	}

	return &CalculationEngine{
		cache:            cache,
		options:          opts,
		optimizer:        optimizer,
		executor:         interpreter,
		constantRegistry: constantRegistry,
		functionRegistry: functionRegistry,
	}, nil
}

/*
	Parse and calculate from the given [formulaText] string using the given variables [vars].
	Returns an error if the given expression has invalid syntax.
*/
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

/*
	Parse the expression from the given [formulaText] string and build a Formula.
	Returns an error if the given expression has invalid syntax.
*/
func (this *CalculationEngine) Build(formulaText string) (Formula, error) {
	return this.BuildWithConstants(formulaText, nil)
}

/*
	Parse the expression from the given [formulaText] string and build a Formula with the given constants.
	Returns an error if the given expression has invalid syntax.
*/
func (this *CalculationEngine) BuildWithConstants(formulaText string, vars map[string]interface{}) (Formula, error) {

	if len(strings.TrimSpace(formulaText)) == 0 {
		return nil, errors.New("the parameter 'formula' is required")
	}

	compiledConstantsRegistry := newConstantRegistry(*this.options.caseSensitive)

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

/*
	Add a custom constant to the calculation engine.
*/
func (this *CalculationEngine) AddConstant(name string, value interface{}, isOverwritable bool) {
	val, _ := toFloat64(value)
	this.constantRegistry.registerConstant(name, val, isOverwritable)
	this.cache.Invalidate()
}

/*
	Add a custom function to the calculation engine.
*/
func (this *CalculationEngine) AddFunction(name string, body Delegate, isIdempotent bool) {
	this.functionRegistry.registerFunction(name, body, true, isIdempotent)
	this.cache.Invalidate()
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string, compiledConstants *constantRegistry) (operation, error) {

	tokenReader := newTokenReader(*this.options.decimalSeparator, *this.options.argumentSeparator)
	astBuilder := newAstBuilder(*this.options.caseSensitive, this.functionRegistry, this.constantRegistry, compiledConstants)

	tokens, err := tokenReader.read(formula)
	if err != nil {
		return nil, err
	}

	operation, err := astBuilder.build(tokens)
	if err != nil {
		return nil, err
	}

	if *this.options.optimizeEnabled {
		optimizedOperation := this.optimizer.optimize(operation, this.functionRegistry, this.constantRegistry)
		return optimizedOperation, nil
	}

	return operation, nil
}
