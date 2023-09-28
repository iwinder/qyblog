package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type FileLibTypePO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name;comment:媒体库类型名称"`
	Identifier        int    `json:"identifier" gorm:"column:identifier;default: 0;comment:类型标识：1本地， 2七牛，3阿里OSS"`
	Ftype             string `json:"type" gorm:"column:type;default: USER;comment:类型"`
}
type FileLibTypePOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*FileLibTypePO `json:"items"`
}

type FileLibTypePOListOption struct {
	metaV1.ListOptions `json:"page"`
	FileLibTypePO      `json:"item"`
}

func (o *FileLibTypePO) TableName() string {
	return "qy_file_lib_type"
}

func (o *FileLibTypePO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *FileLibTypePO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "file-lib-type-")
	o.Sort = int(o.ID)
	return tx.Save(o).Error
}

func (o *FileLibTypePO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *FileLibTypePO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
