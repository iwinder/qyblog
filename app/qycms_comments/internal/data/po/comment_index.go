package po

import metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"

// CommentIndexPO 关注楼层
type CommentIndexPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	objId             int64 `json:"objId" gorm:"column:obj_id;default 0;comment: 对象ID "  `
	objType           int   `json:"objType" gorm:"column:obj_type;default 0;comment: 对象类型 "  `
	memberId          int64 `json:"memberId" gorm:"column:member_id;default 0;comment: 发布者用户ID，0 游客 "  `
	rootId            int64 `json:"rootId" gorm:"column:root_id;default 0;comment: 根评论ID，不为0是回复评论 "  `
	parentId          int64 `json:"parentId" gorm:"column:parent_id;default 0;comment: 父级评论ID，为0是 root 评论"  `
	floor             int32 `json:"floor" gorm:"column:floor;default 0;comment: 评论楼层 "  `
	count             int32 `json:"count" gorm:"column:count;default 0;comment:评论总数，不包含已删除 "  `
	root_count        int32 `json:"root_count" gorm:"column:root_count;default 0;comment: 根评论总数"  `
	LikeCount         int32 `json:"likeCount" gorm:"column:like_Count;default 0;comment: 点赞数"  `
	HateCount         int32 `json:"hateCount" gorm:"column:hate_Count;default 0;comment: 点踩数"  `
	attrs             int32 `json:"attrs" gorm:"column:attrs;default 0;comment: 属性预留"  `
}
