package gojacego

import (
	"errors"
	"math"
	"strings"

	"github.com/mrxrsd/gojacego/cache"
)

type JaceOptions struct {
	decimalSeparator rune
	caseSensitive    bool
	optimizeEnabled  bool
	defaultConstants bool
}

type CalculationEngine struct {
	cache            *cache.Memorycache
	options          *JaceOptions
	optimizer        IOptimizer
	executor         *Interpreter
	constantRegistry *ConstantRegistry
}

func NewCalculationEngine(options *JaceOptions) *CalculationEngine {
	cache := cache.NewCache()

	if options == nil {
		options = &JaceOptions{
			decimalSeparator: '.',
			caseSensitive:    false,
			optimizeEnabled:  true,
			defaultConstants: true,
		}
	}

	interpreter := &Interpreter{}
	optimizer := &Optimizer{executor: *interpreter}
	constantRegistry := NewConstantRegistry(options.caseSensitive)

	if options.defaultConstants {
		constantRegistry.RegisterConstant("e", math.E, false)
		constantRegistry.RegisterConstant("pi", math.Pi, false)
	}

	return &CalculationEngine{
		cache:            cache,
		options:          options,
		optimizer:        optimizer,
		executor:         interpreter,
		constantRegistry: constantRegistry,
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
		ret, err := this.executor.Execute(item.(Operation), formulaVariables)
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

	ret, err := this.executor.Execute(op, formulaVariables)
	if err != nil {
		return 0, nil
	}

	return ret, nil
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string) (Operation, error) {

	tokenReader := NewTokenReader(this.options.decimalSeparator)
	astBuilder := NewAstBuilder(this.options.caseSensitive, this.constantRegistry)

	tokens, err := tokenReader.Read(formula)
	if err != nil {
		return nil, err
	}

	operation, err := astBuilder.Build(tokens)
	if err != nil {
		return nil, err
	}

	if this.options.optimizeEnabled {
		optimizedOperation := this.optimizer.Optimize(operation)
		return optimizedOperation, nil
	}

	return operation, nil
}
