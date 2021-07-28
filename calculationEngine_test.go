package gojacego

import (
	"math"
	"testing"
)

type CalculationTestScenario struct {
	formula        string
	variables      map[string]interface{}
	expectedResult float64
	fnCallback     func(float64) float64
}

func TestCalculationFormula1FloatingPoint(test *testing.T) {
	engine := NewCalculationEngine(nil)
	result, _ := engine.Calculate("sin(14)", nil)

	if result != math.Sin(14) {
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
		{
			formula:        "-(1*2)^2",
			expectedResult: -4.0,
		},
		{
			formula:        "2*pi",
			expectedResult: 2 * math.Pi,
		},
		{
			formula:        "2*pI",
			expectedResult: 2 * math.Pi,
		},
		{
			formula:        "1+2-3*4/5+6-7*8/9+0",
			expectedResult: 0.378,
			fnCallback: func(f float64) float64 {
				return math.Round(f*1000) / 1000
			},
		},
		{
			formula: "var1 < var2",
			variables: map[string]interface{}{
				"var1": 2,
				"var2": 4.2,
			},
			expectedResult: 1.0,
		},
		{
			formula: "v_var1 + v_var2",
			variables: map[string]interface{}{
				"v_var1": 1,
				"v_var2": 2.0,
			},
			expectedResult: 3.0,
		},
		{
			formula:        "sin(14)",
			expectedResult: math.Sin(14),
		},
		{
			formula:        "max(5,6,3,-4,5,3,7,8,13,100)",
			expectedResult: 100,
		},
		{
			formula:        "max(sin(67), cos(67))",
			expectedResult: -0.518,
			fnCallback: func(f float64) float64 {
				return math.Round(f*1000) / 1000
			},
		},
	}

	for _, test := range scenarios {
		result, err := engine.Calculate(test.formula, test.variables)
		if err != nil {
			t.Logf("test:%s => Error: %s", test.formula, err.Error())
		}

		if test.fnCallback != nil {
			result = test.fnCallback(result)
		}

		if result != test.expectedResult {
			t.Logf("test: %s => expected: %f, got: %f", test.formula, test.expectedResult, result)
			t.Fail()
		}
	}
}

func TestCustomFunctions(test *testing.T) {
	engine := NewCalculationEngine(nil)

	engine.AddFunction("addTwo", func(arguments ...float64) (float64, error) {
		return arguments[0] + 2, nil
	}, true)

	result, _ := engine.Calculate("addTwo(2)", nil)

	if result != 4 {
		test.Errorf("exptected: 4.0, got: %f", result)
	}
}
