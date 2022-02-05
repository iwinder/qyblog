package idUtil

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/utils/stringUtil"
	hashids "github.com/speps/go-hashids"
)

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func GetInstanceID(uid uint64, prefix string) string {
	hd := hashids.NewData()
	hd.Alphabet = Alphabet62
	hd.MinLength = 6
	hd.Salt = "jJK7Hj" + prefix
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	i, err := h.Encode([]int{int(uid)})
	if err != nil {
		panic(err)
	}
	return prefix + stringUtil.Reverse(i)
}
