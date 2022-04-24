package v1

import (
	"encoding/json"
	metaV1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/idUtil"
	"gorm.io/gorm"
)

type Privilege struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Url               string `json:"url" gorm:"url"`
	Icon              string `json:"icon" gorm:"icon"`
	identifier        string `json:"identifier" gorm:"identifier"`
	hided             bool   `json:"hided" gorm:"hided"`
	description       string `json:"description" gorm:"description"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (p *Privilege) TableName() string {
	return "qy_sys_privilege"
}

func (p *Privilege) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (p *Privilege) AfterCreate(tx *gorm.DB) (err error) {
	p.InstanceID = idUtil.GetInstanceID(p.ID, "privilege-")
	return tx.Save(p).Error
}

func (p *Privilege) BeforeUpdate(tx *gorm.DB) (err error) {
	p.ExtendShadow = p.Extend.String()
	return err
}

func (p *Privilege) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(p.ExtendShadow), &p.Extend); err != nil {
		return err
	}
	return nil
}
