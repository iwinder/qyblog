package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ShortLinkPO struct {
	metaV1.ObjectMeta
	Url         string `json:"url" gorm:"column:url;comment:原地址"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
	Identifier  string `json:"identifier" gorm:"unique;column:identifier;comment:短地址"`
}

type ShortLinkPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ShortLinkPO `json:"items"`
}

type ShortLinkPOListOption struct {
	metaV1.ListOptions `json:"page"`
	ShortLinkPO        `json:"item"`
}

func (o *ShortLinkPO) TableName() string {
	return "qy_sys_short_link"
}

func (o *ShortLinkPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *ShortLinkPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "short-link-")
	return tx.Save(o).Error
}

func (o *ShortLinkPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *ShortLinkPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
