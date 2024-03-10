package idUtil

import (
	"github.com/iwinder/qyblog/internal/pkg/utils/iputil"
	"github.com/iwinder/qyblog/internal/pkg/utils/stringUtil"

	"github.com/sony/sonyflake"
	hashids "github.com/speps/go-hashids"
	"math/rand"
)

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		ip := iputil.GetLocalIP()

		return uint16([]byte(ip)[2])<<8 + uint16([]byte(ip)[3]), nil
	}

	sf = sonyflake.NewSonyflake(st)
}

func GetIntID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}

	return id
}

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
func GetUUID36(prefix string) string {
	id := GetIntID()
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36

	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	i, err := h.Encode([]int{int(id)})
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
