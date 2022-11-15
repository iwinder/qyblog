package po

import "time"

type ArticleTagsPO struct {
	ArticleID uint64     ` gorm:"primary_key;column:article_id;comment:文章主键"`
	TagID     uint64     `gorm:"primary_key;column:tag_id;comment:Tag标签主键"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (r *ArticleTagsPO) TableName() string {
	return "qy_blog_article_tags"
}
