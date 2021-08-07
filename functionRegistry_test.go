package gojacego

import (
	"testing"
)

func TestAddAndGetFunction(test *testing.T) {
	registryCaseInsensitive := newFunctionRegistry(false)
	registryCaseSensitive := newFunctionRegistry(true)

	fn := func(args ...interface{}) float64 {
		return args[0].(float64) + args[1].(float64)
	}

	registryCaseInsensitive.registerFunction("test", fn, true, false)
	registryCaseSensitive.registerFunction("test", fn, true, false)

	_, found := registryCaseInsensitive.get("test")
	if found != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found1 := registryCaseInsensitive.get("TesT")
	if found1 != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found2 := registryCaseSensitive.get("test")
	if found2 != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found3 := registryCaseSensitive.get("TesT")
	if found3 != false {
		test.Errorf("exptected: false, got: true")
	}
}

func TestFunctionOverwritable(test *testing.T) {
	registry := newFunctionRegistry(false)

	fnAddTwo := func(args ...interface{}) float64 {
		return args[0].(float64) + 2
	}

	fnAddFour := func(args ...interface{}) float64 {
		return args[0].(float64) + 4
	}

	registry.registerFunction("test", fnAddTwo, true, true)
	registry.registerFunction("test", fnAddFour, true, true)

	fn, _ := registry.get("test")
	if item := fn.function(0.0); item != 4 {
		test.Errorf("exptected: 4, got: %f", item)
	}
}

func TestFunctionNotOverwritable(test *testing.T) {
	registry := newFunctionRegistry(false)

	fn := func(args ...interface{}) float64 {
		return args[0].(float64) + 2
	}

	registry.registerFunction("test", fn, false, true)

	shouldPanic(test, func() {
		registry.registerFunction("test", fn, false, true)
	}, "TestNotOverwritable - Panic expected")
}
