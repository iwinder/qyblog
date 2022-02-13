package v1

import (
	"encoding/json"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/utils/idUtil"
	"gorm.io/gorm"
)

type Menu struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"name"`
	Path              string `json:"path" gorm:"path"`
	TargetId          uint64 `json:"targetId" gorm:"target_id"`
	TargetType        string `json:"targetType" gorm:"target_type"`
	blanked           bool   `json:"blanked" gorm:"blanked"`
	parentId          uint64 `json:"parentId" gorm:"column:parent_id"`
}

func (m *Menu) TableName() string {
	return "qy_sys_menu"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (m *Menu) AfterCreate(tx *gorm.DB) (err error) {
	m.InstanceID = idUtil.GetInstanceID(m.ID, "user-")
	return tx.Save(m).Error
}

func (u *Menu) BeforeUpdate(tx *gorm.DB) (err error) {
	//u.Password, err = auth.Encrypt(u.Password + u.Salt)
	u.ExtendShadow = u.Extend.String()
	return err
}

func (m *Menu) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(m.ExtendShadow), &m.Extend); err != nil {
		return err
	}
	return nil
}
