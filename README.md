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

## Examples 

```go
engine := NewCalculationEngine(&JaceOptions{
			decimalSeparator: '.',
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
