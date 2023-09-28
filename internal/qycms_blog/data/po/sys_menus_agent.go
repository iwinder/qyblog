package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type MenusAgentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name" gorm:"column:name;comment:展示名称"`               // 展示名称
	Ftype             string `json:"ftype" gorm:"column:ftype;default: USER;comment:类型"` // 父菜单ID
}

type MenusAgentPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*MenusAgentPO `json:"items"`
}

type MenusAgentPOListOption struct {
	metaV1.ListOptions `json:"page"`
	MenusAgentPO       `json:"item"`
}

func (o *MenusAgentPO) TableName() string {
	return "qy_sys_menus_agent"
}

func (o *MenusAgentPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *MenusAgentPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "menus-agent-")
	return tx.Save(o).Error
}

func (o *MenusAgentPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *MenusAgentPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
