# Gojacego (W.I.P)
Gojacego is a high performance calculation engine for Go and it is a port of Jace.NET. 

'Jace' stands for "Just Another Calculation Engine".
 
## What does it do?
Gojacego can interprete and execute strings containing mathematical formulas. These formulas can rely on variables. If variables are used, values can be provided for these variables at execution time of the mathematical formula.

## Architecture
Gojacego follows a design similar to most of the modern compilers. Interpretation and execution is done in a number of phases:

### Tokenizing
During the tokenizing phase, the string is converted into the different kind of tokens: variables, operators and constants.
### Abstract Syntax Tree Creation
During the abstract syntax tree creation phase, the tokenized input is converted into a hierarchical tree representing the mathematically formula. This tree unambiguously stores the mathematical calculations that must be executed.
### Optimization
During the optimization phase, the abstract syntax tree is optimized for executing.

## Getting Started 

```go
engine := NewCalculationEngine(&JaceOptions{
			decimalSeparator: '.',
			argumentSeparator: ',',
			caseSensitive:    false,
			optimizeEnabled:  true,
		})

vars := map[string]interface{}{
   "a":2,
   "b":5
}

result, _ := engine.Calculate("a*b", vars)
// 10.0
```

## Features

### Basic Operations 

The following mathematical operations are supported:
* Addition: +
* Subtraction: -
* Multiplication: *
* Division: /
* Modulo: %
* Exponentiation: ^

### Standard Constants

| Constant        |  Description | More Information |
| ------------- | -------|----|
| e |   Euler's number  | https://oeis.org/A001113 |
| pi |   Pi| https://oeis.org/A000796 |

```go
result, _ := engine.Calculate("2*pi", nil)
// 6.283185307179586
```

### Standard Functions (W.I.P)

The following mathematical functions are out of the box supported:

| Function      | Arguments  |  Description | More Information |
| ------------- |-------------| -------|----|
| sin | sin(A1)|  Sine  | https://pkg.go.dev/math#Sin |
| cos | cos(A1)|  Cosine| https://pkg.go.dev/math#Cos |
| max| max(A1, ..., An)|Maximum||


```go
max, _ := engine.Calculate("max(5,6,3,-4,5,3,7,8,13,100)", nil)
// 100.0


vars := map[string]interface{}{
   "a":2,
}
ret, _ := engine.Calculate("sin(100)+a", vars)
// 1.4936343588902412
```

### Custom Functions (W.I.P)

Custom functions allow programmers to add additional functions besides the ones already supported (sin, cos, asin, …). Functions are required to have a unique name (this name is case insensitive). The existing functions cannot be overwritten.

### Compile Time Constants (W.I.P)

Variables as defined in a formula can be replaced by a constant value at compile time. This feature is useful in case that a number of the parameters don't frequently change and that the formula needs to be executed many times.

