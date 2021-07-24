package gojacego

import (
	"testing"
)

func TestCalculationFormula1FloatingPoint(test *testing.T) {
	engine := NewCalculationEngine()
	result, _ := engine.Calculate("2.0+3.0", nil)

	if result != 5.0 {
		test.Errorf("exptected: 5.0, got: %f", result)
	}
}
