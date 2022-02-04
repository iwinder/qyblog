package v1

import "gitee.com/windcoder/qingyucms/internal/qysystem/store"

type Service interface {
	Users() UserSrv
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}