package v1

import metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"

type Privilege struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Url               string `json:"url" gorm:"url"`
	Icon              string `json:"icon" gorm:"icon"`
	identifier        string `json:"identifier" gorm:"identifier"`
	Sort              int    `json:"sort" gorm:"sort"`
	hided             bool   `json:"hided" gorm:"hided"`
	description       string `json:"description" gorm:"description"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (p *Privilege) TableName() string {
	return "qy_sys_privilege"
}
