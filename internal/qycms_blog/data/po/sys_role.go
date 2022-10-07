package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type RolePO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string          `json:"name" gorm:"column:name;comment:角色名称""`
	Identifier        string          `json:"identifier" gorm:"column:identifier;comment:唯一英文标识"`
	Users             []*UserPO       `gorm:"-"`
	MenusAdmins       []*MenusAdminPO `gorm:"-"`
	Apis              []*ApiPO        `gorm:"-"`
}

type RolePOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*RolePO `json:"items"`
}

type RolePOListOption struct {
	metaV1.ListOptions `json:"page"`
	RolePO             `json:"item"`
}

func (r *RolePO) TableName() string {
	return "qy_sys_role"
}

func (r *RolePO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (r *RolePO) AfterCreate(tx *gorm.DB) (err error) {
	r.InstanceID = idUtil.GetInstanceID(r.ID, "role-")
	return tx.Save(r).Error
}

func (r *RolePO) BeforeUpdate(tx *gorm.DB) (err error) {
	r.ExtendShadow = r.Extend.String()
	return err
}

func (r *RolePO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(r.ExtendShadow), &r.Extend); err != nil {
		return err
	}
	return nil
}
