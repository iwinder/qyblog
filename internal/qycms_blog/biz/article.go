package biz

import (
	"context"
	"fmt"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/stringUtil"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_DATA_NOT_FOUND.String(), "user not found")
)

// ArticleDO is a ArticleDO model.
type ArticleDO struct {
	metaV1.ObjectMeta
	Title          string
	PermaLink      string
	CanonicalLink  string
	Summary        string
	Thumbnail      string
	Password       string
	Atype          int
	CategoryId     uint64
	CategoryName   string
	CommentAgentId uint64
	Published      bool
	ViewCount      int32
	LikeCount      int32
	HateCount      int32
	Nickname       string
	PublishedAt    time.Time
	TagStrings     []string
	Tags           []*TagsDO
	Content        string
	ContentHtml    string
}

type ArticleDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*ArticleDO `json:"items"`
}

type ArticleDOListOption struct {
	metaV1.ListOptions `json:"page"`
	ArticleDO          `json:"item"`
}

// ArticleRepo is a Greater repo.
type ArticleRepo interface {
	Save(context.Context, *ArticleDO) (*ArticleDO, error)
	Update(context.Context, *ArticleDO) (*ArticleDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*ArticleDO, error)
	FindByAgentID(context.Context, uint64) (*ArticleDO, error)
	CountByPermaLink(ctx context.Context, str string) (int64, error)
	ListAll(context.Context, ArticleDOListOption) (*ArticleDOList, error)
}

// ArticleUsecase   is a ArticleDO usecase.
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
	ac   *ArticleContentUsecase
	at   *ArticleTagsUsecase
	ca   *CommentAgentUsecase
}

// NewArticleUsecase new a ArticleDO usecase.
func NewArticleUsecase(repo ArticleRepo, logger log.Logger,
	ac *ArticleContentUsecase, at *ArticleTagsUsecase, ca *CommentAgentUsecase,
) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger),
		ac: ac, at: at, ca: ca,
	}
}

// Create creates a ArticleDO, and returns the new ArticleDO.
func (uc *ArticleUsecase) Create(ctx context.Context, g *ArticleDO) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("CreateArticle: %v", g.Title)
	// 新增 评论
	cad := &CommentAgentDO{
		ObjType:   int32(g.Atype),
		MemberId:  g.CreatedBy,
		Count:     0,
		RootCount: 0,
		AllCount:  0,
		Attrs:     0,
	}
	cad.CreatedBy = g.CreatedBy
	cado, caErr := uc.ca.CreateCommentAgent(ctx, cad)
	if caErr != nil {
		uc.log.WithContext(ctx).Error("CreateArticle: %v", caErr)
	}
	g.CommentAgentId = cado.ID
	g.Summary = stringUtil.RemoveHtmlAndSubstring(g.ContentHtml)
	data, err := uc.repo.Save(ctx, g)
	if err != nil {
		uc.log.WithContext(ctx).Error("CreateArticle: %v", err)
	}

	// 保存内容
	g.ID = data.ID
	_, cerr := uc.ac.CreateByArticle(ctx, g)
	if cerr != nil {
		uc.log.WithContext(ctx).Error("CreateArticle: %v", cerr)
	}
	// 保存Tags
	uc.at.SaveTagsForArticle(ctx, g)
	// 更新 评论
	cad.UpdatedBy = g.CreatedBy
	cado.ObjId = data.ID
	uc.ca.Update(ctx, cado)
	return data, nil
}

// Update 更新
func (uc *ArticleUsecase) Update(ctx context.Context, g *ArticleDO) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", g.Title)
	g.Summary = stringUtil.RemoveHtmlAndSubstring(g.ContentHtml)
	data, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	// 保存内容
	_, cerr := uc.ac.UpdateByArticle(ctx, g)
	if cerr != nil {
		log.Error(cerr)
	}
	// 保存Tags
	uc.at.SaveTagsForArticle(ctx, g)
	return data, nil
}

// Delete 根据ID删除
func (uc *ArticleUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	err := uc.repo.Delete(ctx, id)
	return err
}

// DeleteList 根据ID批量删除
func (uc *ArticleUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}
func (uc *ArticleUsecase) InitArticlePermaLink(ctx context.Context, title string) string {
	uc.log.WithContext(ctx).Infof("InitArticlePermaLink: %v", title)
	//return uc.repo.DeleteList(ctx, ids)
	link := stringUtil.PinyinConvert(title)
	count, err := uc.repo.CountByPermaLink(ctx, link)
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
	}
	if count > 0 {
		link = fmt.Sprintf("%s-%d", link, count+1)
	}
	return link
}

// FindOneByID 根据ID查询信息
func (uc *ArticleUsecase) FindOneByID(ctx context.Context, id uint64) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	g, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	// 内容
	cont, cerr := uc.ac.FindOneByID(ctx, id)
	if cerr != nil {
		uc.log.WithContext(ctx).Error(cerr)
	}
	g.Content = cont.Content
	g.ContentHtml = cont.ContentHtml

	// 标签
	g.TagStrings = uc.getTagsStringByAid(ctx, id)
	return g, nil
}

// FindOneByAgentID 根据ID查询信息
func (uc *ArticleUsecase) FindOneByAgentID(ctx context.Context, id uint64) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByAgentID: %v", id)
	g, err := uc.repo.FindByAgentID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return g, nil
}

// ListAll 批量查询
func (uc *ArticleUsecase) ListAll(ctx context.Context, opts ArticleDOListOption) (*ArticleDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	for i, _ := range dataDOs.Items {
		dataDOs.Items[i].TagStrings = uc.getTagsStringByAid(ctx, dataDOs.Items[i].ID)
	}

	return dataDOs, nil
}

func (uc *ArticleUsecase) getTagsStringByAid(ctx context.Context, aid uint64) []string {
	oldTsgs, _ := uc.at.FindAllByArticleID(ctx, aid)
	tagStrins := make([]string, 0, len(oldTsgs))
	for _, tag := range oldTsgs {
		tagStrins = append(tagStrins, tag.Name)
	}
	return tagStrins
}
