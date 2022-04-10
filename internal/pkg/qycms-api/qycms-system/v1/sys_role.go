package v1

import (
	"encoding/json"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/idUtil"
	"gorm.io/gorm"
)

type Role struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name""`
	Key               string `json:"key" gorm:"key"`
}

func (r *Role) TableName() string {
	return "qy_sys_role"
}

func (r *Role) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (r *Role) AfterCreate(tx *gorm.DB) (err error) {
	r.InstanceID = idUtil.GetInstanceID(r.ID, "role-")
	return tx.Save(r).Error
}

func (r *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	r.ExtendShadow = r.Extend.String()
	return err
}

func (r *Role) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(r.ExtendShadow), &r.Extend); err != nil {
		return err
	}
	return nil
}
