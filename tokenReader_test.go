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
	reader := NewTokenReader('.', ',')
	_, err := reader.Read("")

	if !errorContains(err, "formula cannot be empty") {
		test.Errorf("unexpected error: %v", err)
	}
}

func TestTokenReader1(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("42+31")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}

func TestTokenReader2(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("(42+31)")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 5 {
		test.Errorf("Count - expected: 5, got: %d", len(ret))
	}
}

func TestTokenReader3(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("(42+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 4 {
		test.Errorf("Count - expected: 4, got: %d", len(ret))
	}
}

func TestTokenReader4(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("(42+ 8) *2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 7 {
		test.Errorf("Count - expected: 7, got: %d", len(ret))
	}
}

func TestTokenReader5(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("(42.87+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 4 {
		test.Errorf("Count - expected: 4, got: %d", len(ret))
	}
}

func TestTokenReader6(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("(var+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 4 {
		test.Errorf("Count - expected: 4, got: %d", len(ret))
	}
}

func TestTokenReader12(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("-2.1")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 1 {
		test.Errorf("Count - expected: 1, got: %d", len(ret))
	}
}

func TestTokenReader32(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("-e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 2 {
		test.Errorf("Count - expected: 2, got: %d", len(ret))
	}
}

func TestTokenReader33(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("1-e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}

func TestTokenReader34(test *testing.T) {
	reader := NewTokenReader('.', ',')
	ret, err := reader.Read("1+e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}
