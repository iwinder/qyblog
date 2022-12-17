package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type ArticleVisitorDO struct {
	ID        uint64
	ArticleId uint64
	Ip        string
	Agent     string
	Atype     int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type ArticleVisitorRepo interface {
	Save(context.Context, *ArticleVisitorDO) (*ArticleVisitorDO, error)
}

type ArticleVisitorUsecase struct {
	repo ArticleVisitorRepo
	log  *log.Helper
}

func NewArticleVisitorUsecase(repo ArticleVisitorRepo, logger log.Logger) *ArticleVisitorUsecase {
	return &ArticleVisitorUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *ArticleVisitorUsecase) Create(ctx context.Context, g *ArticleVisitorDO) (*ArticleVisitorDO, error) {
	uc.log.WithContext(ctx).Infof("CreateArticle: %v", g.Ip)
	objDO, err := uc.repo.Save(ctx, g)
	if err != nil {
		uc.log.WithContext(ctx).Error(fmt.Errorf("ArticleVisitor %s 浏览记录保存失败:%w", g.Ip, err))
		return nil, err
	}

	return objDO, nil
}
