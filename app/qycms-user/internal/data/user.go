package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/iwinder/qingyucms/app/qycms-user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, userBiz *biz.UserBiz) (*biz.UserBiz, error) {
	err := r.data.db.Create(&userBiz).Error
	if err != nil {

	}
	return userBiz, nil
}

func (r *userRepo) Update(ctx context.Context, g *biz.UserBiz) (*biz.UserBiz, error) {
	return g, nil
}

func (r *userRepo) FindByID(c context.Context, id int64) (*biz.UserBiz, error) {
	user := &biz.UserBiz{}
	err := r.data.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%d,data not found,error is [%s]:", 111, err)
		}

		return nil, fmt.Errorf("%d,data query error,error is [%s]:", 222, err)
	}

	return user, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.UserBiz, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.UserBiz, error) {
	return nil, nil
}
