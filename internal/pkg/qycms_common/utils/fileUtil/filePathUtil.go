package fileUtil

import (
	"fmt"
	stringUtil "github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/stringUtil"
	"io/ioutil"
	"strings"
)

type IconConfig struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func GetAllFile(pathname string) ([]*IconConfig, error) {
	rd, err := ioutil.ReadDir(pathname)
	objs := make([]*IconConfig, 0, 0)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return objs, err
	}
	var tmpName string
	var tmpLable string

	for _, fi := range rd {
		if !fi.IsDir() {
			tmpName = fi.Name()
			if strings.Contains(tmpName, "Outlined.js") {
				len := strings.Index(tmpName, ".js")
				newTmpName := tmpName[0:len]
				fmt.Println("read dir tmpName:", tmpName, newTmpName)
				tmpLable = stringUtil.SnakeStringWidthByteTag(newTmpName, '-')
				objs = append(objs, &IconConfig{
					Value: newTmpName,
					Name:  tmpLable,
				})

			}

		}
	}
	return objs, nil
}
