package po

import "time"

type ArticleVisitorPO struct {
	ID        uint64     `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENt;column:id;comment:主键"`
	ArticleId uint64     `gorm:"column:article_id;comment: 文章ID" `
	Ip        string     `json:"ip" gorm:"column:ip;default: 0;comment:浏览者ip"  `
	Agent     string     `json:"agent" gorm:"column:agent;type:text;comment: 浏览者客户端"  `
	Atype     int        `json:"atype" gorm:"column:atype;default: 1;comment:类型：1：web，2：微信"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (r *ArticleVisitorPO) TableName() string {
	return "qy_blog_article_visitor"
}
