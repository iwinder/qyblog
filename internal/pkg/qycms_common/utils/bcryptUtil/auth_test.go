package bcryptUtil

import (
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	str := stringUtil.MD5ByStr("xxxx")
	log.Println("结果1：", str)
	//a := PinyinConvert("【译】Oracle调优技巧22：Hash Outer Join")
	a, err := Encrypt(str + "xxxxx")
	if err != nil {
		log.Fatal("异常原因：", err)
	}
	log.Println("结果：", a)
}
