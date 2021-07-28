package gojacego

import (
	"testing"
)

func BenchmarkEvaluationNumericLiteral(bench *testing.B) {

	engine := NewCalculationEngine(nil)
	formula, _ := engine.Build("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

func BenchmarkCalculation(b *testing.B) {

	engine := NewCalculationEngine(nil)
	formula, _ := engine.Build("2.0+3.0")

	for i := 0; i < 10000000; i++ {
		formula(nil)
	}

}
