package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type FileLibConfigPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	AccessKey         string `json:"accessKey" gorm:"column:access_key;comment:密钥AccessKey"`
	SecretKey         string `json:"secretKey" gorm:"column:secret_key;comment:密钥SecretKey"`
	Bucket            string `json:"bucket" gorm:"column:bucket;comment:存储空间"`
	Prefix            string `json:"prefix" gorm:"column:prefix;comment:前缀"`
	Domain            string `json:"domain" gorm:"column:domain;comment:域名"`
	Endpoint          string `json:"endpoint" gorm:"column:endpoint;comment:绑定域名"`
	TypeId            uint64 `json:"typeId" gorm:"column:type_id;default: 0;comment:媒体库类型ID"`
}
type FileLibConfigPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*FileLibConfigPO `json:"items"`
}

type FileLibConfigPOListOption struct {
	metaV1.ListOptions `json:"page"`
	FileLibConfigPO    `json:"item"`
}

func (o *FileLibConfigPO) TableName() string {
	return "qy_file_lib_config"
}

func (o *FileLibConfigPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *FileLibConfigPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "file-lib-config-")
	return tx.Save(o).Error
}

func (o *FileLibConfigPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *FileLibConfigPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
