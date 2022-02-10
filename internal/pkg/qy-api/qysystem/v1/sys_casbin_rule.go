package v1

import metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"

type CasbinRule struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	PType             string `json:"p_type" gorm:"size:100;"`
	V0                string `json:"v0" gorm:"size:100;"`
	V1                string `json:"v1" gorm:"size:100;"`
	V2                string `json:"v2" gorm:"size:100;"`
	V3                string `json:"v3" gorm:"size:100;"`
	V4                string `json:"v4" gorm:"size:100;"`
	V5                string `json:"v5" gorm:"size:100;"`
	RType             int    `json:"r_type" gorm:"size:1;DEFAULT:0;"`
}

func (c *CasbinRule) TableName() string {
	return "qy_sys_casbin_rule"
}
