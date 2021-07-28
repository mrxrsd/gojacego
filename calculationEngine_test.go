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

func TestDebug(test *testing.T) {
	engine := NewCalculationEngine()
	result, _ := engine.Calculate("2*pi", nil)

	if result != 1 {

	}

}

func TestCalculationDefaultEngine(t *testing.T) {
	engine := NewCalculationEngine()

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
		{
			formula: "requests_made > requests_succeeded",
			variables: map[string]interface{}{
				"requests_made":      99.0,
				"requests_succeeded": 90.0,
			},
			expectedResult: 1.0,
		},
		{
			formula: "(requests_made * requests_succeeded / 100) >= 90",
			variables: map[string]interface{}{
				"requests_made":      99.0,
				"requests_succeeded": 90.0,
			},
			expectedResult: 0.0,
		},
		{
			formula:        "(0 || 0)",
			expectedResult: 0.0,
		},
		{
			formula:        "(1 || 0)",
			expectedResult: 1.0,
		},
		{
			formula:        "(1 && 0)",
			expectedResult: 0.0,
		},
		{
			formula:        "(1 && 1)",
			expectedResult: 1.0,
		},
		{
			formula: "var_var_1 + var_var_2",
			variables: map[string]interface{}{
				"var_var_1": 1.0,
				"var_var_2": 2.0,
			},
			expectedResult: 3.0,
		},
		{
			formula: "var1 == 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 1.0,
		},
		{
			formula: "var1 != 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 0.0,
		},
		{
			formula: "var1 > 2",
			variables: map[string]interface{}{
				"var1": 7,
			},
			expectedResult: 1.0,
		},
		{
			formula: "var1 > 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 0.0,
		},
		{
			formula: "var1 >= 2",
			variables: map[string]interface{}{
				"var1": 7,
			},
			expectedResult: 1.0,
		},
		{
			formula: "var1 >= 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 1.0,
		},
		{
			formula: "var1 >= 2",
			variables: map[string]interface{}{
				"var1": -2,
			},
			expectedResult: 0.0,
		},
		{
			formula: "var1 < 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 0.0,
		},
		{
			formula: "var1 <= 2",
			variables: map[string]interface{}{
				"var1": 2,
			},
			expectedResult: 1.0,
		},
		{
			formula: "var1 <= 2",
			variables: map[string]interface{}{
				"var1": 1,
			},
			expectedResult: 1.0,
		},
		{
			formula:        "1+2-3*4/5+6-7*8/9+0",
			expectedResult: 0.378,
			fnCallback: func(f float64) float64 {
				return math.Round(f*1000) / 1000
			},
		},
		{
			formula: "$var1 + 2",
			variables: map[string]interface{}{
				"$var1": 1,
			},
			expectedResult: 3.0,
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
	engine := NewCalculationEngine()

	engine.AddFunction("addTwo", func(arguments ...float64) (float64, error) {
		return arguments[0] + 2, nil
	}, true)

	result, _ := engine.Calculate("addTwo(2)", nil)

	if result != 4 {
		test.Errorf("exptected: 4.0, got: %f", result)
	}
}

func TestCompiledConstants(test *testing.T) {
	engine := NewCalculationEngine()

	constants := map[string]interface{}{
		"a": 1.0,
	}
	var fn, _ = engine.BuildWithConstants("a+b+c", constants)

	input := map[string]float64{
		"b": 2.0,
		"c": 3.0,
	}
	result := fn(input)

	if result != 6 {
		test.Errorf("exptected: 6.0, got: %f", result)
	}
}

func TestCaseUnsensitive(test *testing.T) {
	engine := NewCalculationEngineWithOptions(&JaceOptions{
		decimalSeparator:  '.',
		argumentSeparador: ',',
		caseSensitive:     false,
		optimizeEnabled:   true,
		defaultConstants:  true,
		defaultFunctions:  true,
	})

	engine.AddFunction("addTwo", func(arguments ...float64) (float64, error) {
		return arguments[0] + 2, nil
	}, true)

	resultFn, _ := engine.Calculate("addtwo(0)", nil)
	if resultFn != 2 {
		test.Errorf("exptected: 2 got: %f", resultFn)
	}

	resultPi, _ := engine.Calculate("PI", nil)

	if resultPi != math.Pi {
		test.Errorf("exptected: %f, got: %f", math.Pi, resultPi)
	}

	resultPilo, _ := engine.Calculate("pi", nil)

	if resultPilo != math.Pi {
		test.Errorf("exptected: %f, got: %f", math.Pi, resultPilo)
	}

	resultE, _ := engine.Calculate("E", nil)
	if resultE != math.E {
		test.Errorf("exptected: %f, got: %f", math.E, resultE)
	}

	resultElo, _ := engine.Calculate("e", nil)
	if resultElo != math.E {
		test.Errorf("exptected: %f, got: %f", math.E, resultElo)
	}

	vars := map[string]interface{}{
		"var1": 1,
		"var2": 2,
	}

	result, _ := engine.Calculate("vAr1 + VaR2", vars)
	if result != 3.0 {
		test.Errorf("exptected: 3.0, got: %f", result)
	}
}
