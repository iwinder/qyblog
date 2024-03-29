package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "github.com/iwinder/qyblog/api/qycms_bff/admin/v1"
	"github.com/iwinder/qyblog/internal/qycms_blog/conf"
	"github.com/iwinder/qyblog/internal/qycms_blog/service"
)

// NewGRPCServer new a gRPC server.  gRPC 示例，本项目暂未使用
func NewGRPCServer(c *conf.Server, greeter *service.BlogAdminUserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterQyAdminUserServer(srv, greeter)
	return srv
}
