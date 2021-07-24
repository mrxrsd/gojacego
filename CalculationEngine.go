package gojacego

type CalculationEngine struct {
}

func NewCalculationEngine() *CalculationEngine {
	return &CalculationEngine{}
}

func (this *CalculationEngine) Calculate(formula string, vars map[string]float64) (float64, error) {

	return 0, nil
}
