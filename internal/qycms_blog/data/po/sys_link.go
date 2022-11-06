package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type LinkPO struct {
	metaV1.ObjectMeta
	Name        string `json:"name" gorm:"column:name;comment:名称"`
	Url         string `json:"url" gorm:"column:url;comment:网址"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
	Ftype       int32  `json:"ftype" gorm:"column:ftype;comment:类型：1首页，2内页"`
}

type LinkPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*LinkPO `json:"items"`
}

type LinkPOListOption struct {
	metaV1.ListOptions `json:"page"`
	LinkPO             `json:"item"`
}

func (o *LinkPO) TableName() string {
	return "qy_sys_link"
}

func (o *LinkPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *LinkPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "link-")
	return tx.Save(o).Error
}

func (o *LinkPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *LinkPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
