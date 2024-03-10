package stringUtil

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/google/uuid"
	shortuuid "github.com/lithammer/shortuuid/v4"
	pinyin "github.com/mozillazg/go-pinyin"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	SUMMARYDEFNUM = 99
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

func MD5(str []byte) string {
	has := md5.Sum(str)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func MD5ByStr(str string) string {
	data := []byte(str) //切片
	return MD5(data)
}

func GetUUID() string {
	key := uuid.New().String()
	return strings.ReplaceAll(key, "-", "")
}
func GetShortUuid() string {
	u := shortuuid.New()
	return u
}

var pinyinArgs = pinyin.NewArgs()

func PinyinConvert(data string) string {

	pinyinArgs.Separator = ""
	pinyinArgs.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	newStr := strings.ToLower(pinyin.Slug(data, pinyinArgs))
	sampleRegexp := regexp.MustCompile(`\s`)
	result1 := sampleRegexp.ReplaceAllString(newStr, "-")
	sampleRegexp = regexp.MustCompile(`[^\w]`)
	result := sampleRegexp.ReplaceAllString(result1, "-")
	if strings.Index(result, "-") == 0 {
		result = result[1:len(result)]
	}

	return result
}

func RemoveHtml(data string) string {
	sampleRegexp := regexp.MustCompile(".*?<body.*?>(.*?)<\\/body>")
	result1 := sampleRegexp.ReplaceAllString(data, "$1")
	sampleRegexp = regexp.MustCompile("</?[a-zA-Z]+[^><]*>")
	result := sampleRegexp.ReplaceAllString(result1, "")
	content := strings.ReplaceAll(result, "\n", "")
	return content
}

func RemoveHtmlAndSubstring(data string) string {
	str := RemoveHtml(data)
	temp := []rune(str)
	length := SUMMARYDEFNUM
	if len(temp) < SUMMARYDEFNUM {
		length = len(temp)
	}
	return string(temp[:length])
}
