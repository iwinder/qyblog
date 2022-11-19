package po

type ArticleContentPO struct {
	ID          uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENt;column:id;comment:主键"`
	StatusFlag  int    `json:"status" gorm:"column:status;default:1;comment:状态,1:草稿,2:公共：3:私密：4:加密 "  `
	Atype       int    `json:"atype" gorm:"column:atype;default:1;comment:类型：1：文章，2：页面"`
	Content     string `json:"content" gorm:"column:content;type:longtext;comment:内容-mkdown"`
	ContentHtml string `json:"contentHtml" gorm:"column:content_html;type:longtext;comment:内容-html"`
}

func (u *ArticleContentPO) TableName() string {
	return "qy_blog_article_content"
}
