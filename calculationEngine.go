package gojacego

import (
	"errors"
	"strings"

	"github.com/mrxrsd/gojacego/cache"
)

type JaceOptions struct {
	decimalSeparator rune
	caseSensitive    bool
	optimizeEnabled  bool
}

type CalculationEngine struct {
	cache     *cache.Memorycache
	options   *JaceOptions
	optimizer IOptimizer
	executor  *Interpreter
}

func NewCalculationEngine(options *JaceOptions) *CalculationEngine {
	cache := cache.NewCache()

	if options == nil {
		options = &JaceOptions{
			decimalSeparator: '.',
			caseSensitive:    false,
			optimizeEnabled:  true,
		}
	}

	interpreter := &Interpreter{}
	optimizer := &Optimizer{executor: *interpreter}

	return &CalculationEngine{
		cache:     cache,
		options:   options,
		optimizer: optimizer,
		executor:  interpreter,
	}
}

func (this *CalculationEngine) Calculate(formula string, vars map[string]interface{}) (float64, error) {

	if len(strings.TrimSpace(formula)) == 0 {
		return 0, errors.New("the parameter 'formula' is requred")
	}

	trimmedFormula := strings.TrimSpace(formula)
	item, found := this.cache.Get(trimmedFormula)

	if found {
		ret, err := this.executor.Execute(item.(Operation), vars)
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

	ret, err := this.executor.Execute(op, vars)
	if err != nil {
		return 0, nil
	}

	return ret, nil
}

func (this *CalculationEngine) buildAbstractSyntaxTree(formula string) (Operation, error) {

	tokenReader := NewTokenReader(this.options.decimalSeparator)
	astBuilder := NewAstBuilder(this.options.caseSensitive)

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
