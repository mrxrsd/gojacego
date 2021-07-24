package gojacego

import (
	memorycache "github.com/mrxrsd/gojacego/Cache"
)

type CalculationEngine struct {
	cache *memorycache.Memorycache
}

func NewCalculationEngine() *CalculationEngine {
	cache := memorycache.NewCache()
	return &CalculationEngine{
		cache: cache,
	}
}

func (this *CalculationEngine) Calculate(formula string, vars map[string]float64) (float64, error) {

	return 0, nil
}
