package po

import (
	"database/sql"
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ApiPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Path              string        `json:"path" gorm:"column:path;comment:API路径"`
	Method            string        `json:"method" gorm:"column:method;comment:请求方法"`
	Description       string        `json:"description" gorm:"column:description;comment:API简介"`
	GroupId           sql.NullInt64 `json:"parentId" gorm:"column:group_id;default: 0;comment:分组ID"` // 分组ID
	ApiGroup          string        `json:"name" gorm:"column:api_group;comment:API分组名称""`
	Identifier        string        `json:"identifier" gorm:"column:identifier;comment:分组标识，唯一英文"`
	Roles             []*RolePO     `gorm:"-"`
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
	o.InstanceID = idUtil.GetInstanceID(o.ID, "apis-")
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
