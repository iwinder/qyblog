// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_comment.proto

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

const OperationQyAdminCommentCreateQyAdminComment = "/api.qycms_bff.admin.v1.QyAdminComment/CreateQyAdminComment"
const OperationQyAdminCommentDeleteQyAdminComment = "/api.qycms_bff.admin.v1.QyAdminComment/DeleteQyAdminComment"
const OperationQyAdminCommentGetQyAdminCommentCount = "/api.qycms_bff.admin.v1.QyAdminComment/GetQyAdminCommentCount"
const OperationQyAdminCommentListQyAdminComment = "/api.qycms_bff.admin.v1.QyAdminComment/ListQyAdminComment"
const OperationQyAdminCommentUpdateQyAdminCommentContent = "/api.qycms_bff.admin.v1.QyAdminComment/UpdateQyAdminCommentContent"
const OperationQyAdminCommentUpdateQyAdminCommentState = "/api.qycms_bff.admin.v1.QyAdminComment/UpdateQyAdminCommentState"

type QyAdminCommentHTTPServer interface {
	CreateQyAdminComment(context.Context, *CreateQyAdminCommentRequest) (*CreateQyAdminCommentReply, error)
	DeleteQyAdminComment(context.Context, *DeleteQyAdminCommentRequest) (*DeleteQyAdminCommentReply, error)
	GetQyAdminCommentCount(context.Context, *GetQyAdminCommentCountRequest) (*GetQyAdminCommentCountReply, error)
	ListQyAdminComment(context.Context, *ListQyAdminCommentRequest) (*ListQyAdminCommentReply, error)
	UpdateQyAdminCommentContent(context.Context, *UpdateQyAdminCommentRequest) (*UpdateQyAdminCommentReply, error)
	UpdateQyAdminCommentState(context.Context, *UpdateQyAdminCommentStateRequest) (*UpdateQyAdminCommentStateReply, error)
}

func RegisterQyAdminCommentHTTPServer(s *http.Server, srv QyAdminCommentHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/comment", _QyAdminComment_CreateQyAdminComment0_HTTP_Handler(srv))
	r.POST("/api/admin/v1/comment/content", _QyAdminComment_UpdateQyAdminCommentContent0_HTTP_Handler(srv))
	r.POST("/api/admin/v1/comment/state", _QyAdminComment_UpdateQyAdminCommentState0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/comment", _QyAdminComment_DeleteQyAdminComment0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/comment", _QyAdminComment_ListQyAdminComment0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/comment/total", _QyAdminComment_GetQyAdminCommentCount0_HTTP_Handler(srv))
}

func _QyAdminComment_CreateQyAdminComment0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentCreateQyAdminComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminComment(ctx, req.(*CreateQyAdminCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminCommentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminComment_UpdateQyAdminCommentContent0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentUpdateQyAdminCommentContent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminCommentContent(ctx, req.(*UpdateQyAdminCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminCommentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminComment_UpdateQyAdminCommentState0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminCommentStateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentUpdateQyAdminCommentState)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminCommentState(ctx, req.(*UpdateQyAdminCommentStateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminCommentStateReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminComment_DeleteQyAdminComment0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentDeleteQyAdminComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminComment(ctx, req.(*DeleteQyAdminCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminCommentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminComment_ListQyAdminComment0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentListQyAdminComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminComment(ctx, req.(*ListQyAdminCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminCommentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminComment_GetQyAdminCommentCount0_HTTP_Handler(srv QyAdminCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetQyAdminCommentCountRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminCommentGetQyAdminCommentCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetQyAdminCommentCount(ctx, req.(*GetQyAdminCommentCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetQyAdminCommentCountReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminCommentHTTPClient interface {
	CreateQyAdminComment(ctx context.Context, req *CreateQyAdminCommentRequest, opts ...http.CallOption) (rsp *CreateQyAdminCommentReply, err error)
	DeleteQyAdminComment(ctx context.Context, req *DeleteQyAdminCommentRequest, opts ...http.CallOption) (rsp *DeleteQyAdminCommentReply, err error)
	GetQyAdminCommentCount(ctx context.Context, req *GetQyAdminCommentCountRequest, opts ...http.CallOption) (rsp *GetQyAdminCommentCountReply, err error)
	ListQyAdminComment(ctx context.Context, req *ListQyAdminCommentRequest, opts ...http.CallOption) (rsp *ListQyAdminCommentReply, err error)
	UpdateQyAdminCommentContent(ctx context.Context, req *UpdateQyAdminCommentRequest, opts ...http.CallOption) (rsp *UpdateQyAdminCommentReply, err error)
	UpdateQyAdminCommentState(ctx context.Context, req *UpdateQyAdminCommentStateRequest, opts ...http.CallOption) (rsp *UpdateQyAdminCommentStateReply, err error)
}

type QyAdminCommentHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminCommentHTTPClient(client *http.Client) QyAdminCommentHTTPClient {
	return &QyAdminCommentHTTPClientImpl{client}
}

func (c *QyAdminCommentHTTPClientImpl) CreateQyAdminComment(ctx context.Context, in *CreateQyAdminCommentRequest, opts ...http.CallOption) (*CreateQyAdminCommentReply, error) {
	var out CreateQyAdminCommentReply
	pattern := "/api/admin/v1/comment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCommentCreateQyAdminComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCommentHTTPClientImpl) DeleteQyAdminComment(ctx context.Context, in *DeleteQyAdminCommentRequest, opts ...http.CallOption) (*DeleteQyAdminCommentReply, error) {
	var out DeleteQyAdminCommentReply
	pattern := "/api/admin/v1/comment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCommentDeleteQyAdminComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCommentHTTPClientImpl) GetQyAdminCommentCount(ctx context.Context, in *GetQyAdminCommentCountRequest, opts ...http.CallOption) (*GetQyAdminCommentCountReply, error) {
	var out GetQyAdminCommentCountReply
	pattern := "/api/admin/v1/comment/total"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminCommentGetQyAdminCommentCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCommentHTTPClientImpl) ListQyAdminComment(ctx context.Context, in *ListQyAdminCommentRequest, opts ...http.CallOption) (*ListQyAdminCommentReply, error) {
	var out ListQyAdminCommentReply
	pattern := "/api/admin/v1/comment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminCommentListQyAdminComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCommentHTTPClientImpl) UpdateQyAdminCommentContent(ctx context.Context, in *UpdateQyAdminCommentRequest, opts ...http.CallOption) (*UpdateQyAdminCommentReply, error) {
	var out UpdateQyAdminCommentReply
	pattern := "/api/admin/v1/comment/content"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCommentUpdateQyAdminCommentContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminCommentHTTPClientImpl) UpdateQyAdminCommentState(ctx context.Context, in *UpdateQyAdminCommentStateRequest, opts ...http.CallOption) (*UpdateQyAdminCommentStateReply, error) {
	var out UpdateQyAdminCommentStateReply
	pattern := "/api/admin/v1/comment/state"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminCommentUpdateQyAdminCommentState))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
