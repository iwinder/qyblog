package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type SiteConfigPO struct {
	metaV1.ObjectMeta
	ConfigKey   string `json:"configKey" gorm:"unique;column:config_key;comment:配置key"`
	ConfigValue string `json:"configValue" gorm:"column:config_value;comment:配置value"`
	ConfigName  string `json:"configName" gorm:"column:config_name;comment:配置名称"`
	ConfigTip   string `json:"configTip" gorm:"column:config_tip;comment:配置说明"`
	Ftype       int    `json:"ftype" gorm:"column:ftype;comment:分类"`
}
type SiteConfigPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*SiteConfigPO `json:"items"`
}

type SiteConfigPOListOption struct {
	metaV1.ListOptions `json:"page"`
	SiteConfigPO       `json:"item"`
}

func (o *SiteConfigPO) TableName() string {
	return "qy_sys_site_config"
}

func (o *SiteConfigPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *SiteConfigPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "site-config-")
	return tx.Save(o).Error
}

func (o *SiteConfigPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *SiteConfigPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
