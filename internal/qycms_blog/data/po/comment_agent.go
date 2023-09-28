package po

import (
	"encoding/json"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

// CommentAgentPO 主题表
type CommentAgentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	ObjId             uint64 `json:"objId" gorm:"column:obj_id;default: 0;comment: 对象ID "  `
	ObjType           int32  `json:"objType" gorm:"column:obj_type;default: 0;comment: 对象类型 "  `
	MemberId          uint64 `json:"memberId" gorm:"column:member_id;default: 0;comment: 作者ID "  `
	Count             int32  `json:"count" gorm:"column:count;default: 0;comment:评论总数，不包含已删除 "  `
	RootCount         int32  `json:"root_count" gorm:"column:root_count;default: 0;comment: 根评论总数"  `
	AllCount          int32  `json:"all_count" gorm:"column:all_count;default: 0;comment: 评论+回复总数 "  `
	//State             int8   `json:"state" gorm:"column:state;default: 0;comment::状态,是否开启：0不开启， 1开启  "  `
	Attrs string `json:"attrs" gorm:"column:attrs;default: 0;comment: 属性预留"  `
}

type CommentAgentPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CommentAgentPO `json:"items"`
}

type CommentAgentPOListOption struct {
	metaV1.ListOptions `json:"page"`
	CommentAgentPO     `json:"item"`
}

func (o *CommentAgentPO) TableName() string {
	return "qy_blog_comment_agent"
}

func (o *CommentAgentPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *CommentAgentPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "comment-agent-")
	if o.Sort <= 0 {
		o.Sort = int(o.ID)
	}
	return tx.Save(o).Error
}

func (o *CommentAgentPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *CommentAgentPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
