package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/iwinder/qingyucms/api/qycms_user/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
)

var (
	ErrAuthFailed = errors.New("authentication failed")
)

// UserService is a greeter service. 可用于后期拆分单个微服务
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

func (s *UserService) GetMyInfo(ctx context.Context, in *v1.GetMyInfoRequest) (*v1.GetUserReply, error) {
	var uId uint64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwtV4.MapClaims)
		if c["ID"] == nil {
			return nil, ErrAuthFailed
		}
		uId = uint64(c["ID"].(float64))
	}

	user, err := s.uc.FindOneByID(ctx, uId)
	if err != nil {
		return nil, err
	}
	u := UserResponse(user)
	return &v1.GetUserReply{User: &u}, nil
}

func (s *UserService) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	opts := biz.UserDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = -1
	opts.ListOptions.PageSize = 20
	if in.PageInfo != nil {
		opts.ListOptions.Pages = int64(in.PageInfo.Pages)
		opts.ListOptions.Current = int64(in.PageInfo.Page)
		opts.ListOptions.PageSize = int64(in.PageInfo.Size)
	}

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

//func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
//	rv, err := s.uc.VerifyPassword(ctx, &biz.UserDO{Username: req.Username, Password: req.Password, Salt: s.conf.Token})
//	if err != nil {
//		return nil, err
//	}
//
//	return &v1.VerifyPasswordReply{
//		Ok: rv,
//	}, nil
//}

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
