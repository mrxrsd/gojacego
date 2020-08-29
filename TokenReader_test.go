package gojacego

import (
	"strings"
	"testing"
)

func errorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestTokenReader(test *testing.T) {
	reader := NewTokenReader('.')
	_, err := reader.Read("")

	if !errorContains(err, "formula cannot be empty") {
		test.Errorf("unexpected error: %v", err)
	}
}

func TestTokenReader1(test *testing.T) {
	reader := NewTokenReader('.')
	ret, err := reader.Read("42+31")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}
