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

func testLen(test *testing.T, tokens []token, expected int) {
	if len(tokens) != expected {
		test.Errorf("Count - expected: %d, got: %d", expected, len(tokens))
	}
}

func testToken(test *testing.T, token token, value string, start int, len int) {

	if tokenValue, ok := token.Value.(string); ok && tokenValue != value {
		test.Errorf("value expected: %s, got: %s", value, tokenValue)
	}

	if tokenValue, ok := token.Value.(rune); ok && string(tokenValue) != value {
		test.Errorf("value expected: %s, got: %s", value, string(tokenValue))
	}

	if token.StartPosition != start {
		test.Errorf("start expected: %d, got: %d", start, token.StartPosition)
	}

	if token.Length != len {
		test.Errorf("Length expected: %d, got: %d", len, token.Length)
	}
}

func TestTokenReader(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("")

	if !errorContains(err, "formula cannot be empty") {
		test.Errorf("unexpected error: %v", err)
	}
}

func TestTokenReader1(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("42+31")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
}

func TestTokenReader2(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("(42+31)")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 5)
}

func TestTokenReader3(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("(42+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 4)
}

func TestTokenReader4(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("(42+ 8) *2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 7)
}

func TestTokenReader5(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("(42.87+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 4)
}

func TestTokenReader6(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("(var+31.0")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 4)
}

func TestTokenReader7(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("varb")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 1)
	testToken(test, ret[0], "varb", 0, 4)
}

func TestTokenReader8(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("varb(")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 2)
	testToken(test, ret[0], "varb", 0, 4)
	testToken(test, ret[1], "(", 4, 1)
}

func TestTokenReader9(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("+varb(")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "+", 0, 1)
	testToken(test, ret[1], "varb", 1, 4)
	testToken(test, ret[2], "(", 5, 1)
}

func TestTokenReader10(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("var1+2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "var1", 0, 4)
	testToken(test, ret[1], "+", 4, 1)
	testToken(test, ret[2], "2", 5, 1)
}

func TestTokenReader11(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("5.1%2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "5.1", 0, 3)
	testToken(test, ret[1], "%", 3, 1)
	testToken(test, ret[2], "2", 4, 1)
}

func TestTokenReader12(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("-2.1")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 1)
	testToken(test, ret[0], "-2.1", 0, 4)
}

func TestTokenReader13(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("5-2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "5", 0, 1)
	testToken(test, ret[1], "-", 1, 1)
	testToken(test, ret[2], "2", 2, 1)
}

func TestTokenReader14(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("5*-2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "5", 0, 1)
	testToken(test, ret[1], "*", 1, 1)
	testToken(test, ret[2], "-2", 2, 2)
}

func TestTokenReader15(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("5*(-2)")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 5)
	testToken(test, ret[0], "5", 0, 1)
	testToken(test, ret[1], "*", 1, 1)
	testToken(test, ret[2], "(", 2, 1)
	testToken(test, ret[3], "-2", 3, 2)
	testToken(test, ret[4], ")", 5, 1)
}

func TestTokenReader16(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("5*-(2+43)")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 8)
	testToken(test, ret[0], "5", 0, 1)
	testToken(test, ret[1], "*", 1, 1)
	testToken(test, ret[2], "_", 2, 1)
	testToken(test, ret[3], "(", 3, 1)
	testToken(test, ret[4], "2", 4, 1)
	testToken(test, ret[5], "+", 5, 1)
	testToken(test, ret[6], "43", 6, 2)
	testToken(test, ret[7], ")", 8, 1)
}

func TestTokenReader17(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("logn(2,5)")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 6)
	testToken(test, ret[0], "logn", 0, 4)
	testToken(test, ret[1], "(", 4, 1)
	testToken(test, ret[2], "2", 5, 1)
	testToken(test, ret[3], ",", 6, 1)
	testToken(test, ret[4], "5", 7, 1)
	testToken(test, ret[5], ")", 8, 1)
}

func TestTokenReader18(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("var_1+2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "var_1", 0, 5)
	testToken(test, ret[1], "+", 5, 1)
	testToken(test, ret[2], "2", 6, 1)
}

func TestTokenReader20(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("2.11E-3")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 1)
	testToken(test, ret[0], "2.11E-3", 0, 7)
}

func TestTokenReader21(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("var_1+2.11E-3")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "var_1", 0, 5)
	testToken(test, ret[1], "+", 5, 1)
	testToken(test, ret[2], "2.11E-3", 6, 7)
}

func TestTokenReader23(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("2.11e3")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 1)
	testToken(test, ret[0], "2.11e3", 0, 6)
}

func TestTokenReader24(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("1 * e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 3)
	testToken(test, ret[0], "1", 0, 1)
	testToken(test, ret[1], "*", 2, 1)
	testToken(test, ret[2], "e", 4, 1)
}

func TestTokenReader25(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 1)
	testToken(test, ret[0], "e", 0, 1)
}

func TestTokenReader26(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("2.11e3+1.23E4")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}

func TestTokenReader27(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("-(1)^2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	testLen(test, ret, 6)
	testToken(test, ret[0], "_", 0, 1)
	testToken(test, ret[1], "(", 1, 1)
	testToken(test, ret[2], "1", 2, 1)
	testToken(test, ret[3], ")", 3, 1)
	testToken(test, ret[4], "^", 4, 1)
	testToken(test, ret[5], "2", 5, 1)
}

func TestTokenReader28(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read(".")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader29(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("..")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader30(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("..1")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader31(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("0..1")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader32(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("-e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 2 {
		test.Errorf("Count - expected: 2, got: %d", len(ret))
	}
}

func TestTokenReader33(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("1-e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}

func TestTokenReader34(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("1+e")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 3 {
		test.Errorf("Count - expected: 3, got: %d", len(ret))
	}
}

func TestTokenReader35(test *testing.T) {
	reader := newTokenReader('.', ',')
	ret, err := reader.read("1e-3*5+2")

	if err != nil {
		test.Log(err)
		test.Fail()
	}

	if len(ret) != 5 {
		test.Errorf("Count - expected: 5, got: %d", len(ret))
	}
}

func TestTokenReader36(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("2.11E-e3")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader37(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("2.11E-e")

	if err == nil {
		test.Errorf("error should not be null")
	}
}

func TestTokenReader38(test *testing.T) {
	reader := newTokenReader('.', ',')
	_, err := reader.read("3e")

	if err == nil {
		test.Errorf("error should not be null")
	}
}
