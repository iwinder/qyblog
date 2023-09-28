package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	metaV1 "github.com/iwinder/qyblog/internal/pkg/qycms_common/meta/v1"
)

type ArticleResourceDO struct {
	metaV1.ObjectMeta
	ArticleID uint64
	Name      string
	Url       string
	Password  string
}

type ArticleResourceRepo interface {
	CreateInBatches(ctx context.Context, articleID uint64, articleResource []*ArticleResourceDO) error
	UpdateInBatches(ctx context.Context, articleResource []*ArticleResourceDO) error
	FindAllByArticleID(ctx context.Context, articleId uint64) ([]*ArticleResourceDO, error)
	FindAllByArticlePermaLink(ctx context.Context, permaLink string) ([]*ArticleResourceDO, error)
	DeleteByArticleID(ctx context.Context, articleId uint64) error
	DeleteByIDs(ctx context.Context, ids []uint64) error
	CountByArticleID(ctx context.Context, articleId uint64) (int64, error)
}

type ArticleResourceUsecase struct {
	repo ArticleResourceRepo
	log  *log.Helper
}

// NewArticleResourceUsecase new a ArticleResourceDO usecase.
func NewArticleResourceUsecase(repo ArticleResourceRepo, logger log.Logger) *ArticleResourceUsecase {
	return &ArticleResourceUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ArticleResourceUsecase) SaveResourcesForArticle(ctx context.Context, articleID uint64, articleResources []*ArticleResourceDO) error {

	err := uc.repo.CreateInBatches(ctx, articleID, articleResources)
	if err != nil {
		uc.log.WithContext(ctx).Error("SaveResourcesForArticle: %v", err)
		return err
	}
	return nil
}

func (uc *ArticleResourceUsecase) UpdateResourcesForArticle(ctx context.Context, articleID uint64, articleResources []*ArticleResourceDO) error {
	// 查找当前文章附件数量
	oldList, err := uc.FindAllByArticleID(ctx, articleID)
	if err != nil {
		uc.log.WithContext(ctx).Error("FindAllByArticleID: %v", err)
		return err
	}
	oldLen := len(oldList)
	nowLen := len(articleResources)

	if oldLen == 0 {
		// 没有直接新增
		if nowLen > 0 {
			uc.SaveResourcesForArticle(ctx, articleID, articleResources)
		}
	} else if oldLen <= nowLen {
		// 小于等于新列表时 更新已有
		newLen := nowLen - oldLen
		uerr := uc.UpdateResourcesByOldFiles(ctx, articleID, oldLen, oldList, articleResources)
		if uerr != nil {
			uc.log.WithContext(ctx).Error("UpdateInBatches: ", err)
			return uerr
		}
		// 再新增未有
		if newLen > 0 {
			newFiles := make([]*ArticleResourceDO, 0, newLen)
			for i := oldLen; i < nowLen; i++ {
				newFiles = append(newFiles, &ArticleResourceDO{
					ArticleID: articleID,
					Name:      articleResources[i].Name,
					Url:       articleResources[i].Url,
					Password:  articleResources[i].Password,
				})
			}
			serr := uc.SaveResourcesForArticle(ctx, articleID, newFiles)
			if serr != nil {
				uc.log.WithContext(ctx).Error("SaveNewFilesForArticle: ", serr)
				return serr
			}
		}
	} else {
		// 如果新的为空，直接全删
		if nowLen == 0 {
			delErr := uc.repo.DeleteByArticleID(ctx, articleID)
			if delErr != nil {
				uc.log.WithContext(ctx).Error("DeleteByArticleID: ", delErr)
				return delErr
			}
		}
		// 大于新列表时，先更新
		newLen := oldLen - nowLen
		uerr := uc.UpdateResourcesByOldFiles(ctx, articleID, nowLen, oldList, articleResources)
		if uerr != nil {
			uc.log.WithContext(ctx).Error("UpdateInBatches: ", err)
			return uerr
		}
		// 再删除
		delFileIDs := make([]uint64, 0, newLen)
		for i := nowLen; i < oldLen; i++ {
			delFileIDs = append(delFileIDs, oldList[i].ID)
		}

		derr := uc.repo.DeleteByIDs(ctx, delFileIDs)
		if derr != nil {
			uc.log.WithContext(ctx).Error("articleResources DeleteByIDs: ", derr)
			return derr
		}
	}
	return nil
}
func (uc *ArticleResourceUsecase) UpdateResourcesByOldFiles(ctx context.Context, articleID uint64, needLen int,
	oldList []*ArticleResourceDO, articleResources []*ArticleResourceDO) error {
	nowFiles := make([]*ArticleResourceDO, 0, needLen)
	for i := 0; i < needLen; i++ {
		nowFiles = append(nowFiles, &ArticleResourceDO{
			ObjectMeta: metaV1.ObjectMeta{
				ID: oldList[i].ID,
			},
			ArticleID: articleID,
			Name:      articleResources[i].Name,
			Url:       articleResources[i].Url,
			Password:  articleResources[i].Password,
		})
	}
	err := uc.repo.UpdateInBatches(ctx, nowFiles)
	return err
}
func (uc *ArticleResourceUsecase) FindAllByArticleID(ctx context.Context, aid uint64) ([]*ArticleResourceDO, error) {
	oldFiles, err := uc.repo.FindAllByArticleID(ctx, aid)
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
	}
	return oldFiles, nil
}

func (uc *ArticleResourceUsecase) FindAllByArticlePermaLink(ctx context.Context, permaLink string) ([]*ArticleResourceDO, error) {
	oldFiles, err := uc.repo.FindAllByArticlePermaLink(ctx, permaLink)
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
	}
	return oldFiles, nil
}

func (uc *ArticleResourceUsecase) CountByArticleID(ctx context.Context, articleId uint64) int64 {
	num, err := uc.repo.CountByArticleID(ctx, articleId)
	if err != nil {
		uc.log.WithContext(ctx).Warn(err)
	}
	return num
}
