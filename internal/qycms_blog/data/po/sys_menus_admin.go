package po

import (
	"database/sql"
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type MenusAdminPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string         `json:"name" gorm:"column:name;comment:展示名称"`                         // 展示名称
	BreadcrumbName    string         `json:"breadcrumbName" gorm:"column:breadcrumbName;comment:标签页名称"`    // 标签页名称
	Identifier        string         `json:"identifier" gorm:"unique;column:identifier;comment:路由名称，唯一英文"` // 路由名称
	ParentId          sql.NullInt64  `json:"parentId" gorm:"column:parent_id;default: 0;comment:父菜单ID"`    // 父菜单ID
	Icon              string         `json:"icon" gorm:"column:icon;comment:icon图标"`                       // Icon图标
	MType             int            `json:"mtype" gorm:"column:mtype;default: 1;comment:路由类型:1 菜单,2 按钮"`  // 路由类型
	Path              string         `json:"path" gorm:"column:path;comment:路由Path"`                       // 路由 path
	Redirect          string         `json:"redirect" gorm:"column:redirect;comment:路由重定向地址"`              // 路由重定向地址
	Component         sql.NullString `json:"component" gorm:"column:component;comment:对应前端文件路径"`           // 对应前端文件路径
	Sort              int            `json:"sort" gorm:"column:sort;comment:排序标记"`                         // 排序标记
	Level             int            `json:"-"`
	//Roles             []*RolePO `gorm:"many2many:qy_sys_role_menus_admin;"`
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
	o.InstanceID = idUtil.GetInstanceID(o.ID, "menus-admin-")
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
