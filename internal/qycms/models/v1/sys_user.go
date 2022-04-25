package v1

import (
	"encoding/json"
	metaV1 "gitee.com/windcoder/qingyucms/internal/pkg/qycms-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/auth"
	"gitee.com/windcoder/qingyucms/internal/pkg/qygo-common/utils/idUtil"
	"gorm.io/gorm"
)

type User struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Username          string `json:"username,omitempty" gorm:"unique;colum:username;type:varchar(255);not null"`
	Nickname          string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	Avatar            string `json:"avatar" gorm:"column:avatar" validate:"omitempty"`
	Password          string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Salt              string `json:"-" gorm:"-" validate:"omitempty"`
	Email             string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone             string `json:"phone" gorm:"column:phone" validate:"omitempty"`
	AdminFlag         bool   `json:"adminFlag,omitempty" gorm:"column:admin_flag" validate:"omitempty"`
}

type UserList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

type UserListOption struct {
	metaV1.ListOptions `json:"page"`
	User               `json:"user"`
}

func (u *User) TableName() string {
	return "qy_sys_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = auth.Encrypt(u.Password + u.Salt)
	return err
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idUtil.GetInstanceID(u.ID, "user-")
	u.Sort = int(u.ID)
	return tx.Save(u).Error
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.ExtendShadow = u.Extend.String()
	return err
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.DeletedAt.Valid {
		// 不为空,则已删除
		u.StatusFlag = 2
	}
	if err := json.Unmarshal([]byte(u.ExtendShadow), &u.Extend); err != nil {
		return err
	}
	return nil
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}
