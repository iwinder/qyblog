package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type MenusAdminPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string    `json:"name" gorm:"column:name;comment:展示名称"` // 路由name
	Identifier        string    `json:"identifier" gorm:"unique;column:identifier;comment:路由名称，唯一英文"`
	ParentId          uint64    `json:"parentId" gorm:"column:name;comment:父菜单ID"`     // 父菜单ID
	Path              string    `json:"path" gorm:"column:name;comment:路由path"`        // 路由path
	Hidden            bool      `json:"hidden" gorm:"column:name;comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component         string    `json:"component" gorm:"column:name;comment:对应前端文件路径"` // 对应前端文件路径
	Sort              int       `json:"sort" gorm:"column:name;comment:排序标记"`          // 排序标记
	Level             uint      `json:"-"`
	Roles             []*RolePO `gorm:"many2many:qy_sys_user_role;"`
	//children          []MenusAdminPO `json:"children" gorm:"-"`
	//Parameters        []SysBaseMenuParameter                     `json:"parameters"`
	//MenuBtn           []SysBaseMenuBtn                           `json:"menuBtn"`
}

type MenusAdminPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusAdminPO `json:"items"`
}

type MenusAdminPOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusAdminPO       `json:"item"`
}

func (o *MenusAdminPO) TableName() string {
	return "qy_sys_menus_admin"
}

func (o *MenusAdminPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *MenusAdminPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "menus_admin-")
	return tx.Save(o).Error
}

func (o *MenusAdminPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *MenusAdminPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
