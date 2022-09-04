package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type CasbinRulePO struct {
	metaV1.ObjectNoInstMeta `json:"metadata,omitempty"`
	PType                   string `json:"ptype" gorm:"column:ptype;size:100;"` // 类型
	V0                      string `json:"v0" gorm:"column:v0;size:100;"`       // sub
	V1                      string `json:"v1" gorm:"column:v1;size:100;"`       // dom
	V2                      string `json:"v2" gorm:"column:v2;size:100;"`       // obj
	V3                      string `json:"v3" gorm:"column:v3;size:100;"`       // act
	V4                      string `json:"v4" gorm:"column:v4;size:100;"`
	V5                      string `json:"v5" gorm:"column:v5;size:100;"`
}

func (c *CasbinRulePO) TableName() string {
	return "qy_sys_casbin_rule"
}

func (c *CasbinRulePO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (c *CasbinRulePO) AfterCreate(tx *gorm.DB) (err error) {
	c.InstanceID = idUtil.GetInstanceID(c.ID, "casbinru-")
	return tx.Save(c).Error
}

func (c *CasbinRulePO) BeforeUpdate(tx *gorm.DB) (err error) {
	c.ExtendShadow = c.Extend.String()
	return err
}

func (c *CasbinRulePO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(c.ExtendShadow), &c.Extend); err != nil {
		return err
	}
	return nil
}
