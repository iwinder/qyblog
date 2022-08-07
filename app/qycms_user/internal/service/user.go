package service

import (
	"context"
	"github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/biz"
	"github.com/iwinder/qingyucms/app/qycms_user/internal/conf"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer
	conf *conf.Qycms
	uc   *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, conf *conf.Qycms) *UserService {
	return &UserService{uc: uc, conf: conf}
}

// CreateUser implements server.CreateUser. 创建用户
func (s *UserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	userDO := &biz.UserDO{
		Username: in.Username,
		Nickname: in.NickName,
		Password: in.Password,
		Salt:     s.conf.Token,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Phone:    in.Phone,
	}
	user, err := s.uc.CreateUser(ctx, userDO)
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.CreateUserReply{User: &u}, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	userDO := &biz.UserDO{
		Username: in.Username,
		Nickname: in.NickName,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Phone:    in.Phone,
	}
	user, err := s.uc.Update(ctx, userDO)
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.UpdateUserReply{User: &u}, nil
}

// DeleteUser 根据ID删除用户
func (s *UserService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.uc.Delete(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{}, nil
}

// DeleteUsers 根据ID批量删除用户
func (s *UserService) DeleteUsers(ctx context.Context, in *v1.DeleteUsersRequest) (*v1.DeleteUsersReply, error) {
	err := s.uc.DeleteList(ctx, in.Uids)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUsersReply{}, nil
}

// GetUser 通过ID获取用户
func (s *UserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.FindOneByID(ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.GetUserReply{User: &u}, nil
}

func (s *UserService) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	opts := biz.UserListOption{}
	opts.ListOptions.Pages = int64(in.PageInfo.Pages)
	opts.ListOptions.Page = int64(in.PageInfo.Page)
	opts.ListOptions.PageSize = int64(in.PageInfo.Size)
	opts.ListOptions.Init()
	userList, err := s.uc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.PageInfo{
		Page:      uint64(userList.Pages),
		Size:      uint64(userList.PageSize),
		Total:     uint64(userList.TotalCount),
		Pages:     uint64(userList.Pages),
		FirstFlag: userList.FirstFlag,
		LastFlag:  userList.LastFlag,
	}
	users := make([]*v1.UserInfoResponse, 0, len(userList.Items))
	for _, user := range userList.Items {
		users = append(users, &v1.UserInfoResponse{
			Uid:      user.ID,
			Username: user.Username,
			NickName: user.Nickname,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Phone:    user.Phone,
		})
	}
	return &v1.ListUserReply{PageInfo: pageInfo, Items: users}, nil
}

func UserResponse(user *biz.UserDO) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Uid:      user.ID,
		Username: user.Username,
		NickName: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
	}

	return userInfoRsp
}
