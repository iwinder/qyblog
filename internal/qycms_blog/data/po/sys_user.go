package po

import (
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/bcryptUtil"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type UserPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Username          string    `json:"username,omitempty" gorm:"unique;colum:username;type:varchar(255);not null"`
	Nickname          string    `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	Avatar            string    `json:"avatar" gorm:"column:avatar" validate:"omitempty"`
	Password          string    `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Salt              string    `json:"-" gorm:"-" validate:"omitempty"`
	Email             string    `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone             string    `json:"phone" gorm:"column:phone" validate:"omitempty"`
	AdminFlag         bool      `json:"adminFlag,omitempty" gorm:"column:admin_flag;default: 0;" validate:"omitempty"`
	Roles             []*RolePO `gorm:"-"`
}

type UserPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*UserPO `json:"items"`
}

type UserPOListOption struct {
	metaV1.ListOptions `json:"page"`
	UserPO             `json:"item"`
}

func (u *UserPO) TableName() string {
	return "qy_sys_user"
}

func (u *UserPO) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = bcryptUtil.Encrypt(u.Password + u.Salt)
	return err
}

func (u *UserPO) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idUtil.GetInstanceID(u.ID, "user-")
	u.Sort = int(u.ID)
	return tx.Updates(u).Error
}
