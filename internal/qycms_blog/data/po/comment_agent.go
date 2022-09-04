package po

import metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"

// CommentAgentPO 主题表
type CommentAgentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	ObjId             uint64 `json:"objId" gorm:"column:obj_id;default: 0;comment: 对象ID "  `
	ObjType           int32  `json:"objType" gorm:"column:obj_type;default: 0;comment: 对象类型 "  `
	MemberId          uint64 `json:"memberId" gorm:"column:member_id;default: 0;comment: 作者ID "  `
	Count             int32  `json:"count" gorm:"column:count;default: 0;comment:评论总数，不包含已删除 "  `
	RootCount         int32  `json:"root_count" gorm:"column:root_count;default: 0;comment: 根评论总数"  `
	AllCount          int32  `json:"all_count" gorm:"column:all_count;default: 0;comment: 评论+回复总数 "  `
	State             int8   `json:"state" gorm:"column:state;default: 0;comment::状态,是否开启：0不开启， 1开启  "  `
	Attrs             int32  `json:"attrs" gorm:"column:attrs;default: 0;comment: 属性预留"  `
}
