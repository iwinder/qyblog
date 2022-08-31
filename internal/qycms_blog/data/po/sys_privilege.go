package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type PrivilegePO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Url               string `json:"url" gorm:"url"`
	Icon              string `json:"icon" gorm:"icon"`
	identifier        string `json:"identifier" gorm:"identifier"`
	hided             bool   `json:"hided" gorm:"hided"`
	description       string `json:"description" gorm:"description"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (p *PrivilegePO) TableName() string {
	return "qy_sys_privilege"
}

func (p *PrivilegePO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (p *PrivilegePO) AfterCreate(tx *gorm.DB) (err error) {
	p.InstanceID = idUtil.GetInstanceID(p.ID, "privilege-")
	return tx.Save(p).Error
}

func (p *PrivilegePO) BeforeUpdate(tx *gorm.DB) (err error) {
	p.ExtendShadow = p.Extend.String()
	return err
}

func (p *PrivilegePO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(p.ExtendShadow), &p.Extend); err != nil {
		return err
	}
	return nil
}
