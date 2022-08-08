package po

import metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"

// CommentAgentPO 主题表
type CommentAgentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	objId             int64 `json:"objId" gorm:"column:obj_id;default 0;comment: 对象ID "  `
	objType           int   `json:"objType" gorm:"column:obj_type;default 0;comment: 对象类型 "  `
	memberId          int64 `json:"memberId" gorm:"column:member_id;default 0;comment: 作者ID "  `
	count             int32 `json:"count" gorm:"column:count;default 0;comment:评论总数，不包含已删除 "  `
	root_count        int32 `json:"root_count" gorm:"column:root_count;default 0;comment: 根评论总数"  `
	all_count         int32 `json:"all_count" gorm:"column:all_count;default 0;comment: 评论+回复总数 "  `
	state             int8  `json:"state" gorm:"column:state;default 0;comment::状态,是否开启：0不开启， 1开启  "  `
	attrs             int32 `json:"attrs" gorm:"column:attrs;default 0;comment: 属性预留"  `
}
