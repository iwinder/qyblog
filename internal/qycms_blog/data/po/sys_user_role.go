package po

import "time"

type UserRolePO struct {
	UserID    uint64     ` gorm:"primary_key;column:user_id;comment:用户主键"`
	RoleID    uint64     `gorm:"primary_key;column:role_id;comment:角色主键"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (r *UserRolePO) TableName() string {
	return "qy_sys_user_role"
}
