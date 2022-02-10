package v1

import metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"

type Menu struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Path              string `json:"path" gorm:"path"`
	TargetId          uint64 `json:"targetId" gorm:"target_id"`
	TargetType        string `json:"targetType" gorm:"target_type"`
	Sort              int    `json:"sort" gorm:"sort"`
	blanked           bool   `json:"blanked" gorm:"blanked"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (m *Menu) TableName() string {
	return "qy_sys_menu"
}
