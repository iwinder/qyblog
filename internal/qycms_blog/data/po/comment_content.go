package po

import (
	"database/sql"
	"encoding/json"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/idUtil"
	"gorm.io/gorm"
)

// CommentContentPO 关注内容,id 与 CommentIndexPO 相同
// StatusFlag 0 全部 1 已审核 2 待审核
type CommentContentPO struct {
	metaV1.ObjectMeta `json:"metadata,omitempty"`
	MemberId          uint64        `json:"memberId" gorm:"column:member_id;default: 0;comment: 评论者ID，0 为游客 "  `
	AgentId           uint64        `json:"agentId" gorm:"column:agent_id;default: 0;comment: 主题表ID "  `
	AtMemberIds       string        `json:"atMemberIds" gorm:"column:at_member_id;default: 0;comment: 圈的人员ID列表，使用英文逗号隔开"  `
	Agent             string        `json:"agent" gorm:"column:agent;type:text;comment: 评论者客户端"  `
	MemberName        string        `json:"memberName" gorm:"column:member_name;comment: 评论者用户名"  `
	Ip                string        `json:"ip" gorm:"column:ip;default: 0;comment:评论者ip"  `
	Email             string        `json:"email" gorm:"column:email;comment:评论者邮箱"  `
	Url               string        `json:"url" gorm:"column:url;comment:评论者网址"  `
	RootId            sql.NullInt64 `json:"rootId" gorm:"column:root_id;default: 0;comment: 根评论ID，不为0是回复评论 "  `
	Content           string        `json:"content" gorm:"column:content;type:longtext;comment: 评论内容"  `
	Meta              string        `json:"attrs" gorm:"column:attrs;comment: 属性预留"  `
}

type CommentContentPOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*CommentContentPO `json:"items"`
}

type CommentContentPOListOption struct {
	metaV1.ListOptions `json:"page"`
	CommentContentPO   `json:"item"`
}

func (o *CommentContentPO) TableName() string {
	return "qy_blog_comment_content"
}

func (o *CommentContentPO) BeforeCreate(tx *gorm.DB) (er error) {
	return
}

func (o *CommentContentPO) AfterCreate(tx *gorm.DB) (err error) {
	o.InstanceID = idUtil.GetInstanceID(o.ID, "comment-content-")
	if o.Sort <= 0 {
		o.Sort = int(o.ID)
	}
	return tx.Save(o).Error
}

func (o *CommentContentPO) BeforeUpdate(tx *gorm.DB) (err error) {
	o.ExtendShadow = o.Extend.String()
	return err
}

func (o *CommentContentPO) AfterFind(tx *gorm.DB) (err error) {
	if err := json.Unmarshal([]byte(o.ExtendShadow), &o.Extend); err != nil {
		return err
	}
	return nil
}
