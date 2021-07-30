package gojacego

import (
	"testing"
)

func createEngine() *CalculationEngine {
	return NewCalculationEngineWithOptions(JaceOptions{
		DecimalSeparator:  '.',
		ArgumentSeparador: ',',
		CaseSensitive:     true,
		OptimizeEnabled:   true,
		DefaultConstants:  false,
		DefaultFunctions:  false,
	})
}

func BenchmarkEvaluationNumericLiteral(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("(2) > (1)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

/*
  Benchmarks evaluation times of literals with modifiers
*/
func BenchmarkEvaluationLiteralModifiers(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("(2) + (2) == (4)")

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(nil)
	}
}

func BenchmarkEvaluationParameter(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("requests_made")
	parameters := map[string]float64{
		"requests_made": 99.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

/*
  Benchmarks evaluation times of parameters
*/
func BenchmarkEvaluationParameters(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("requests_made > requests_succeeded")
	parameters := map[string]float64{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

/*
  Benchmarks evaluation times of parameters + literals with modifiers
*/
func BenchmarkEvaluationParametersModifiers(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("(requests_made * requests_succeeded / 100) >= 90")
	parameters := map[string]float64{
		"requests_made":      99.0,
		"requests_succeeded": 90.0,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}

func BenchmarkExpr(bench *testing.B) {

	engine := createEngine()
	formula, _ := engine.Build("(Origin == 1 || Country == 55) && (Value >= 100 || Adults == 1)")
	parameters := map[string]float64{
		"Origin":  1,
		"Country": 55,
		"Value":   100,
		"Adults":  1,
	}

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		formula(parameters)
	}
}
