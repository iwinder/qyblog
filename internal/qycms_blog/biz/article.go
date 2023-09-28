package biz

import (
	"context"
	"fmt"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qyblog/internal/pkg/qycms_common/utils/stringUtil"

	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrArticleNotFound is article not found.
	ErrArticleNotFound = errors.NotFound("103404", "article not found")
)

// ArticleDO is a ArticleDO model.
type ArticleDO struct {
	metaV1.ObjectMeta
	Title           string
	PermaLink       string
	CanonicalLink   string
	Summary         string
	Thumbnail       string
	Password        string
	Atype           int
	CategoryId      uint64
	CategoryName    string
	CommentAgentId  uint64
	CommentFlag     bool
	Published       bool
	ViewCount       int32
	CommentCount    int32
	LikeCount       int32
	HateCount       int32
	Nickname        string
	PublishedAt     time.Time
	TagStrings      []string
	Tags            []*TagsDO
	Category        *CategoryDO
	Content         string
	ContentHtml     string
	TagName         string
	Resource        []*ArticleResourceDO
	ResourceStrings []string
	ResourceCount   int64
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
	FindByLink(ctx context.Context, link string) (*ArticleDO, error)
	SetArticleCache(ctx context.Context, user *ArticleDO, key string)
	GetUserFromCache(ctx context.Context, key string) (*ArticleDO, error)
	UpdateCommentContByAgentIds(context.Context) error
	UpdateAllPostsCount(context.Context)
	AddPostViewCount(context.Context, uint64, string)
	GetPostViewCount(context.Context, uint64) int64
}

// ArticleUsecase   is a ArticleDO usecase.
type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
	ac   *ArticleContentUsecase
	at   *ArticleTagsUsecase
	ca   *CommentAgentUsecase
	cu   *CategoryUsecase
	fu   *ArticleResourceUsecase
}

// NewArticleUsecase new a ArticleDO usecase.
func NewArticleUsecase(repo ArticleRepo, logger log.Logger,
	ac *ArticleContentUsecase, at *ArticleTagsUsecase, ca *CommentAgentUsecase,
	cu *CategoryUsecase, fu *ArticleResourceUsecase,
) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: log.NewHelper(logger),
		ac: ac, at: at, ca: ca, cu: cu, fu: fu,
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
		Attrs:     "",
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
	// 新增 相关文件 信息
	uc.fu.SaveResourcesForArticle(ctx, g.ID, g.Resource)
	return data, nil
}

// Update 更新
func (uc *ArticleUsecase) Update(ctx context.Context, g *ArticleDO) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", g.Title)
	g.Summary = stringUtil.RemoveHtmlAndSubstring(g.ContentHtml)
	data, err := uc.repo.Update(ctx, g)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	// 保存内容
	_, cerr := uc.ac.UpdateByArticle(ctx, g)
	if cerr != nil {
		log.Error(cerr)
		return nil, cerr
	}
	// 保存Tags
	terr := uc.at.SaveTagsForArticle(ctx, g)
	if terr != nil {
		return nil, terr
	}

	// 更新 相关文件 信息
	ferr := uc.fu.UpdateResourcesForArticle(ctx, g.ID, g.Resource)
	if ferr != nil {
		return nil, ferr
	}
	return data, nil
}

// Delete 根据ID删除
func (uc *ArticleUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	err := uc.repo.Delete(ctx, id)
	return err
}

// UpdateCommentContByAgentIds 更新评论总数
func (uc *ArticleUsecase) UpdateCommentContByAgentIds(ctx context.Context) error {
	uc.log.WithContext(ctx).Infof("UpdateCommentContByAgentIds")
	err := uc.repo.UpdateCommentContByAgentIds(ctx)
	return err
}
func (uc *ArticleUsecase) UpdateAllPostsCount(ctx context.Context) {
	uc.log.WithContext(ctx).Infof("UpdateAllPostsCount")
	uc.repo.UpdateAllPostsCount(ctx)
}
func (uc *ArticleUsecase) AddPostViewCount(ctx context.Context, id uint64, ip string) {
	uc.log.WithContext(ctx).Infof("UpdateAllPostsCount")
	uc.repo.AddPostViewCount(ctx, id, ip)
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
			return nil, ErrArticleNotFound
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
	// 查询列表
	files, ferr := uc.fu.FindAllByArticleID(ctx, id)
	if ferr != nil {
		uc.log.WithContext(ctx).Error(ferr)
	}
	g.Resource = files
	return g, nil
}
func (uc *ArticleUsecase) FindOneByLink(ctx context.Context, link string) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByLink: %v", link)
	dataDo, err := uc.repo.GetUserFromCache(ctx, link)
	if err != nil {
		g, aerr := uc.repo.FindByLink(ctx, link)
		if aerr != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				dataDo = &ArticleDO{PermaLink: "err404"}
				uc.repo.SetArticleCache(ctx, dataDo, link)
				return nil, ErrArticleNotFound
			} else if err.Error() == "redis: nil" {
				return nil, ErrArticleNotFound
			}
			return nil, err
		}
		// 内容
		cont, cerr := uc.ac.FindOneByID(ctx, g.ID)
		if cerr != nil {
			uc.log.WithContext(ctx).Error(cerr)
		}
		g.ContentHtml = cont.ContentHtml
		g.Content = cont.Content

		// 标签
		tsgs, _ := uc.at.FindAllByArticleID(ctx, g.ID)
		g.Tags = tsgs
		if g.CategoryId > 0 {
			// 分类
			category, _ := uc.cu.FindOneByID(ctx, g.CategoryId)
			g.Category = category
		}

		g.ViewCount = g.ViewCount
		dataDo = g

		// 查询 文件
		files, ferr := uc.fu.FindAllByArticleID(ctx, g.ID)
		if ferr != nil {
			uc.log.WithContext(ctx).Error(ferr)
		}
		g.Resource = files
		uc.repo.SetArticleCache(ctx, dataDo, link)
	}
	if dataDo.PermaLink == "err404" {
		return nil, ErrArticleNotFound
	}

	return dataDo, nil
}
func (uc *ArticleUsecase) GetPostViewCount(ctx context.Context, id uint64) int64 {
	return uc.repo.GetPostViewCount(ctx, id)
}

// FindOneByAgentID 根据ID查询信息
func (uc *ArticleUsecase) FindOneByAgentID(ctx context.Context, id uint64) (*ArticleDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByAgentID: %v", id)
	g, err := uc.repo.FindByAgentID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
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
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	for i, _ := range dataDOs.Items {
		dataDOs.Items[i].TagStrings = uc.getTagsStringByAid(ctx, dataDOs.Items[i].ID)
	}

	return dataDOs, nil
}

func (uc *ArticleUsecase) ListAllForWeb(ctx context.Context, opts ArticleDOListOption) (*ArticleDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
	}
	for i, dataDo := range dataDOs.Items {
		tsgs, _ := uc.at.FindAllByArticleID(ctx, dataDo.ID)
		dataDOs.Items[i].Tags = tsgs
		// 分类
		category, _ := uc.cu.FindOneByID(ctx, dataDo.CategoryId)
		dataDOs.Items[i].Category = category
	}

	return dataDOs, nil
}
func (uc *ArticleUsecase) GeneratorMapListAll(ctx context.Context, opts ArticleDOListOption) (*ArticleDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	dataDOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, err
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
