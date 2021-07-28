package gojacego

func ToLowerFast(s string) string {
	b := make([]byte, len(s))
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		b[i] = byte(c)
	}
	return string(b)
}
