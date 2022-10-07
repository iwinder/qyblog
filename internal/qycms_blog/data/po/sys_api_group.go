package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ApiGroupPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name;comment:分组名称""`
	Identifier        string `json:"identifier" gorm:"column:identifier;comment:分组标识，唯一英文"`
}

type ApiGroupPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ApiGroupPO `json:"items"`
}

type ApiGroupPOListOption struct {
	metaV1.ListOptions `json:"page"`
	ApiGroupPO         `json:"item"`
}

func (o *ApiGroupPO) TableName() string {
	return "qy_sys_api_group"
}

func (o *ApiGroupPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *ApiGroupPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "apis-group-")
	return tx.Save(o).Error
}

func (o *ApiGroupPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *ApiGroupPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
