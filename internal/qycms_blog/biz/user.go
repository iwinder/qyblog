package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	jwt2 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/auth/jwt"
	metaV1 "github.com/iwinder/qingyucms/internal/pkg/qycms_common/meta/v1"
	"github.com/iwinder/qingyucms/internal/pkg/qycms_common/utils/bcryptUtil"
	"github.com/iwinder/qingyucms/internal/qycms_blog/conf"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UserDO is a UserDO model.
type UserDO struct {
	metaV1.ObjectMeta
	Username  string
	Nickname  string
	Avatar    string
	Password  string
	Salt      string
	Email     string
	Phone     string
	AdminFlag bool
	Roles     []*RoleDO
}
type UserInfoDO struct {
	metaV1.ObjectMeta
	Nickname  string
	Avatar    string
	Email     string
	Phone     string
	Token     string
	RoleNames string
}
type UserDOList struct {
	metaV1.ListMeta `json:",inline"`
	Items           []*UserDO `json:"items"`
}

type UserDOListOption struct {
	metaV1.ListOptions `json:"page"`
	UserDO             `json:"item"`
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *UserDO) (*UserDO, error)
	Update(context.Context, *UserDO) (*UserDO, error)
	Delete(context.Context, uint64) error
	DeleteList(c context.Context, uids []uint64) error
	FindByID(context.Context, uint64) (*UserDO, error)
	FindByUsername(c context.Context, username string) (*UserDO, error)
	ListAll(c context.Context, opts UserDOListOption) (*UserDOList, error)
	ChangePassword(c context.Context, user *UserDO) error
	//VerifyPassword(ctx context.Context, u *UserDO) (bool, error)
}

// UserUsecase is a UserDO usecase.
type UserUsecase struct {
	repo     UserRepo
	role     *RoleUsecase
	userRole *UserRoleUsecase
	log      *log.Helper
}

// NewUserUsecase new a UserDO usecase.
func NewUserUsecase(repo UserRepo, role *RoleUsecase, userRole *UserRoleUsecase, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, role: role, userRole: userRole, log: log.NewHelper(logger)}
}

// CreateUser creates a UserDO, and returns the new UserDO.
func (uc *UserUsecase) CreateUser(ctx context.Context, user *UserDO) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", user.Username)
	userDO, err := uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}
	// 关联
	if user.Roles != nil && len(user.Roles) > 0 {
		// 权限
		rerr := uc.userRole.UpdateRoleForUser(ctx, user)
		if rerr != nil {
			return nil, rerr
		}
	}
	return userDO, nil
}

// Update 更新用户
func (uc *UserUsecase) Update(ctx context.Context, user *UserDO) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", user.Username)
	userDO, err := uc.repo.Update(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 权限
	rerr := uc.userRole.UpdateRoleForUser(ctx, user)
	if rerr != nil {
		return nil, rerr
	}
	return userDO, nil
}

// Delete 根据ID删除用户
func (uc *UserUsecase) Delete(ctx context.Context, id uint64) error {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

// DeleteList 根据ID批量删除用户
func (uc *UserUsecase) DeleteList(ctx context.Context, ids []uint64) error {
	uc.log.WithContext(ctx).Infof("DeleteList: %v", ids)
	return uc.repo.DeleteList(ctx, ids)
}

// FindOneByID 根据ID查询用户信息
func (uc *UserUsecase) FindOneByID(ctx context.Context, id uint64) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByID: %v", id)
	user, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	// 获取角色
	roles, ree := uc.role.FindByUserId(ctx, id)
	user.Roles = roles
	if ree != nil {
		user.Roles = make([]*RoleDO, 0)
	}
	return user, nil
}

// FindOneByUsername 根据用户名查询用户信息
func (uc *UserUsecase) FindOneByUsername(ctx context.Context, username string) (*UserDO, error) {
	uc.log.WithContext(ctx).Infof("FindOneByUsername: %v", username)
	user, err := uc.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// 获取角色
	roles, ree := uc.role.FindByUserId(ctx, user.ID)
	user.Roles = roles
	if ree != nil {
		user.Roles = make([]*RoleDO, 0)
	}
	return user, nil
}

// ListAll 批量查询
func (uc *UserUsecase) ListAll(ctx context.Context, opts UserDOListOption) (*UserDOList, error) {
	uc.log.WithContext(ctx).Infof("ListAll")
	userPOs, err := uc.repo.ListAll(ctx, opts)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	for _, user := range userPOs.Items {
		roles, ree := uc.role.FindByUserId(ctx, user.ID)
		user.Roles = roles
		if ree != nil {
			user.Roles = make([]*RoleDO, 0)
		}
	}
	return userPOs, nil
}

func (uc *UserUsecase) VerifyPassword(ctx context.Context, u *UserDO, authConf *conf.Auth) (*UserInfoDO, error) {
	userInfo, err := uc.FindOneByUsername(ctx, u.Username)
	if err != nil {
		uc.log.WithContext(ctx).Error(fmt.Errorf("登录失败-查询用户:%v", err))
		return nil, errors.New("登录失败:账号或密码错误")
	}
	aerr := bcryptUtil.Compare(userInfo.Password, u.Password+u.Salt)
	if aerr != nil {
		uc.log.WithContext(ctx).Error(fmt.Errorf("登录失败-密码比较失败:%v", aerr))
		return nil, errors.New("登录失败:账号或密码错误")
	}
	var roleNames []string
	var roleIds []string
	if len(userInfo.Roles) > 0 {
		roleNames = make([]string, 0, len(userInfo.Roles))
		for _, obj := range userInfo.Roles {
			roleIds = append(roleIds, strconv.FormatUint(obj.ID, 10))
			roleNames = append(roleNames, obj.Name)
		}
	} else {
		roleIds = make([]string, 0, 0)
		roleNames = make([]string, 0, 0)
	}
	roleNameStr := strings.Join(roleNames, ",")
	roleIdStr := strings.Join(roleIds, ",")
	// 获取角色信息
	claims := jwt2.SecurityUser{
		ID:            userInfo.ID,
		NickName:      userInfo.Nickname,
		AuthorityName: userInfo.Username,
		RoleIds:       roleIdStr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(authConf.Jwt.ExpireDuration.AsDuration())), // 设置token的过期时间
		},
	}
	// 生成token
	token, jerr := jwt2.CreateToken(claims, authConf.Jwt.JwtSecret)
	if jerr != nil {
		uc.log.WithContext(ctx).Error("登录失败，生成token失败：%v", jerr)
		return nil, errors.New("登录失败:账号或密码错误")
	}

	user := &UserInfoDO{
		Nickname:  userInfo.Nickname,
		Avatar:    userInfo.Avatar,
		Token:     token,
		RoleNames: roleNameStr,
	}

	return user, jerr
}

func (uc *UserUsecase) ChangePassword(ctx context.Context, user *UserDO) error {
	var err error
	user.Password, err = bcryptUtil.Encrypt(user.Password + user.Salt)
	if err != nil {
		return err
	}
	return uc.repo.ChangePassword(ctx, user)
}
