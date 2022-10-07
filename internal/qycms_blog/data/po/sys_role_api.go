package po

import "time"

type RoleApiPO struct {
	RoleID    uint64     `gorm:"primary_key;column:role_id;comment:角色主键"`
	ApiID     uint64     ` gorm:"primary_key;column:api_id;comment:api主键"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (r *RoleApiPO) TableName() string {
	return "qy_sys_role_api"
}
