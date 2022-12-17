package db

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/data/po"
)

type articleVisitorRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleVisitorRepo(data *Data, logger log.Logger) biz.ArticleVisitorRepo {
	return &articleVisitorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *articleVisitorRepo) Save(ctx context.Context, do *biz.ArticleVisitorDO) (*biz.ArticleVisitorDO, error) {
	po := &po.ArticleVisitorPO{
		ArticleId: do.ArticleId,
		Ip:        do.Ip,
		Agent:     do.Agent,
		Atype:     do.Atype,
	}
	err := r.data.Db.Create(po).Error
	if err != nil {
		return nil, err
	}
	dataDO := &biz.ArticleVisitorDO{ID: po.ID}
	return dataDO, nil
}
