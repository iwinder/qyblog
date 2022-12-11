package po

import (
	"database/sql"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
	"time"
)

type ArticlePO struct { //Status    `json:"status" gorm:"column:status;default: 1;comment:状态,1:草稿,2:公共：3:私密：4:加密 "  `
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	Title             string       `json:"title,omitempty" gorm:"colum:title;type:varchar(255);not null;comment:标题"`
	PermaLink         string       `json:"permaLink" gorm:"unique;column:perma_link;type:varchar(255);comment:链接"  `
	CanonicalLink     string       `json:"canonicalLink" gorm:"column:canonical_link;type:varchar(255);comment:规范链接"  `
	Summary           string       `json:"summary" gorm:"column:summary;type:varchar(255);comment:摘要 "  `
	Thumbnail         string       `json:"thumbnail" gorm:"column:thumbnail;type:varchar(255);comment:缩略图 "  `
	Password          string       `json:"password" gorm:"column:password;type:varchar(255);comment:密码 "  `
	Atype             int          `json:"atype" gorm:"column:atype;default: 1;comment:类型：1：文章，2：页面"`
	CategoryId        uint64       `json:"categoryId" gorm:"column:category_id;comment: 所属分类"  `
	CategoryName      string       `json:"categoryName" gorm:"column:category_name;comment: 所属分类"  `
	CommentAgentId    uint64       `json:"commentAgentId" gorm:"column:comment_agent_id;comment: 评论组件id"  `
	CommentFlag       sql.NullBool `json:"commentFlag" gorm:"column:comment_flag;default: 1;comment: 评论开启状态 true 开启，false 不开启"  `
	Nickname          string       `json:"nickname" gorm:"column:nickname;comment: 作者昵称" validate:"required,min=1,max=30"`
	Published         sql.NullBool `json:"published" gorm:"column:published;default: 0;comment: 是否发布：0不发布， 1 发布"  `
	ViewCount         int32        `json:"viewCount" gorm:"column:view_count;default: 0;comment:浏览人数"  `
	CommentCount      int32        `json:"commentCount" gorm:"column:comment_count;default: 0;comment:评论人数"  `
	LikeCount         int32        `json:"likeCount" gorm:"column:like_Count;default: 0;comment: 点赞数"  `
	HateCount         int32        `json:"hateCount" gorm:"column:hate_Count;default: 0;comment: 点踩数"  `
	PublishedAt       time.Time    `json:"publishedAt" gorm:"column:published_at;comment:发布日期"  `
	TopStatus         int          `json:"status" gorm:"column:top_status;default: 1;comment:状态,1:普通,2:置顶 "  `
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
