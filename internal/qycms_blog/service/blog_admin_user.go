package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	v1 "github.com/iwinder/qingyucms/api/qycms_bff/admin/v1"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/qycms_blog/biz"
)

var (
	ErrAuthFailed = errors.New("authentication failed")
)

// CreateUser 创建用户
func (s *BlogAdminUserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	userDO := &biz.UserDO{
		Username: in.Username,
		Nickname: in.Nickname,
		Password: in.Password,
		Salt:     s.conf.Token,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Phone:    in.Phone,
	}
	userDO.StatusFlag = int(in.StatusFlag)
	user, err := s.uc.CreateUser(ctx, userDO)
	if err != nil {
		return nil, err
	}
	u := bizToUserResponse(user)
	return &v1.CreateUserReply{Data: &u}, nil
}

// UpdateUser 更新用户
func (s *BlogAdminUserService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	userDO := &biz.UserDO{
		Username: in.Username,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Phone:    in.Phone,
	}
	userDO.ID = in.Id
	userDO.StatusFlag = int(in.StatusFlag)
	if in.Roles != nil && len(in.Roles) > 0 {
		roles := make([]*biz.RoleDO, 0, len(in.Roles))
		for _, obj := range in.Roles {
			roles = append(roles, &biz.RoleDO{
				ObjectMeta: metaV1.ObjectMeta{
					ID: obj.Id,
				},
				Name:       obj.Name,
				Identifier: obj.Identifier,
			})
		}
		userDO.Roles = roles
	}
	user, err := s.uc.Update(ctx, userDO)
	if err != nil {
		return nil, err
	}
	u := bizToUserResponse(user)
	return &v1.UpdateUserReply{Data: &u}, nil
}

// DeleteUser 根据ID删除用户
func (s *BlogAdminUserService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.uc.Delete(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{}, nil
}

// DeleteUsers 根据ID批量删除用户
func (s *BlogAdminUserService) DeleteUsers(ctx context.Context, in *v1.DeleteUsersRequest) (*v1.DeleteUsersReply, error) {

	err := s.uc.DeleteList(ctx, in.Ids)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteUsersReply{}, nil
}

// GetUser 根据ID获取用户信息
func (s *BlogAdminUserService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.FindOneByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	u := bizToUserResponse(user)
	return &v1.GetUserReply{Data: &u}, nil
}

// GetMyInfo 获取用户个人信息
func (s *BlogAdminUserService) GetMyInfo(ctx context.Context, in *v1.GetMyInfoRequest) (*v1.GetUserReply, error) {
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
	u := bizToUserResponse(user)
	return &v1.GetUserReply{Data: &u}, nil
}

// ListUser 获取用户列表
func (s *BlogAdminUserService) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	opts := biz.UserDOListOption{}
	opts.ListOptions.Pages = 0
	opts.ListOptions.Current = 0
	opts.ListOptions.PageSize = 20
	if in.Current > 0 {
		opts.ListOptions.Pages = in.Pages
		opts.ListOptions.Current = in.Current
		opts.ListOptions.PageSize = in.PageSize
	}
	opts.Username = in.Username
	opts.Nickname = in.Nickname
	opts.Email = in.Email
	opts.StatusFlag = int(in.StatusFlag)
	opts.ListOptions.Init()
	userList, err := s.uc.ListAll(ctx, opts)
	if err != nil {
		return nil, err
	}
	pageInfo := &v1.PageInfo{
		Current:   userList.Current,
		PageSize:  userList.PageSize,
		Total:     userList.TotalCount,
		Pages:     userList.Pages,
		FirstFlag: userList.FirstFlag,
		LastFlag:  userList.LastFlag,
	}
	users := make([]*v1.UserInfoResponse, 0, len(userList.Items))
	for _, user := range userList.Items {
		temp := bizToUserResponse(user)
		users = append(users, &temp)
	}
	return &v1.ListUserReply{PageInfo: pageInfo, Items: users}, nil
}
func (s *BlogAdminUserService) ChangePassword(ctx context.Context, in *v1.ChangePasswordRequest) (*v1.CreateUserReply, error) {
	userDO := &biz.UserDO{
		Password: in.Password,
		Salt:     s.conf.Token,
	}
	userDO.ID = in.Id
	err := s.uc.ChangePassword(ctx, userDO)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{}, nil
}
func bizToUserResponse(user *biz.UserDO) v1.UserInfoResponse {
	userInfoRsp := v1.UserInfoResponse{
		Id:         user.ID,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Avatar:     user.Avatar,
		Email:      user.Email,
		StatusFlag: int32(user.StatusFlag),
	}
	alen := 0
	if user.Roles != nil && len(user.Roles) >= 0 {
		alen = len(user.Roles)
	}
	roles := make([]*v1.URoleInfo, 0, alen)
	if user.Roles != nil {
		for _, obj := range user.Roles {
			roles = append(roles, &v1.URoleInfo{
				Id:         obj.ID,
				Name:       obj.Name,
				Identifier: obj.Identifier,
			})
		}
		userInfoRsp.Roles = roles
	}

	return userInfoRsp
}
