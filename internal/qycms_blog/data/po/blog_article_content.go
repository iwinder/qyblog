package po

import (
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

type ArticleContentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Status            int    `json:"status" gorm:"column:status;default:0;comment:状态,0:草稿,1:公共：2:私密：3:加密 "  `
	Atype             int    `json:"atype" gorm:"column:atype;default:1;comment:类型：1：文章，2：页面"`
	Content           string `json:"content" gorm:"column:content;type:longtext;comment:内容-mkdown"`
	ContentHtml       string `json:"contentHtml" gorm:"column:content_html;type:longtext;comment:内容-html"`
}

func (u *ArticleContentPO) TableName() string {
	return "qy_blog_article_content"
}
func (u *ArticleContentPO) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idUtil.GetInstanceID(u.ID, "article-content-")
	u.Sort = int(u.ID)
	return tx.Save(u).Error
}
