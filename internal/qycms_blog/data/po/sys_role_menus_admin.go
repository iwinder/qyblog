package po

import "time"

type RoleMenusPO struct {
	RoleID    uint64     `gorm:"primary_key;column:role_id;comment:角色主键"`
	MenusID   uint64     ` gorm:"primary_key;column:menus_id;comment:菜单主键"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (r *RoleMenusPO) TableName() string {
	return "qy_sys_role_menus"
}
