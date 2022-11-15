package stringUtil

import (
	"log"
	"testing"
)

func TestRunServer(t *testing.T) {
	//a := PinyinConvert("【译】Oracle调优技巧22：Hash Outer Join")
	a := PinyinConvert("Kava")
	log.Println("结果：", a)
}
