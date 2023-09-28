package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type TagsPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name;comment:分组名称""`
	Identifier        string `json:"identifier" gorm:"unique;column:identifier;comment:分组标识，唯一英文"`
	Description       string `json:"description" gorm:"column:description;comment:简述"`
}

type TagsPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*TagsPO `json:"items"`
}

type TagsPOListOption struct {
	metaV1.ListOptions `json:"page"`
	TagsPO             `json:"item"`
}

func (o *TagsPO) TableName() string {
	return "qy_blog_tags"
}

func (o *TagsPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *TagsPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "tags-")
	if o.Sort <= 0 {
		o.Sort = int(o.ID)
	}
	return tx.Save(o).Error
}

func (o *TagsPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *TagsPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
