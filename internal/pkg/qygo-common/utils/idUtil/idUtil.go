package idUtil

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/stringUtil"
	//"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/stringUtil"
	hashids "github.com/speps/go-hashids"
	"math/rand"
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

func NewSecretID() string {
	return randString(Alphabet62, 36)
}

func NewSecretKey() string {
	return randString(Alphabet62, 32)
}

func randString(letters string, n int) string {
	output := make([]byte, n)

	randomness := make([]byte, n)

	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(letters)

	for pos := range output {
		random := randomness[pos]
		randomPos := random % uint8(l)
		output[pos] = letters[randomPos]
	}
	return string(output)

}
