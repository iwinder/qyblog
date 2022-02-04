package v1

import (
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
)

type User struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Username          string `json:"username,omitempty" gorm:"colum:username;type:varchar(255);not null"`
	Nickname          string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`
	Password          string `json:"password,omitempty" gorm:"column:passeord" validate:"required"`
	Email             string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`
	Phone             string `json:"phone" gorm:"column:phone" validate:"omitempty"`
	AdminFlag         int    `json:"adminFlag,omitempty" gorm:"column:admin_flag" validate:"omitempty"`
}

type UserList struct {
	metav1.ListMeta `json:",inline"`
	Items           []*User `json:"items"`
}

func (u *User) TableName() string {
	return "qy_sys_user"
}
