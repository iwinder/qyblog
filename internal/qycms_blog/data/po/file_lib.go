package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type FileLibPO struct {
	metaV1.ObjectMeta
	OriginFileName string `json:"accessKey" gorm:"column:origin_fileName;comment:原始名称"`
	Fname          string `json:"fname" gorm:"column:fname;comment:文件名"`
	Fsize          uint64 `json:"fsize" gorm:"column:fsize;comment:文件大小"`
	Extention      string `json:"extention" gorm:"column:extention;comment:文件扩展名"`
	MimeType       string `json:"mimeType" gorm:"column:mime_type;comment:资源的 MIME 类型"`
	Fmd5           string `json:"fmd5" gorm:"column:fmd5;comment: 文件md5值"`
	RelativePath   string `json:"relativePath" gorm:"column:relative_path;comment:通过浏览器访问的相对路径"`
}
type FileLibPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*FileLibPO `json:"items"`
}

type FileLibPOListOption struct {
	metaV1.ListOptions `json:"page"`
	FileLibPO          `json:"item"`
}

func (o *FileLibPO) TableName() string {
	return "qy_file_lib"
}

func (o *FileLibPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *FileLibPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "file-lib-")
	return tx.Save(o).Error
}

func (o *FileLibPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *FileLibPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
