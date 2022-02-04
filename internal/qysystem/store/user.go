package store

import (
	"context"
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-api/qysystem/v1"
	metav1 "gitee.com/windcoder/qingyucms/internal/pkg/qy-common/meta/v1"
)

type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
}
