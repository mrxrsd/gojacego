package gojacego

import (
	"testing"
)

func TestAddAndGetFunction(test *testing.T) {
	registryCaseInsensitive := NewFunctionRegistry(false)
	registryCaseSensitive := NewFunctionRegistry(true)

	fn := func(args ...float64) (float64, error) {
		return args[0] + args[1], nil
	}

	registryCaseInsensitive.RegisterFunction("test", fn, true, false)
	registryCaseSensitive.RegisterFunction("test", fn, true, false)

	_, found := registryCaseInsensitive.Get("test")
	if found != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found1 := registryCaseInsensitive.Get("TesT")
	if found1 != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found2 := registryCaseSensitive.Get("test")
	if found2 != true {
		test.Errorf("exptected: true, got: false")
	}

	_, found3 := registryCaseSensitive.Get("TesT")
	if found3 != false {
		test.Errorf("exptected: false, got: true")
	}
}

func TestFunctionOverwritable(test *testing.T) {
	registry := NewFunctionRegistry(false)

	fnAddTwo := func(args ...float64) (float64, error) {
		return args[0] + 2, nil
	}

	fnAddFour := func(args ...float64) (float64, error) {
		return args[0] + 4, nil
	}

	registry.RegisterFunction("test", fnAddTwo, true, true)
	registry.RegisterFunction("test", fnAddFour, true, true)

	fn, _ := registry.Get("test")
	if item, _ := fn.function(0); item != 4 {
		test.Errorf("exptected: 4, got: %f", item)
	}
}

func TestFunctionNotOverwritable(test *testing.T) {
	registry := NewFunctionRegistry(false)

	fn := func(args ...float64) (float64, error) {
		return args[0] + 2, nil
	}

	registry.RegisterFunction("test", fn, false, true)

	shouldPanic(test, func() {
		registry.RegisterFunction("test", fn, false, true)
	}, "TestNotOverwritable - Panic expected")
}
