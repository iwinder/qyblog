package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ArticleResourcePO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	ArticleID         uint64 ` gorm:"primary_key;column:article_id;comment:文章主键"`
	Name              string `json:"name" gorm:"column:name;comment:名称"`
	Url               string `json:"url" gorm:"column:url;comment:地址"`
	Password          string `json:"type" gorm:"column:password;comment:密码"`
}
type ArticleResourcePOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ArticleResourcePO `json:"items"`
}

type ArticleResourcePOListOption struct {
	metaV1.ListOptions `json:"page"`
	ArticleResourcePO  `json:"item"`
}

func (o *ArticleResourcePO) TableName() string {
	return "qy_blog_article_resource"
}

func (o *ArticleResourcePO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *ArticleResourcePO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "article-resource-")
	o.Sort = int(o.ID)
	return tx.Save(o).Error
}

func (o *ArticleResourcePO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *ArticleResourcePO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
