package gojacego

import (
	"testing"
)

func TestBasicInterpreterSubstraction(test *testing.T) {
	ret, _ := Execute(NewSubtractionOperation(Integer,
		NewConstantOperation(Integer, 6),
		NewConstantOperation(Integer, 9)))

	if ret != -3.0 {
		test.Errorf("Expected: -3.0, got: %f", ret)
	}
}
