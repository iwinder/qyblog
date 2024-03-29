// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_link.proto

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

const OperationQyAdminLinkCreateQyAdminLink = "/api.qycms_bff.admin.v1.QyAdminLink/CreateQyAdminLink"
const OperationQyAdminLinkDeleteQyAdminLinks = "/api.qycms_bff.admin.v1.QyAdminLink/DeleteQyAdminLinks"
const OperationQyAdminLinkListQyAdminLink = "/api.qycms_bff.admin.v1.QyAdminLink/ListQyAdminLink"
const OperationQyAdminLinkUpdateQyAdminLink = "/api.qycms_bff.admin.v1.QyAdminLink/UpdateQyAdminLink"

type QyAdminLinkHTTPServer interface {
	CreateQyAdminLink(context.Context, *CreateQyAdminLinkRequest) (*CreateQyAdminLinkReply, error)
	DeleteQyAdminLinks(context.Context, *DeleteQyAdminLinksRequest) (*DeleteQyAdminLinkReply, error)
	ListQyAdminLink(context.Context, *ListQyAdminLinkRequest) (*ListQyAdminLinkReply, error)
	UpdateQyAdminLink(context.Context, *UpdateQyAdminLinkRequest) (*UpdateQyAdminLinkReply, error)
}

func RegisterQyAdminLinkHTTPServer(s *http.Server, srv QyAdminLinkHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/link", _QyAdminLink_CreateQyAdminLink0_HTTP_Handler(srv))
	r.PUT("/api/admin/v1/link/{id}", _QyAdminLink_UpdateQyAdminLink0_HTTP_Handler(srv))
	r.POST("/api/admin/v1/link", _QyAdminLink_DeleteQyAdminLinks0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/link", _QyAdminLink_ListQyAdminLink0_HTTP_Handler(srv))
}

func _QyAdminLink_CreateQyAdminLink0_HTTP_Handler(srv QyAdminLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminLinkCreateQyAdminLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminLink(ctx, req.(*CreateQyAdminLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminLinkReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminLink_UpdateQyAdminLink0_HTTP_Handler(srv QyAdminLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminLinkUpdateQyAdminLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminLink(ctx, req.(*UpdateQyAdminLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminLinkReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminLink_DeleteQyAdminLinks0_HTTP_Handler(srv QyAdminLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminLinksRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminLinkDeleteQyAdminLinks)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminLinks(ctx, req.(*DeleteQyAdminLinksRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminLinkReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminLink_ListQyAdminLink0_HTTP_Handler(srv QyAdminLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminLinkRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminLinkListQyAdminLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminLink(ctx, req.(*ListQyAdminLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminLinkReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminLinkHTTPClient interface {
	CreateQyAdminLink(ctx context.Context, req *CreateQyAdminLinkRequest, opts ...http.CallOption) (rsp *CreateQyAdminLinkReply, err error)
	DeleteQyAdminLinks(ctx context.Context, req *DeleteQyAdminLinksRequest, opts ...http.CallOption) (rsp *DeleteQyAdminLinkReply, err error)
	ListQyAdminLink(ctx context.Context, req *ListQyAdminLinkRequest, opts ...http.CallOption) (rsp *ListQyAdminLinkReply, err error)
	UpdateQyAdminLink(ctx context.Context, req *UpdateQyAdminLinkRequest, opts ...http.CallOption) (rsp *UpdateQyAdminLinkReply, err error)
}

type QyAdminLinkHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminLinkHTTPClient(client *http.Client) QyAdminLinkHTTPClient {
	return &QyAdminLinkHTTPClientImpl{client}
}

func (c *QyAdminLinkHTTPClientImpl) CreateQyAdminLink(ctx context.Context, in *CreateQyAdminLinkRequest, opts ...http.CallOption) (*CreateQyAdminLinkReply, error) {
	var out CreateQyAdminLinkReply
	pattern := "/api/admin/v1/link"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminLinkCreateQyAdminLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminLinkHTTPClientImpl) DeleteQyAdminLinks(ctx context.Context, in *DeleteQyAdminLinksRequest, opts ...http.CallOption) (*DeleteQyAdminLinkReply, error) {
	var out DeleteQyAdminLinkReply
	pattern := "/api/admin/v1/link"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminLinkDeleteQyAdminLinks))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminLinkHTTPClientImpl) ListQyAdminLink(ctx context.Context, in *ListQyAdminLinkRequest, opts ...http.CallOption) (*ListQyAdminLinkReply, error) {
	var out ListQyAdminLinkReply
	pattern := "/api/admin/v1/link"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminLinkListQyAdminLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminLinkHTTPClientImpl) UpdateQyAdminLink(ctx context.Context, in *UpdateQyAdminLinkRequest, opts ...http.CallOption) (*UpdateQyAdminLinkReply, error) {
	var out UpdateQyAdminLinkReply
	pattern := "/api/admin/v1/link/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminLinkUpdateQyAdminLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
