package v1

import metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"

type Role struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name""`
	Key               string `json:"key" gorm:"key"`
	Sort              int    `json:"Sort" gorm:"sort"`
}

func (r *Role) TableName() string {
	return "qy_sys_role"
}
