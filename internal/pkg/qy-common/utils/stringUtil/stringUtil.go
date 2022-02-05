package stringUtil

import "unicode/utf8"

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for i := 0; i < size; {
		r, n := utf8.DecodeRuneInString(s[i:])
		i += n
		utf8.EncodeRune(buf[size-i:], r)
	}
	return string(buf)
}
