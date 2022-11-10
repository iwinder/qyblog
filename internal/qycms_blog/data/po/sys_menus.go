package po

import (
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
)

type MenusPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string         `json:"name" gorm:"column:name;comment:展示名称"`                         // 展示名称
	Url    string         `json:"url" gorm:"column:url;comment:标签页名称"`    // 标签页名称
	Blanked        string         `json:"blanked" gorm:"unique;column:blanked;comment:路由名称，唯一英文"` // 路由名称
	ParentId          sql.NullInt64  `json:"parentId" gorm:"column:parent_id;default: 0;comment:父菜单ID"`    // 父菜单ID
	TargetId             sql.NullInt64            `json:"targetId" gorm:"column:target_id;default: 0;comment:路由类型:1 菜单,2 按钮"`  // 路由类型

}

type MenusPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusPO `json:"items"`
}

type MenusPOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusPO       `json:"item"`
}

func (o *MenusPO) TableName() string {
	return "qy_sys_menus"
}

func (o *MenusPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *MenusPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "menus-")
	return tx.Save(o).Error
}

func (o *MenusPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *MenusPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
