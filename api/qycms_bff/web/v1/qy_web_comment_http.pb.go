// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/web/v1/qy_web_comment.proto

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

const OperationQyWebCommentCreateQyWebComment = "/api.qycms_bff.web.v1.QyWebComment/CreateQyWebComment"
const OperationQyWebCommentListQyWebComment = "/api.qycms_bff.web.v1.QyWebComment/ListQyWebComment"

type QyWebCommentHTTPServer interface {
	CreateQyWebComment(context.Context, *CreateQyWebCommentRequest) (*CreateQyWebCommentReply, error)
	ListQyWebComment(context.Context, *ListQyWebCommentRequest) (*ListQyWebCommentReply, error)
}

func RegisterQyWebCommentHTTPServer(s *http.Server, srv QyWebCommentHTTPServer) {
	r := s.Route("/")
	r.POST("/api/web/v1/comment", _QyWebComment_CreateQyWebComment0_HTTP_Handler(srv))
	r.GET("/api/web/v1/comment", _QyWebComment_ListQyWebComment0_HTTP_Handler(srv))
}

func _QyWebComment_CreateQyWebComment0_HTTP_Handler(srv QyWebCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyWebCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyWebCommentCreateQyWebComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyWebComment(ctx, req.(*CreateQyWebCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyWebCommentReply)
		return ctx.Result(200, reply)
	}
}

func _QyWebComment_ListQyWebComment0_HTTP_Handler(srv QyWebCommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyWebCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyWebCommentListQyWebComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyWebComment(ctx, req.(*ListQyWebCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyWebCommentReply)
		return ctx.Result(200, reply)
	}
}

type QyWebCommentHTTPClient interface {
	CreateQyWebComment(ctx context.Context, req *CreateQyWebCommentRequest, opts ...http.CallOption) (rsp *CreateQyWebCommentReply, err error)
	ListQyWebComment(ctx context.Context, req *ListQyWebCommentRequest, opts ...http.CallOption) (rsp *ListQyWebCommentReply, err error)
}

type QyWebCommentHTTPClientImpl struct {
	cc *http.Client
}

func NewQyWebCommentHTTPClient(client *http.Client) QyWebCommentHTTPClient {
	return &QyWebCommentHTTPClientImpl{client}
}

func (c *QyWebCommentHTTPClientImpl) CreateQyWebComment(ctx context.Context, in *CreateQyWebCommentRequest, opts ...http.CallOption) (*CreateQyWebCommentReply, error) {
	var out CreateQyWebCommentReply
	pattern := "/api/web/v1/comment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyWebCommentCreateQyWebComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyWebCommentHTTPClientImpl) ListQyWebComment(ctx context.Context, in *ListQyWebCommentRequest, opts ...http.CallOption) (*ListQyWebCommentReply, error) {
	var out ListQyWebCommentReply
	pattern := "/api/web/v1/comment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyWebCommentListQyWebComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
