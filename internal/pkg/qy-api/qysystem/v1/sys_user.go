package v1

import (
	"encoding/json"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/auth"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/utils/idUtil"
	"gorm.io/gorm"
)

type User struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Username          string `json:"username,omitempty" gorm:"colum:username;type:varchar(255);not null"`
	Nickname          string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	Avatar            string `json:"avatar" gorm:"column:avatar" validate:"omitempty"`
	Password          string `json:"password,omitempty" gorm:"column:password" validate:"required"`
	Salt              string `json:"-" gorm:"-" validate:"omitempty"`
	Email             string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone             string `json:"phone" gorm:"column:phone" validate:"omitempty"`
	AdminFlag         bool   `json:"adminFlag,omitempty" gorm:"column:admin_flag" validate:"omitempty"`
}

type UserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

func (u *User) TableName() string {
	return "qy_sys_user"
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idUtil.GetInstanceID(u.ID, "user-")
	u.Sort = int(u.ID)
	return tx.Save(u).Error
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password, err = auth.Encrypt(u.Password + u.Salt)
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
