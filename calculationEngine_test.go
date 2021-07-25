package gojacego

import (
	"testing"
)

type CalculationTestScenario struct {
	formula        string
	variables      map[string]interface{}
	expectedResult float64
}

func TestCalculationFormula1FloatingPoint(test *testing.T) {
	engine := NewCalculationEngine(nil)
	result, _ := engine.Calculate("-(1+2+(3+4))", nil)

	if result != 5.0 {
		test.Errorf("exptected: 5.0, got: %f", result)
	}
}

func TestCalculationDefaultEngine(t *testing.T) {
	engine := NewCalculationEngine(nil)

	scenarios := []CalculationTestScenario{
		{
			formula:        "2.0+3.0",
			expectedResult: 5.0,
		},
		{
			formula:        "2+3",
			expectedResult: 5.0,
		},
		{
			formula:        "5 % 3.0",
			expectedResult: 2.0,
		},
		{
			formula:        "2^3.0",
			expectedResult: 8.0,
		},
		{
			formula: "var1*var2",
			variables: map[string]interface{}{
				"var1": 2.5,
				"var2": 3.4,
			},
			expectedResult: 8.5,
		},
		{
			formula: "vAr1*VaR2",
			variables: map[string]interface{}{
				"VaR1": 2.5,
				"vAr2": 3.4,
			},
			expectedResult: 8.5,
		},
		{
			formula:        "-100",
			expectedResult: -100.0,
		},
		{
			formula:        "5*-100",
			expectedResult: -500.0,
		},
		{
			formula:        "-(1+2+(3+4))",
			expectedResult: -10.0,
		},
		{
			formula:        "5+(-(1*2))",
			expectedResult: 3.0,
		},
		{
			formula:        "5*(-(1*2)*3)",
			expectedResult: -30.0,
		},
		{
			formula:        "5* -(1*2)",
			expectedResult: -10.0,
		},
		{
			formula:        "-(1*2)^3",
			expectedResult: -8.0,
		},
		{
			formula: "var1+2*(3*age)",
			variables: map[string]interface{}{
				"var1": 2,
				"age":  4,
			},
			expectedResult: 26.0,
		},
	}

	for _, test := range scenarios {
		result, err := engine.Calculate(test.formula, test.variables)
		if err != nil {
			t.Logf("Error: %s", err.Error())
		}

		if result != test.expectedResult {
			t.Logf("exptected: %f, got: %f", test.expectedResult, result)
			t.Fail()
		}
	}
}
