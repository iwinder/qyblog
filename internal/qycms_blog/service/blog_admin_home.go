package service

import (
	"context"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
)

func (s *BlogAdminUserService) UpdateContentCountJobsQyAdminHome(ctx context.Context, request *v1.CreateQyAdminHomeRequest) (*v1.CreateQyAdminHomeReply, error) {
	err := s.ctu.UpdateContentCountAndObjIds(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminHomeReply{}, nil
}
func (s *BlogAdminUserService) GeneratorMapJobQyAdminHome(ctx context.Context, request *v1.CreateQyAdminHomeRequest) (*v1.CreateQyAdminHomeReply, error) {
	err := s.siteMap.GeneratorMap(ctx, s.conf.SiteMapPath)
	if err != nil {
		return nil, err
	}
	return &v1.CreateQyAdminHomeReply{}, nil
}

func (s *BlogAdminUserService) UpdateAllPostsCountJobsQyAdminHome(ctx context.Context, request *v1.CreateQyAdminHomeRequest) (*v1.CreateQyAdminHomeReply, error) {
	s.au.UpdateAllPostsCount(ctx)
	return &v1.CreateQyAdminHomeReply{}, nil
}
func (s *BlogAdminUserService) EmailToNotSendCountJobsQyAdminHome(ctx context.Context, request *v1.CreateQyAdminHomeRequest) (*v1.CreateQyAdminHomeReply, error) {
	s.ctu.EmailToNotSend(ctx, s.conf)
	return &v1.CreateQyAdminHomeReply{}, nil
}
