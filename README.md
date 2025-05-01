[![Build](https://github.com/mrxrsd/gojacego/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/mrxrsd/gojacego/actions/workflows/build.yml)
[![Test](https://github.com/mrxrsd/gojacego/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/mrxrsd/gojacego/actions/workflows/test.yml)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrxrsd/gojacego)](https://goreportcard.com/report/github.com/mrxrsd/gojacego)
[![Coverage Status](https://coveralls.io/repos/github/mrxrsd/gojacego/badge.svg)](https://coveralls.io/github/mrxrsd/gojacego)
[![Godoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=ff69b4)](https://godoc.org/github.com/mrxrsd/gojacego)

# goJACEgo

> 🚀 A lightning-fast mathematical expression engine for Go

goJACEgo is a high-performance calculation engine that brings the power of dynamic mathematical expressions to your Go applications. Built with pure Go, it delivers blazing-fast performance while maintaining simplicity and ease of use.

## Why goJACEgo?

✨ **Ultra-Fast Performance**: Consistently outperforms other expression evaluators in benchmarks
🔧 **Simple Integration**: Just a few lines of code to get started
🛡️ **Production-Ready**: Extensively tested with high code coverage
🎯 **Dynamic Expressions**: Evaluate mathematical formulas on the fly
⚡ **Variable Support**: Use dynamic variables in your expressions

## What Can You Build With It?

- 📊 Financial calculations and analytics
- 🎮 Game scoring and mechanics
- 📈 Real-time data processing
- 🧮 Dynamic business rules
- 🔬 Scientific computations

## 🌐 Under the Hood

Built on proven compiler design principles, goJACEgo's architecture ensures reliability and performance:

### 🔢 Smart Tokenization
Converts expressions into optimized tokens with low overhead

### 🌳 Intelligent AST Creation
Builds a smart hierarchical tree that represents your mathematical formulas with precision

### ⚡ Performance Optimization
Automatically optimizes the execution path for maximum speed

![Architecture Overview](https://github.com/mrxrsd/gojacego/blob/master/.github/imgs/1.png?raw=true)

[Learn more about the architecture](https://pieterderycke.wordpress.com/2012/11/04/jace-net-just-another-calculation-engine-for-net/)

## 🔥 Quick Start

Get up and running in seconds:

### Simple Calculation

```go
engine, _ := gojacego.NewCalculationEngine()

vars := map[string]interface{}{
   "price": 29.99,
   "quantity": 5
}

total, _ := engine.Calculate("price * quantity", vars)
// 149.95
```

### Advanced Usage
Build optimized formulas for repeated use:

```go
engine, _ := gojacego.NewCalculationEngine(
    gojacego.WithOptimizeEnabled(true),    // 🚀 Speed optimization
    gojacego.WithDefaultConstants(true),    // 📊 Built-in constants
    gojacego.WithDefaultFunctions(true),    // ⚙️ Standard functions
    gojacego.WithCaseSensitive(false)      // 🔧 Flexible syntax
)

formula := engine.Build("price * quantity * (1 - discount)")

result := formula(map[string]interface{}{
    "price": 99.99,
    "quantity": 3,
    "discount": 0.15
})
// 254.97
```

## ✨ Feature Highlights

### 📊 Core Mathematical Operations

The following mathematical operations are supported:
* Addition: +
* Subtraction: -
* Multiplication: *
* Division: /
* Modulo: %
* Exponentiation: ^

### Boolean Operations

The following boolean operations are supported:

* Less than: <
* Less than or equal: <=
* More than: >
* More than or equal: >=
* Equal: ==
* Not Equal: !=

The boolean operations map true to 1.0 and false to 0.0. All functions accepting a condition will consider 0.0 as false and any other value as true.

```go
result, _ := engine.Calculate("5 > 1", nil)
// 1.0
```
### Scientific Notation

```go
result, _ := engine.Calculate("1E-3*5+2", nil)
// 2.005
```

### 🔑 Flexible Variable Support

Use descriptive variable names that make sense for your business:
```go
vars := map[string]interface{}{
	"$a":  1,
	"B":   2,
	"c_c": 3,
	"d1":  4,
	"VaR_vAr": 10
}

result, _ := engine.Calculate("$a + B + c_c + d1 + 10 + VaR_vAr", vars)
// 30.0
```
- Can contains letters ( a-z | A-Z ), underscore ( _ ), dolar sign ( $ ) or a number ( 0-9 ).
- Cannot start with a number.
- Cannot start with underscore.

### 📍 Built-in Constants

| Constant        |  Description | More Information |
| ------------- | -------|----|
| e |   Euler's number  | https://oeis.org/A001113 |
| pi |   Pi| https://oeis.org/A000796 |

```go
result, _ := engine.Calculate("2*pi", nil)
// 6.283185307179586
```

### 🔧 Comprehensive Function Library

The following mathematical functions are out of the box supported:

| Function | Arguments       | Description         | More Information                                                                               |
| -------- | --------------- | ------------------- | ---------------------------------------------------------------------------------------------- |
| sin      | sin(x)          | Sine                | https://pkg.go.dev/math#Sin                                                                    |
| cos      | cos(x)          | Cosine              | https://pkg.go.dev/math#Cos                                                                    |
| asin     | asin(x)         | Arcsine             | [https://pkg.go.dev/math#Asin](https://pkg.go.dev/math#Asin)                                   |
| acos     | acos(x)         | Arccosine           | https://pkg.go.dev/math#Acos                                                                   |
| tan      | tan(x)          | Tangent             | https://pkg.go.dev/math#Tan                                                                    |
| atan     | atan(x)         | Arctangent          | https://pkg.go.dev/math#Atan                                                                   |
| log      | log(x)          | Logarithm           | https://pkg.go.dev/math#Log                                                                    |
| sqrt     | sqrt(x)         | Square Root         | https://pkg.go.dev/math#Sqrt                                                                   |
| trunc    | trunc(x)        | Truncate            | https://pkg.go.dev/math#Trunc                                                                  |
| floor    | floor(x)        | Floor               | https://pkg.go.dev/math#Floor                                                                  |
| ceil     | ceil(x)         | Ceil                | https://pkg.go.dev/math#Ceil                                                                   |
| round    | round(x \[,y\]) | Round               | Rounds a number to a specified number of digits where 'x' is the number and 'y' is the digits. |
| random   | random(x)       | Random              | Generate a random double value between 0.0 and 1.0 where 'x' is the seed.                      |
| if       | if(a,b,c)       | Excel's IF Function | IF 'a' IS true THEN 'b' ELSE 'c'.                                                              |
| max      | max(x1,…,xn)    | Maximum             | Return the maximum number of a series.                                                         |
| min      | min(x1,…,xn)    | Minimum             | Return the minimum number of a series.                                                         |


```go

// Sin (ordinary function)
vars := map[string]interface{}{
   "a":2,
}
ret, _ := engine.Calculate("sin(100)+a", vars)
// 1.4936343588902412

// Round
retRound, _ := engine.Calculate("round(1.234567,2)", nil)
// 1.23

// If
vars := map[string]interface{}{
   "a":4,
}

ifresult, _ := engine.Calculate("if(2+2==a, 10, 5)", varsIf)
// 10.0

// MAX
max, _ := engine.Calculate("max(5,6,3,-4,5,3,7,8,13,100)", nil)
// 100.0



```

### Custom Functions 

Custom functions allow programmers to add additional functions besides the ones already supported (sin, cos, asin, …). Functions are required to have a unique name. The existing functions cannot be overwritten.

```go
engine.AddFunction("addTwo", func(arguments ...interface{}) float64{
		return arguments[0] + 2
}, true)

result, _ := engine.Calculate("addTwo(2.0)", nil)
// 4.0

```

### Compile Time Constants

Variables as defined in a formula can be replaced by a constant value at compile time. This feature is useful in case that a number of the parameters don't frequently change and that the formula needs to be executed many times. Thusfore it is better because constants could be optimizated on 'Optimization phase'.


```go

consts := map[string]interface{}{
   "a":1,
}
formula := engine.BuildWithConstants("a+b+c", consts)
// It's the same as 'engine.Build("1+b+c")' but without dealing with string replace

vars := map[string]interface{}{
   "b":2,
   "c":5
}

result, := formula(vars)
// 8.0
```

## Benchmark 

https://github.com/mrxrsd/golang-expression-evaluation-comparison

### goJACEgo vs Others

| Test                         |                     | 
|------------------------------|---------------------| 
| Benchmark_bexpr-8            |        2278 ns/op   | 
| Benchmark_celgo-8            |         127.0 ns/op | 
| Benchmark_evalfilter-8       |        1646 ns/op   | 
| Benchmark_expr-8             |         119.1 ns/op | 
| Benchmark_goja-8             |         306.9 ns/op | 
| Benchmark_gojacego-8         |         117.3 ns/op | 
| Benchmark_govaluate-8        |         259.9 ns/op | 
| Benchmark_gval-8             |         295.0 ns/op | 
| Benchmark_otto-8             |         951.2 ns/op | 
| Benchmark_starlark-8         |        5971 ns/op   | 

### goJACEgo vs Govaluate vs Expr vs Gval

| Test                                   | Gojacego    | Govaluate   |  Expr        | Gval         |
| -------------------------------------- | ----------- | ----------- |--------------|--------------|
| BenchmarkEvaluationNumericLiteral      |  5.42 ns/op | 71.73 ns/op |  87.71 ns/op |   1.89 ns/op |
| BenchmarkEvaluationLiteralModifiers    |  5.63 ns/op | 180.8 ns/op |  69.92 ns/op |   1.85 ns/op |
| BenchmarkEvaluationParameter           | 11.25 ns/op | 72.47 ns/op |  69.75 ns/op |  147.1 ns/op |
| BenchmarkEvaluationParameters          | 31.91 ns/op | 122.0 ns/op |  202.2 ns/op |  315.6 ns/op |
| BenchmarkEvaluationParametersModifiers | 56.32 ns/op | 233.3 ns/op |  368.6 ns/op |  378.5 ns/op |
| BenchmarkComplexPrecedenceMath         |  4.73 ns/op | 18.20 ns/op |  67.96 ns/op |   1.93 ns/op |
| BenchmarkMath                          | 39.22 ns/op | 243.7 ns/op |  252.1 ns/op |  395.3 ns/op |


Disclaimer: GoJACEgo has only mathematical and logical operators while others has more features.  

