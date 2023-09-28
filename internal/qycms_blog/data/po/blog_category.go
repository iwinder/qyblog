package po

import (
	"database/sql"
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type CategoryPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string        `json:"name" gorm:"column:name;comment:分组名称""`
	Identifier        string        `json:"identifier" gorm:"unique;column:identifier;comment:分组标识，唯一英文"`
	Description       string        `json:"description" gorm:"column:description;comment:简述"`
	ParentId          sql.NullInt64 `json:"parentId" gorm:"column:parent_id;default: 0;comment:父ID"`
}

type CategoryPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CategoryPO `json:"items"`
}

type CategoryPOListOption struct {
	metaV1.ListOptions `json:"page"`
	CategoryPO         `json:"item"`
}

func (o *CategoryPO) TableName() string {
	return "qy_blog_category"
}

func (o *CategoryPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *CategoryPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "category-")
	if o.Sort <= 0 {
		o.Sort = int(o.ID)
	}
	return tx.Save(o).Error
}

func (o *CategoryPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *CategoryPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
