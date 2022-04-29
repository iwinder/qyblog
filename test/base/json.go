package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Company struct {
	Id          string `json:"id"`
	CmpPoolName string `json:"cmpPoolName"`
	Text        string `json:"text"`
	Bm          string `json:"bm"`
}

func main() {
	getList()
}

func getList() {
	var list []*Company
	t1 := New("dazgs", "dazgs", "总公司", "00000000")
	list = append(list, t1)
	list = append(list, New("dagd", "dagd", "广东分公司", "44000000"))

	rs, err := json.Marshal(list)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(rs)
	fmt.Println(string(rs))
}

func New(id, name, text, bm string) *Company {
	return &Company{
		Id:          id,
		CmpPoolName: name,
		Text:        text,
		Bm:          bm,
	}
}
