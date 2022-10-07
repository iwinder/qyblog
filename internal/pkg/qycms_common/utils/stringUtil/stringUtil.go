package stringUtil

import (
	"strings"
	"unicode/utf8"
)

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

func snakeString(str string) string {

	return SnakeStringWidthByteTag(str, '_')
}

func SnakeStringWidthByteTag(str string, tag byte) string {
	data := make([]byte, 0, len(str)*2)
	j := false
	num := len(str)
	for i := 0; i < num; i++ {
		d := str[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, tag)
		}
		if d != tag {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}
