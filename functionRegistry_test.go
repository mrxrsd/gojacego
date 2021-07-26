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
	registry := NewConstantRegistry(false)

	registry.RegisterConstant("test", 42, true)
	registry.RegisterConstant("test", 26.3, true)

	val, _ := registry.Get("test")
	if val != 26.3 {
		test.Errorf("exptected: 26.3, got: %f", val)
	}
}

func TestFunctionNotOverwritable(test *testing.T) {
	registry := NewConstantRegistry(false)

	registry.RegisterConstant("test", 42, false)

	shouldPanic(test, func() {
		registry.RegisterConstant("test", 26.3, false)
	}, "TestNotOverwritable - Panic expected")
}
