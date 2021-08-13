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

func getCalculationScenarios() []CalculationTestScenario {
	return []CalculationTestScenario{
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
		{
			formula:        "0-6==-6",
			expectedResult: 1.0,
		},
		{
			formula:        "32.9<-10",
			expectedResult: 0.0,
		},
		{
			formula:        "(2766237061056.00--1619439373.00)/17150000000.00",
			expectedResult: 161.39,
			fnCallback: func(f float64) float64 {
				return math.Round(f*100) / 100
			},
		},
		{
			formula:        "1e-3*5+2",
			expectedResult: 2.005,
		},
		{
			formula:        "5*-3",
			expectedResult: -15,
		},
		{
			formula: "$a + b + c_c + d1 + 10",
			variables: map[string]interface{}{
				"$a":  1,
				"b":   2,
				"c_c": 3,
				"d1":  4,
			},
			expectedResult: 20.0,
		},
	}
}

func TestCalculationDefaultEngine(t *testing.T) {
	engine, _ := NewCalculationEngine()
	scenarios := getCalculationScenarios()
	runScenarios(engine, runCalculate, scenarios, t)
}

func TestCalculationWithBuildingDefaultEngine(t *testing.T) {
	engine, _ := NewCalculationEngine()
	scenarios := getCalculationScenarios()
	runScenarios(engine, runCalculateWithBuilding, scenarios, t)
}

func TestStandardFunctions(t *testing.T) {
	engine, _ := NewCalculationEngine()

	scenarios := []CalculationTestScenario{
		{
			formula:        "if(2+2==$a, 10, 5)",
			expectedResult: 10,
			variables: map[string]interface{}{
				"$a": 4,
			},
		},
		{
			formula:        "if(2+2==a, 10, 5)",
			expectedResult: 5,
			variables: map[string]interface{}{
				"a": 8,
			},
		},
		{
			formula:        "sin(5)",
			expectedResult: math.Sin(5),
		},
		{
			formula:        "cos(5)",
			expectedResult: math.Cos(5),
		},
		{
			formula:        "asin(0.2)",
			expectedResult: math.Asin(0.2),
		},
		{
			formula:        "acos(0.2)",
			expectedResult: math.Acos(0.2),
		},
		{
			formula:        "tan(5)",
			expectedResult: math.Tan(5),
		},
		{
			formula:        "atan(5)",
			expectedResult: math.Atan(5),
		},
		{
			formula:        "log(5)",
			expectedResult: math.Log(5),
		},
		{
			formula:        "sqrt(25)",
			expectedResult: math.Sqrt(25),
		},
		{
			formula:        "trunc(1.234567)",
			expectedResult: math.Trunc(1.234567),
		},
		{
			formula:        "ceil(1.234567)",
			expectedResult: math.Ceil(1.234567),
		},
		{
			formula:        "floor(1.234567)",
			expectedResult: math.Floor(1.234567),
		},
		{
			formula:        "round(1.234567,2)",
			expectedResult: 1.23,
		},
		{
			formula:        "round(1.234567)",
			expectedResult: math.Round(1.234567),
		},
		{
			formula:        "random(5)",
			expectedResult: 1.0,
			fnCallback: func(f float64) float64 {
				if f >= 0.0 && f <= 1.0 {
					return 1.0
				} else {
					return 0.0
				}
			},
		},
		{
			formula:        "random(seed)",
			expectedResult: 1.0,
			variables: map[string]interface{}{
				"seed": 8,
			},
			fnCallback: func(f float64) float64 {
				if f >= 0.0 && f <= 1.0 {
					return 1.0
				} else {
					return 0.0
				}
			},
		},
		{
			formula:        "max(5,6,3,-4,99,67,45,34,-85)",
			expectedResult: 99,
		},
		{
			formula:        "min(5,6,3,-4,99,67,45,34,-85)",
			expectedResult: -85,
		},
		{
			formula:        "if(0.57 < (3000-500)/(1500-500), 10, 20)",
			expectedResult: 10,
		},
	}

	runScenarios(engine, runCalculate, scenarios, t)

}

func runCalculate(engine *CalculationEngine, formula string, vars map[string]interface{}) (float64, error) {
	return engine.Calculate(formula, vars)
}

func runCalculateWithBuilding(engine *CalculationEngine, formula string, vars map[string]interface{}) (float64, error) {

	fn, err := engine.Build(formula)
	if err != nil {
		return 0, err
	}

	varsFloat := map[string]interface{}{}

	for k, p := range vars {
		ret, _ := toFloat64(p)
		varsFloat[k] = ret
	}

	return fn(varsFloat)
}

type fnAction func(*CalculationEngine, string, map[string]interface{}) (float64, error)

func runScenarios(engine *CalculationEngine, fn fnAction, scenarios []CalculationTestScenario, t *testing.T) {

	for _, test := range scenarios {
		result, err := fn(engine, test.formula, test.variables)
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

func TestFormulaContext(test *testing.T) {

	engine, _ := NewCalculationEngine()

	engine.AddConstant("teste", 2.0, true)

	fn, _ := engine.Build("teste")

	result, _ := fn(nil)

	if result != 2.0 {
		test.Errorf("expected:2.0, got: %f", result)
	}

	engine.AddConstant("teste", 4.0, true)

	result2, _ := fn(nil)

	if result2 != 2.0 {
		test.Errorf("expected: 2.0, got: %f", result2)
	}

	fnAfter, _ := engine.Build("teste")

	resultAfter, _ := fnAfter(nil)

	if resultAfter != 4.0 {
		test.Errorf("expected: 4.0, got: %f", result2)
	}

	result3, _ := fn(nil)

	if result3 != 2.0 {
		test.Errorf("expected: 2.0, got: %f", result3)
	}

}

func TestCustomFunctions(test *testing.T) {
	engine, _ := NewCalculationEngine()

	engine.AddFunction("addTwo", func(arguments ...interface{}) float64 {
		return arguments[0].(float64) + 2
	}, true)

	engine.AddFunction("sum", func(arguments ...interface{}) float64 {
		return arguments[0].(float64) + arguments[1].(float64)
	}, true)

	result, _ := engine.Calculate("addTwo(2)", nil)

	if result != 4 {
		test.Errorf("expected: 4.0, got: %f", result)
	}

	resultSum, _ := engine.Calculate("sum(2,2)", nil)

	if resultSum != 4 {
		test.Errorf("expected: 4.0, got: %f", result)
	}

	resultSum1, _ := engine.Calculate("sum(2.4,2.4)", nil)

	if resultSum1 != 4.8 {
		test.Errorf("expected: 4.8, got: %f", resultSum1)
	}

}

func TestCompiledConstants(test *testing.T) {
	engine, _ := NewCalculationEngine()

	constants := map[string]interface{}{
		"a": 1.0,
	}
	var fn, _ = engine.BuildWithConstants("a+b+c", constants)

	input := map[string]interface{}{
		"b": 2.0,
		"c": 3.0,
	}
	result, _ := fn(input)

	if result != 6 {
		test.Errorf("expected: 6.0, got: %f", result)
	}
}

func TestCaseUnsensitive(test *testing.T) {

	engine, _ := NewCalculationEngine(WithDecimalSeparator('.'),
		WithArgumentSeparator(','),
		WithCaseSensitive(false),
		WithOptimizeEnabled(true),
		WithDefaultConstants(true),
		WithDefaultFunctions(true))

	engine.AddFunction("addTwo", func(arguments ...interface{}) float64 {
		return arguments[0].(float64) + 2
	}, true)

	resultFn, _ := engine.Calculate("addtwo(0)", nil)
	if resultFn != 2 {
		test.Errorf("expected: 2 got: %f", resultFn)
	}

	resultPi, _ := engine.Calculate("PI", nil)

	if resultPi != math.Pi {
		test.Errorf("expected: %f, got: %f", math.Pi, resultPi)
	}

	resultPilo, _ := engine.Calculate("pi", nil)

	if resultPilo != math.Pi {
		test.Errorf("expected: %f, got: %f", math.Pi, resultPilo)
	}

	resultE, _ := engine.Calculate("E", nil)
	if resultE != math.E {
		test.Errorf("expected: %f, got: %f", math.E, resultE)
	}

	resultElo, _ := engine.Calculate("e", nil)
	if resultElo != math.E {
		test.Errorf("expected: %f, got: %f", math.E, resultElo)
	}

	vars := map[string]interface{}{
		"var1": 1,
		"var2": 2,
	}

	result, _ := engine.Calculate("vAr1 + VaR2", vars)
	if result != 3.0 {
		test.Errorf("expected: 3.0, got: %f", result)
	}
}

func TestFormulaRequired(test *testing.T) {
	engine, _ := NewCalculationEngine()

	_, err := engine.Calculate("", nil)
	if err == nil {
		test.Errorf("error should not be null")
	}

	_, err2 := engine.Build("")
	if err2 == nil {
		test.Errorf("error should not be null")
	}
}

func TestCalculateFormulaVariableNotDefined(test *testing.T) {
	engine, _ := NewCalculationEngine()

	vars := map[string]interface{}{
		"var1": 1,
	}
	_, err := engine.Calculate("var1*var2", vars)
	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestCalculateParameterNotNumerical(test *testing.T) {
	engine, _ := NewCalculationEngine()

	vars := map[string]interface{}{
		"var1": "a",
	}
	_, err := engine.Calculate("var1", vars)
	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestFunctionVariableError(test *testing.T) {

	engine, _ := NewCalculationEngine(WithDecimalSeparator('.'),
		WithArgumentSeparator(','),
		WithCaseSensitive(false),
		WithOptimizeEnabled(true),
		WithDefaultConstants(true),
		WithDefaultFunctions(true))

	engine.AddFunction("addTwo", func(arguments ...interface{}) float64 {
		return arguments[0].(float64) + 2
	}, true)

	vars := map[string]interface{}{
		"x": "a",
	}

	_, err := engine.Calculate("addtwo(x)", vars)

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestFunctionRuntimeError(test *testing.T) {

	engine, _ := NewCalculationEngine(WithDecimalSeparator('.'),
		WithArgumentSeparator(','),
		WithCaseSensitive(false),
		WithOptimizeEnabled(true),
		WithDefaultConstants(true),
		WithDefaultFunctions(true))

	engine.AddFunction("addTwo", func(arguments ...interface{}) float64 {
		return arguments[1].(float64) + 2
	}, true)

	vars := map[string]interface{}{
		"x": 1,
	}

	_, err := engine.Calculate("addtwo(x)", vars)

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestGenerateCacheKey(test *testing.T) {
	engine, _ := NewCalculationEngine()

	key1 := engine.generateFormulaCacheKey("a+b+c", nil)
	if key1 != "a+b+c" {
		test.Errorf(" expected: 'a+b+c', got: %s", key1)
	}

	key2 := engine.generateFormulaCacheKey("a+2+c", nil)
	if key2 != "a+2+c" {
		test.Errorf(" expected: 'a+2+c', got: %s", key2)
	}

	registry := newConstantRegistry(false)
	registry.registerConstant("a", 1, true)
	registry.registerConstant("b", 2.5, true)

	key3 := engine.generateFormulaCacheKey("a+b+c", registry)
	if key3 != "a+b+c@a:1@b:2.5@" {
		test.Errorf(" expected: 'a+b+c@a:1@b:2.5@', got: %s", key3)
	}
}

func BenchmarkGenerateCacheKey(bench *testing.B) {
	engine, _ := NewCalculationEngine()

	registry := newConstantRegistry(false)
	registry.registerConstant("a", 1, true)
	registry.registerConstant("b", 2.5, true)

	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
		engine.generateFormulaCacheKey("a+b+c", registry)
	}
}
