package gojacego

import (
	"testing"
)

func BenchmarkCalculation(b *testing.B) {

	engine := NewCalculationEngine(nil)
	for i := 0; i < 10000000; i++ {
		engine.Calculate("2.0+3.0", nil)
	}

}
