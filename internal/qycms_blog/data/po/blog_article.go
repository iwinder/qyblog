package po

import (
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
	"time"
)

type ArticlePO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Title             string    `json:"title,omitempty" gorm:"colum:title;type:varchar(255);not null;comment:标题"`
	PermaLink         string    `json:"permaLink" gorm:"column:perma_link;type:varchar(255);comment:链接"  `
	CanonicalLink     string    `json:"canonicalLink" gorm:"column:canonical_link;type:varchar(255);comment:规范链接"  `
	Summary           string    `json:"summary" gorm:"column:summary;type:varchar(255);comment:摘要 "  `
	Thumbnail         string    `json:"thumbnail" gorm:"column:thumbnail;type:varchar(255);comment:缩略图 "  `
	Password          string    `json:"password" gorm:"column:password;type:varchar(255);comment:密码 "  `
	Status            int       `json:"status" gorm:"column:status;default 0;comment:状态,0:草稿,1:公共：2:私密：3:加密 "  `
	Atype             int       `json:"atype" gorm:"column:atype;default 1;comment:类型：1：文章，2：页面"`
	AuthorId          uint64    `json:"authorId" gorm:"column:author_id;comment: 作者id"  `
	CategoryId        uint64    `json:"categoryId" gorm:"column:category_id;comment: 所属分类"  `
	CommentAgentId    uint64    `json:"commentAgentId" gorm:"column:comment_agent_id;comment: 评论组件id"  `
	Published         bool      `json:"published" gorm:"column:published;default 0;comment: 是否发布：0不发布， 1 发布"  `
	ViewCount         int32     `json:"viewCount" gorm:"column:view_count;default 0;comment:浏览人数"  `
	LikeCount         int32     `json:"likeCount" gorm:"column:like_Count;default 0;comment: 点赞数"  `
	HateCount         int32     `json:"hateCount" gorm:"column:hate_Count;default 0;comment: 点踩数"  `
	PublishedAt       time.Time `json:"publishedAt" gorm:"column:published_at;comment:发布日期"  `
}

func (u *ArticlePO) TableName() string {
	return "qy_blog_article"
}

func (u *ArticlePO) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idUtil.GetInstanceID(u.ID, "article-")
	u.Sort = int(u.ID)
	return tx.Save(u).Error
}

type ArticlePOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ArticlePO `json:"items"`
}

type ArticlePOListOption struct {
	metaV1.ListOptions `json:"page"`
	ArticlePO          `json:"item"`
}
