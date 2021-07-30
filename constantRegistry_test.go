package gojacego

import (
	"testing"
)

func TestAddAndGetConstant(test *testing.T) {
	registryCaseInsensitive := newConstantRegistry(false)
	registryCaseSensitive := newConstantRegistry(true)

	registryCaseInsensitive.registerConstant("test", 42.0, true)
	registryCaseSensitive.registerConstant("test", 42.0, true)

	val, _ := registryCaseInsensitive.get("test")
	if val != 42.0 {
		test.Errorf("exptected: 42.0, got: %f", val)
	}

	val1, _ := registryCaseInsensitive.get("TesT")
	if val1 != 42.0 {
		test.Errorf("exptected: 42.0, got: %f", val1)
	}

	val2, _ := registryCaseSensitive.get("test")
	if val != 42.0 {
		test.Errorf("exptected: 42.0, got: %f", val2)
	}

	_, found := registryCaseSensitive.get("TesT")
	if found != false {
		test.Errorf("exptected: false, got: true")
	}
}

func TestOverwritable(test *testing.T) {
	registry := newConstantRegistry(false)

	registry.registerConstant("test", 42, true)
	registry.registerConstant("test", 26.3, true)

	val, _ := registry.get("test")
	if val != 26.3 {
		test.Errorf("exptected: 26.3, got: %f", val)
	}
}

// https://gist.github.com/wrunk/4afea3d85cc9feb7fd8fcef5a8a98b5e
func shouldPanic(t *testing.T, f func(), message string) {
	defer func() { recover() }()
	f()
	t.Errorf(message)
}

func TestNotOverwritable(test *testing.T) {
	registry := newConstantRegistry(false)

	registry.registerConstant("test", 42, false)

	shouldPanic(test, func() {
		registry.registerConstant("test", 26.3, false)
	}, "TestNotOverwritable - Panic expected")
}
