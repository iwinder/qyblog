package biz

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/emailUtil"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/fileUtil"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"html/template"
	"strings"
)

type CommentDO struct {
	metaV1.ObjectMeta
	AgentId        uint64
	MemberId       uint64
	AtMemberIds    string
	Agent          string
	MemberName     string
	Ip             string
	Email          string
	Url            string
	RootId         uint64
	ParentId       uint64
	Content        string
	Meta           string
	ParentUserName string
	ObjTitle       string
	ObjLink        string
	Avatar         string
	EmailState     int32
}

type CommentDOList struct {
	metaV1.ListMeta
	Agent *CommentAgentDO
	Items []*CommentContentDO
}

type CommentUsecase struct {
	log  *log.Helper
	au   *ArticleUsecase
	ca   *CommentAgentUsecase
	ci   *CommentIndexUsecase
	cc   *CommentContentUsecase
	uu   *UserUsecase
	site *SiteConfigUsecase
}

func NewCommentUsecase(logger log.Logger, ca *CommentAgentUsecase,
	ci *CommentIndexUsecase, uu *UserUsecase, au *ArticleUsecase,
	cc *CommentContentUsecase, site *SiteConfigUsecase,
) *CommentUsecase {
	return &CommentUsecase{log: log.NewHelper(logger),
		au: au, ca: ca, ci: ci, cc: cc, uu: uu, site: site}
}

// CreateComment 新增评论
func (uc *CommentUsecase) CreateComment(ctx context.Context, g *CommentDO) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateComment: %v-%v", g.MemberId, g.MemberName)
	// 如果有父级，先查询父级
	g.RootId = 0
	if g.ParentId > 0 {
		parent, err := uc.ci.FindByID(ctx, g.ParentId)
		if err != nil {
			log.Error(err)
		}
		// 其父级的根目录为0，则父级为根评论
		if parent.RootId == 0 {
			g.RootId = parent.ID
			// TODO: 计算改为定时任务
		} else {
			g.RootId = parent.RootId
		}
	}

	// 创建index
	ci := &CommentIndexDO{
		AgentId:  g.AgentId,
		MemberId: g.MemberId,
		RootId:   g.RootId,
		ParentId: g.ParentId,
	}
	ci.StatusFlag = g.StatusFlag
	cidata, cierr := uc.ci.CreateCommentIndex(ctx, ci)
	if cierr != nil {
		log.Error(cierr)
	}
	// 创建内容
	cc := &CommentContentDO{
		AgentId:     g.AgentId,
		MemberId:    g.MemberId,
		AtMemberIds: "",
		Agent:       g.Agent,
		MemberName:  g.MemberName,
		Ip:          g.Ip,
		Email:       g.Email,
		Url:         g.Url,
		RootId:      g.RootId,
		Content:     g.Content,
		EmailState:  g.EmailState,
	}
	cc.StatusFlag = g.StatusFlag
	if cidata != nil {
		cc.ID = cidata.ID
	}
	data, err := uc.cc.CreateCommentContent(ctx, cc)
	if err != nil {
		return nil, err
	}
	g.ID = data.ID
	if cc.StatusFlag == 1 {
		// 更新计数
		uc.ca.UpdateAddCountById(ctx, g.AgentId, g.RootId == 0)
		if g.ParentId > 0 {
			uc.ci.UpdateAddCountById(ctx, g.ParentId, g.RootId == 0)
		}
	}
	return data, nil
}

// CreateCommentWeb 新增评论
func (uc *CommentUsecase) CreateCommentWeb(ctx context.Context, g *CommentDO, conf *conf.Qycms) (*CommentContentDO, error) {
	uc.log.WithContext(ctx).Infof("CreateComment: %v-%v", g.MemberId, g.MemberName)
	data, err := uc.CreateComment(ctx, g)
	if err != nil {
		return nil, err
	}
	if g.StatusFlag != 1 {
		// 推送待审核邮件
		url := uc.site.FindValueByKey(ctx, "site_url")
		siteName := uc.site.FindValueByKey(ctx, "site_name")
		logo := uc.site.FindValueByKey(ctx, "site_logo")
		templatePath := conf.DocPath + fileUtil.EmailTemplatePath
		subject := fmt.Sprintf("您的 %s 博客有新的待审核留言", siteName)
		var content strings.Builder
		content.WriteString(g.MemberName)
		content.WriteString("：")
		content.WriteString(stringUtil.RemoveHtmlAndSubstring(g.Content))
		t, _ := template.ParseFiles(templatePath + "pending-template.html")
		var body bytes.Buffer

		t.Execute(&body, struct {
			LOGO         string
			SITENAME     string
			SITEURL      string
			REPLAYCOMENT string
		}{
			LOGO:         logo,
			SITENAME:     siteName,
			SITEURL:      url,
			REPLAYCOMENT: content.String(),
		})
		err = emailUtil.SendMail(conf.Email.Username, conf.Email.Password,
			conf.Email.Host, conf.Email.Port, conf.Email.SenderName, conf.Email.AdminEMail,
			subject, body.String())
		if err != nil {
			uc.log.Error(fmt.Errorf("待审核邮件发送失败: %w", err))
		}
	}
	return data, nil
}

// UpdateCommentComent 更新评论内容
func (uc *CommentUsecase) UpdateCommentComent(ctx context.Context, g *CommentDO) error {
	return uc.cc.UpdaeCommentById(ctx, g.ID, g.Content)
}

// UpdateCommentState 更新状态
func (uc *CommentUsecase) UpdateCommentState(ctx context.Context, ids []uint64, state int) error {
	err := uc.cc.UpdateStateByIDs(ctx, ids, state)
	if err != nil {
		return err
	}
	err = uc.ci.UpdateStateByIDs(ctx, ids, state)
	if err != nil {
		log.Error(err)
	}

	for _, id := range ids {
		//
		idx, _ := uc.ci.FindByID(ctx, id)
		if idx != nil {
			// 更新
			if state == 1 {
				uc.UpdateAddCount(ctx, idx)
			} else {
				uc.UpdateMinusCount(ctx, idx)
			}
		}
	}
	return nil
}
func (uc *CommentUsecase) UpdateAddCount(ctx context.Context, idx *CommentIndexDO) {
	uc.ca.UpdateAddCountById(ctx, idx.AgentId, idx.RootId == 0)
	if idx.ParentId > 0 {
		uc.ci.UpdateAddCountById(ctx, idx.ParentId, idx.RootId == 0)
	}
}
func (uc *CommentUsecase) UpdateMinusCount(ctx context.Context, idx *CommentIndexDO) {
	uc.ca.UpdateMinusCountById(ctx, idx.AgentId, idx.RootId == 0)
	if idx.ParentId > 0 {
		uc.ci.UpdateMinusCountById(ctx, idx.ParentId, idx.RootId == 0)
	}
}

// UpdateContentCountAndObjIds 更新文章的评论数和评论的objId等
func (uc *CommentUsecase) UpdateContentCountAndObjIds(ctx context.Context) error {
	err := uc.au.UpdateCommentContByAgentIds(ctx)
	if err != nil {
		return err
	}
	err = uc.ci.UpdateObjIdByAgentIds(ctx)
	return err
}
func (uc *CommentUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	err := uc.cc.DeleteList(ctx, ids)
	if err != nil {
		return err
	}
	err = uc.ci.DeleteList(ctx, ids)
	if err != nil {
		log.Error(err)
	}
	for _, id := range ids {
		idx, _ := uc.ci.FindByID(ctx, id)
		if idx != nil {
			// 更新
			uc.UpdateMinusCount(ctx, idx)
		}
	}
	return err
}

func (uc *CommentUsecase) EmailToNotSend(ctx context.Context, conf *conf.Qycms) {
	opts := CommentContentDOListOption{}
	opts.Current = 1
	opts.IsWeb = false
	opts.EmailState = 1
	opts.PageSize = 10
	opts.StatusFlag = 1
	opts.ListOptions.Init()
	dataList := uc.ListAllNeedEmail(ctx, opts)

	url := uc.site.FindValueByKey(ctx, "site_url")
	siteName := uc.site.FindValueByKey(ctx, "site_name")
	logo := uc.site.FindValueByKey(ctx, "site_logo")
	siteInfo := make(map[string]string, 0)
	siteInfo["URL"] = url
	siteInfo["NAME"] = siteName
	siteInfo["LOGO"] = logo
	uc.EmailToSend(ctx, dataList.Items, conf, siteInfo)
	opts.TotalCount = dataList.TotalCount
	opts.ListOptions.IsLast()
	totalPage := opts.Pages
	var page int64
	for page = 2; page <= totalPage; page++ {
		opts.Current = page
		opts.ListOptions.Init()
		dataList = uc.ListAllNeedEmail(ctx, opts)
		uc.EmailToSend(ctx, dataList.Items, conf, siteInfo)
	}

}

func (uc *CommentUsecase) EmailToSend(ctx context.Context, items []*CommentContentDO, conf *conf.Qycms, siteInfo map[string]string) {
	name, _ := siteInfo["NAME"]
	url, _ := siteInfo["URL"]
	logo, _ := siteInfo["LOGO"]
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	templatePath := conf.DocPath + fileUtil.EmailTemplatePath
	var sate int32 = 2

	for _, item := range items {

		if item.Children != nil && len(item.Children) > 0 {
			//仅发送有回复记录的
			subject := fmt.Sprintf("您在 %s 博客的留言有了回复", name)
			replyList := make([]string, 0, len(item.Children))
			var content strings.Builder
			for _, obj := range item.Children {
				content.WriteString(obj.MemberName)
				content.WriteString("：")
				content.WriteString(stringUtil.RemoveHtmlAndSubstring(obj.Content))
				replyList = append(replyList, content.String())
			}
			t, _ := template.ParseFiles(templatePath + "replay-template.html")
			var body bytes.Buffer

			t.Execute(&body, struct {
				LOGO             string
				SITENAME         string
				TOUSERNAME       string
				POSTURL          string
				POSTNAME         string
				REPLAYCOMENTLIST []string
			}{
				LOGO:             logo,
				SITENAME:         name,
				TOUSERNAME:       item.MemberName,
				POSTURL:          url + item.ObjLink,
				POSTNAME:         item.ObjTitle,
				REPLAYCOMENTLIST: replyList,
			})
			err := emailUtil.SendMail(conf.Email.Username, conf.Email.Password,
				conf.Email.Host, conf.Email.Port, conf.Email.SenderName, item.Email,
				subject, body.String())
			sate = 2
			if err != nil {
				uc.log.Error(fmt.Errorf("邮件发送失败: %w", err))
				sate = 3
			}
			uc.cc.UpdaeEmailStateById(ctx, item.ID, sate)

		}
	}
}

func (uc *CommentUsecase) ListAllNeedEmail(ctx context.Context, opts CommentContentDOListOption) *CommentDOList {
	result := &CommentDOList{}
	parentList, err := uc.cc.ListAll(ctx, opts)
	if err != nil {
		log.Error(err)
		result.Items = make([]*CommentContentDO, 0, 0)
		return result
	}
	result.ListMeta = parentList.ListMeta
	result.Items = parentList.Items
	for i, data := range result.Items {
		child, cerr := uc.cc.FindAllByParentID(ctx, data.ID, 5)
		if cerr != nil {
			log.Error(err)
			result.Items[i].Children = make([]*CommentContentDO, 0, 0)
			continue
		}
		result.Items[i].Children = child
	}
	return result
}

func (uc *CommentUsecase) ListAllForWeb(ctx context.Context, opts CommentContentDOListOption) *CommentDOList {
	result := &CommentDOList{Agent: &CommentAgentDO{Count: 0}}
	if !opts.IsChild && opts.Current == 1 {
		agent, err := uc.ca.FindByID(ctx, opts.AgentId)
		if err != nil {
			log.Error(err)
			result.Agent = &CommentAgentDO{}
			result.Items = make([]*CommentContentDO, 0, 0)
			return result
		}
		result.Agent = agent
	}
	objList, err := uc.cc.ListAll(ctx, opts)
	if err != nil {
		log.Error(err)
		result.Items = make([]*CommentContentDO, 0, 0)
		return result
	}
	result.Items = objList.Items
	result.ListMeta = objList.ListMeta
	return result
}
