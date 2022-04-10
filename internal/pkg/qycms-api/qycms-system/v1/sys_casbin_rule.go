package v1

import (
	"encoding/json"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/idUtil"
	"gorm.io/gorm"
)

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

func (c *CasbinRule) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (c *CasbinRule) AfterCreate(tx *gorm.DB) (err error) {
	c.InstanceID = idUtil.GetInstanceID(c.ID, "casbinru-")
	return tx.Save(c).Error
}

func (c *CasbinRule) BeforeUpdate(tx *gorm.DB) (err error) {
	c.ExtendShadow = c.Extend.String()
	return err
}

func (c *CasbinRule) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(c.ExtendShadow), &c.Extend); err != nil {
		return err
	}
	return nil
}
