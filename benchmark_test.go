package gojacego

import (
	"testing"
)

func BenchmarkEvaluationSingle(bench *testing.B) {

	engine := NewCalculationEngine(nil)

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		engine.Calculate("1.0", nil)
	}
}
func BenchmarkCalculation(b *testing.B) {

	engine := NewCalculationEngine(nil)
	for i := 0; i < 10000000; i++ {
		engine.Calculate("2.0+3.0", nil)
	}

}
