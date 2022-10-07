package fileUtil

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestPanic(t *testing.T) {
	str, err := GetAllFile("/home/wind/Work/program/web/gitee.com/windcoder/qyblog-web/qy-console/node_modules/@ant-design/icons-vue/lib/icons")
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	// 通过 JSON 序列化字典数据
	data, _ := json.MarshalIndent(str, "", "    ")

	// 将 JSON 格式数据写入当前目录下的d books 文件（文件不存在会自动创建）
	err = ioutil.WriteFile("icons.json", data, 0644)
	if err != nil {
		log.Fatalf("json connect error: %v", err)
	}

	// 从文件 books 中读取数据
	//dataEncoded, _ = ioutil.ReadFile("books")
	//// 将读取到的数据通过 JSON 解码反序列化为原来的数据类型
	//var booksDecoded map[int]*Book
	//json.Unmarshal(dataEncoded, &booksDecoded)
	//fmt.Printf("%#v", booksDecoded[book1.Id])
}
