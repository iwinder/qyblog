package fileUtil

import (
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/dateUtil"
	"os"
	"path/filepath"
	"strings"
)

func GetExt(pathStr string) string {
	suffix := filepath.Ext(pathStr)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return ""
}
func GetLocalPrefixPath(prefix, filename string) (string, string) {
	var path strings.Builder
	path.WriteString(prefix)
	dateStr := dateUtil.GetTimestampOfNowByType(dateUtil.IYYYYiMMi)
	separator := string(filepath.Separator)
	if !strings.HasSuffix(prefix, separator) {
		path.WriteString(separator)
	}
	path.WriteString(ContentPath)
	path.WriteString(dateStr)

	var relativePath strings.Builder
	relativePath.WriteString(VirtualPath)
	relativePath.WriteString(dateStr)
	relativePath.WriteString(filename)
	return path.String(), relativePath.String()
}
func GetOssKey(prefix, filename string) string {
	var path strings.Builder
	if len(prefix) > 0 {
		path.WriteString(prefix)
		separator := string(filepath.Separator)
		if !strings.HasSuffix(prefix, separator) {
			path.WriteString(separator)
		}
	}
	path.WriteString(filename)
	return path.String()
}
func CheckAndMkdirAll(parentPath string) error {
	if _, perr := os.Stat(parentPath); os.IsNotExist(perr) {
		// 创建文件夹，注意这里给的权限时777，可以将这个参数提取出来作为参数传入。
		if perr = os.MkdirAll(parentPath, os.ModePerm); perr != nil {
			return perr
		}
	}
	return nil
}
