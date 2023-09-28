package po

import (
	"database/sql"
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

// CommentIndexPO 关注楼层
type CommentIndexPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	AgentId           uint64        `json:"agentId" gorm:"column:agent_id;default: 0;comment: 主题表ID "  `
	ObjId             uint64        `json:"objId" gorm:"column:obj_id;default:0;comment: 对象ID "  `
	ObjType           int32         `json:"objType" gorm:"column:obj_type;default:0;comment: 对象类型 "  `
	Atype             int32         `json:"Atype" gorm:"column:atype;default:1;comment: 类型 1:正常，2：通知 "  `
	MemberId          sql.NullInt64 `json:"memberId" gorm:"column:member_id;default:0;comment: 发布者用户ID，0 游客 "  `
	RootId            sql.NullInt64 `json:"rootId" gorm:"column:root_id;default:0;comment: 根评论ID，不为0是回复评论 "  `
	ParentId          sql.NullInt64 `json:"parentId" gorm:"column:parent_id;default: 0;comment: 父级评论ID，为0是 root 评论"  `
	Floor             int32         `json:"floor" gorm:"column:floor;default: 0;comment: 评论楼层，楼中楼不算层数 "  `
	Count             int32         `json:"count" gorm:"column:count;default: 0;comment:评论总数，不包含已删除 "  `
	RootCount         int32         `json:"rootCount" gorm:"column:root_count;default: 0;comment: 根评论总数"  `
	LikeCount         int32         `json:"likeCount" gorm:"column:like_Count;default: 0;comment: 点赞数"  `
	HateCount         int32         `json:"hateCount" gorm:"column:hate_Count;default: 0;comment: 点踩数"  `
	Attrs             string        `json:"attrs" gorm:"column:attrs;default: 0;comment: 属性预留"  `
}

type CommentIndexPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CommentIndexPO `json:"items"`
}

type CommentIndexPOListOption struct {
	metaV1.ListOptions `json:"page"`
	CommentIndexPO     `json:"item"`
}

func (o *CommentIndexPO) TableName() string {
	return "qy_blog_comment_index"
}

func (o *CommentIndexPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *CommentIndexPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "comment-index-")
	if o.Sort <= 0 {
		o.Sort = int(o.ID)
	}
	return tx.Save(o).Error
}

func (o *CommentIndexPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *CommentIndexPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
