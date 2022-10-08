// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_menus_admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationQyAdminMenusAdminCreateQyAdminMenusAdmin = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/CreateQyAdminMenusAdmin"
const OperationQyAdminMenusAdminDeleteQyAdminMenusAdmin = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmin"
const OperationQyAdminMenusAdminDeleteQyAdminMenusAdmins = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmins"
const OperationQyAdminMenusAdminGetMyMenusAdminInfo = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetMyMenusAdminInfo"
const OperationQyAdminMenusAdminGetQyAdminMenusAdmin = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetQyAdminMenusAdmin"
const OperationQyAdminMenusAdminListQyAdminMenusAdmin = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/ListQyAdminMenusAdmin"
const OperationQyAdminMenusAdminUpdateQyAdminMenusAdmin = "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/UpdateQyAdminMenusAdmin"

type QyAdminMenusAdminHTTPServer interface {
	CreateQyAdminMenusAdmin(context.Context, *CreateQyAdminMenusAdminRequest) (*CreateQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmin(context.Context, *DeleteQyAdminMenusAdminRequest) (*DeleteQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmins(context.Context, *DeleteQyAdminMenusAdminsRequest) (*DeleteQyAdminMenusAdminsReply, error)
	GetMyMenusAdminInfo(context.Context, *GetMyMenusAdminInfoReq) (*GetMyMenusAdminInfoReply, error)
	GetQyAdminMenusAdmin(context.Context, *GetQyAdminMenusAdminRequest) (*GetQyAdminMenusAdminReply, error)
	ListQyAdminMenusAdmin(context.Context, *ListQyAdminMenusAdminRequest) (*ListQyAdminMenusAdminReply, error)
	UpdateQyAdminMenusAdmin(context.Context, *UpdateQyAdminMenusAdminRequest) (*UpdateQyAdminMenusAdminReply, error)
}

func RegisterQyAdminMenusAdminHTTPServer(s *http.Server, srv QyAdminMenusAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/menusAdmin", _QyAdminMenusAdmin_CreateQyAdminMenusAdmin0_HTTP_Handler(srv))
	r.PUT("/api/admin/v1/menusAdmin/{id}", _QyAdminMenusAdmin_UpdateQyAdminMenusAdmin0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/menusAdmin/{id}", _QyAdminMenusAdmin_DeleteQyAdminMenusAdmin0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/menusAdmin", _QyAdminMenusAdmin_DeleteQyAdminMenusAdmins0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/menusAdmin/{id}", _QyAdminMenusAdmin_GetQyAdminMenusAdmin0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/menusAdmin-my", _QyAdminMenusAdmin_GetMyMenusAdminInfo0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/menusAdmin", _QyAdminMenusAdmin_ListQyAdminMenusAdmin0_HTTP_Handler(srv))
}

func _QyAdminMenusAdmin_CreateQyAdminMenusAdmin0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminMenusAdminRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminCreateQyAdminMenusAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminMenusAdmin(ctx, req.(*CreateQyAdminMenusAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminMenusAdminReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_UpdateQyAdminMenusAdmin0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminMenusAdminRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminUpdateQyAdminMenusAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminMenusAdmin(ctx, req.(*UpdateQyAdminMenusAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminMenusAdminReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_DeleteQyAdminMenusAdmin0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminMenusAdminRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminDeleteQyAdminMenusAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminMenusAdmin(ctx, req.(*DeleteQyAdminMenusAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminMenusAdminReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_DeleteQyAdminMenusAdmins0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminMenusAdminsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminDeleteQyAdminMenusAdmins)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminMenusAdmins(ctx, req.(*DeleteQyAdminMenusAdminsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminMenusAdminsReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_GetQyAdminMenusAdmin0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetQyAdminMenusAdminRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminGetQyAdminMenusAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetQyAdminMenusAdmin(ctx, req.(*GetQyAdminMenusAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetQyAdminMenusAdminReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_GetMyMenusAdminInfo0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyMenusAdminInfoReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminGetMyMenusAdminInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyMenusAdminInfo(ctx, req.(*GetMyMenusAdminInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyMenusAdminInfoReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAdmin_ListQyAdminMenusAdmin0_HTTP_Handler(srv QyAdminMenusAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminMenusAdminRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAdminListQyAdminMenusAdmin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminMenusAdmin(ctx, req.(*ListQyAdminMenusAdminRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminMenusAdminReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminMenusAdminHTTPClient interface {
	CreateQyAdminMenusAdmin(ctx context.Context, req *CreateQyAdminMenusAdminRequest, opts ...http.CallOption) (rsp *CreateQyAdminMenusAdminReply, err error)
	DeleteQyAdminMenusAdmin(ctx context.Context, req *DeleteQyAdminMenusAdminRequest, opts ...http.CallOption) (rsp *DeleteQyAdminMenusAdminReply, err error)
	DeleteQyAdminMenusAdmins(ctx context.Context, req *DeleteQyAdminMenusAdminsRequest, opts ...http.CallOption) (rsp *DeleteQyAdminMenusAdminsReply, err error)
	GetMyMenusAdminInfo(ctx context.Context, req *GetMyMenusAdminInfoReq, opts ...http.CallOption) (rsp *GetMyMenusAdminInfoReply, err error)
	GetQyAdminMenusAdmin(ctx context.Context, req *GetQyAdminMenusAdminRequest, opts ...http.CallOption) (rsp *GetQyAdminMenusAdminReply, err error)
	ListQyAdminMenusAdmin(ctx context.Context, req *ListQyAdminMenusAdminRequest, opts ...http.CallOption) (rsp *ListQyAdminMenusAdminReply, err error)
	UpdateQyAdminMenusAdmin(ctx context.Context, req *UpdateQyAdminMenusAdminRequest, opts ...http.CallOption) (rsp *UpdateQyAdminMenusAdminReply, err error)
}

type QyAdminMenusAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminMenusAdminHTTPClient(client *http.Client) QyAdminMenusAdminHTTPClient {
	return &QyAdminMenusAdminHTTPClientImpl{client}
}

func (c *QyAdminMenusAdminHTTPClientImpl) CreateQyAdminMenusAdmin(ctx context.Context, in *CreateQyAdminMenusAdminRequest, opts ...http.CallOption) (*CreateQyAdminMenusAdminReply, error) {
	var out CreateQyAdminMenusAdminReply
	pattern := "/api/admin/v1/menusAdmin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminCreateQyAdminMenusAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) DeleteQyAdminMenusAdmin(ctx context.Context, in *DeleteQyAdminMenusAdminRequest, opts ...http.CallOption) (*DeleteQyAdminMenusAdminReply, error) {
	var out DeleteQyAdminMenusAdminReply
	pattern := "/api/admin/v1/menusAdmin/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminDeleteQyAdminMenusAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) DeleteQyAdminMenusAdmins(ctx context.Context, in *DeleteQyAdminMenusAdminsRequest, opts ...http.CallOption) (*DeleteQyAdminMenusAdminsReply, error) {
	var out DeleteQyAdminMenusAdminsReply
	pattern := "/api/admin/v1/menusAdmin"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminDeleteQyAdminMenusAdmins))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) GetMyMenusAdminInfo(ctx context.Context, in *GetMyMenusAdminInfoReq, opts ...http.CallOption) (*GetMyMenusAdminInfoReply, error) {
	var out GetMyMenusAdminInfoReply
	pattern := "/api/admin/v1/menusAdmin-my"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminGetMyMenusAdminInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) GetQyAdminMenusAdmin(ctx context.Context, in *GetQyAdminMenusAdminRequest, opts ...http.CallOption) (*GetQyAdminMenusAdminReply, error) {
	var out GetQyAdminMenusAdminReply
	pattern := "/api/admin/v1/menusAdmin/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminGetQyAdminMenusAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) ListQyAdminMenusAdmin(ctx context.Context, in *ListQyAdminMenusAdminRequest, opts ...http.CallOption) (*ListQyAdminMenusAdminReply, error) {
	var out ListQyAdminMenusAdminReply
	pattern := "/api/admin/v1/menusAdmin"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminListQyAdminMenusAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAdminHTTPClientImpl) UpdateQyAdminMenusAdmin(ctx context.Context, in *UpdateQyAdminMenusAdminRequest, opts ...http.CallOption) (*UpdateQyAdminMenusAdminReply, error) {
	var out UpdateQyAdminMenusAdminReply
	pattern := "/api/admin/v1/menusAdmin/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAdminUpdateQyAdminMenusAdmin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
