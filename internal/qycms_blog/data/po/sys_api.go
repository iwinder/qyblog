package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ApiPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	ApiGroup          string    `json:"name" gorm:"column:api_group;comment:API分组""`
	Method            string    `json:"method" gorm:"column:method;comment:请求方法"`
	Path              string    `json:"identifier" gorm:"path;comment:API路径"`
	Description       string    `json:"description" gorm:"description;comment:API简介"`
	Roles             []*RolePO `gorm:"many2many:qy_sys_user_role;"`
}

type ApiPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ApiPO `json:"items"`
}

type ApiPOListOption struct {
	metaV1.ListOptions `json:"page"`
	ApiPO              `json:"item"`
}

func (o *ApiPO) TableName() string {
	return "qy_sys_api"
}

func (o *ApiPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *ApiPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "api-")
	return tx.Save(o).Error
}

func (o *ApiPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *ApiPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
