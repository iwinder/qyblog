// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_menus_agent.proto

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

const OperationQyAdminMenusAgentCreateQyAdminMenusAgent = "/api.qycms_bff.admin.v1.QyAdminMenusAgent/CreateQyAdminMenusAgent"
const OperationQyAdminMenusAgentDeleteQyAdminMenusAgents = "/api.qycms_bff.admin.v1.QyAdminMenusAgent/DeleteQyAdminMenusAgents"
const OperationQyAdminMenusAgentListQyAdminMenusAgent = "/api.qycms_bff.admin.v1.QyAdminMenusAgent/ListQyAdminMenusAgent"
const OperationQyAdminMenusAgentUpdateQyAdminMenusAgent = "/api.qycms_bff.admin.v1.QyAdminMenusAgent/UpdateQyAdminMenusAgent"

type QyAdminMenusAgentHTTPServer interface {
	CreateQyAdminMenusAgent(context.Context, *CreateQyAdminMenusAgentRequest) (*CreateQyAdminMenusAgentReply, error)
	DeleteQyAdminMenusAgents(context.Context, *DeleteQyAdminMenusAgentRequest) (*DeleteQyAdminMenusAgentReply, error)
	ListQyAdminMenusAgent(context.Context, *ListQyAdminMenusAgentRequest) (*ListQyAdminMenusAgentReply, error)
	UpdateQyAdminMenusAgent(context.Context, *UpdateQyAdminMenusAgentRequest) (*UpdateQyAdminMenusAgentReply, error)
}

func RegisterQyAdminMenusAgentHTTPServer(s *http.Server, srv QyAdminMenusAgentHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/v1/menusAgent", _QyAdminMenusAgent_CreateQyAdminMenusAgent0_HTTP_Handler(srv))
	r.PUT("/api/admin/v1/menusAgent/{id}", _QyAdminMenusAgent_UpdateQyAdminMenusAgent0_HTTP_Handler(srv))
	r.DELETE("/api/admin/v1/menusAgent", _QyAdminMenusAgent_DeleteQyAdminMenusAgents0_HTTP_Handler(srv))
	r.GET("/api/admin/v1/menusAgent", _QyAdminMenusAgent_ListQyAdminMenusAgent0_HTTP_Handler(srv))
}

func _QyAdminMenusAgent_CreateQyAdminMenusAgent0_HTTP_Handler(srv QyAdminMenusAgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateQyAdminMenusAgentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAgentCreateQyAdminMenusAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateQyAdminMenusAgent(ctx, req.(*CreateQyAdminMenusAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateQyAdminMenusAgentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAgent_UpdateQyAdminMenusAgent0_HTTP_Handler(srv QyAdminMenusAgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateQyAdminMenusAgentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAgentUpdateQyAdminMenusAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateQyAdminMenusAgent(ctx, req.(*UpdateQyAdminMenusAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateQyAdminMenusAgentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAgent_DeleteQyAdminMenusAgents0_HTTP_Handler(srv QyAdminMenusAgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteQyAdminMenusAgentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAgentDeleteQyAdminMenusAgents)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteQyAdminMenusAgents(ctx, req.(*DeleteQyAdminMenusAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteQyAdminMenusAgentReply)
		return ctx.Result(200, reply)
	}
}

func _QyAdminMenusAgent_ListQyAdminMenusAgent0_HTTP_Handler(srv QyAdminMenusAgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListQyAdminMenusAgentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQyAdminMenusAgentListQyAdminMenusAgent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListQyAdminMenusAgent(ctx, req.(*ListQyAdminMenusAgentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListQyAdminMenusAgentReply)
		return ctx.Result(200, reply)
	}
}

type QyAdminMenusAgentHTTPClient interface {
	CreateQyAdminMenusAgent(ctx context.Context, req *CreateQyAdminMenusAgentRequest, opts ...http.CallOption) (rsp *CreateQyAdminMenusAgentReply, err error)
	DeleteQyAdminMenusAgents(ctx context.Context, req *DeleteQyAdminMenusAgentRequest, opts ...http.CallOption) (rsp *DeleteQyAdminMenusAgentReply, err error)
	ListQyAdminMenusAgent(ctx context.Context, req *ListQyAdminMenusAgentRequest, opts ...http.CallOption) (rsp *ListQyAdminMenusAgentReply, err error)
	UpdateQyAdminMenusAgent(ctx context.Context, req *UpdateQyAdminMenusAgentRequest, opts ...http.CallOption) (rsp *UpdateQyAdminMenusAgentReply, err error)
}

type QyAdminMenusAgentHTTPClientImpl struct {
	cc *http.Client
}

func NewQyAdminMenusAgentHTTPClient(client *http.Client) QyAdminMenusAgentHTTPClient {
	return &QyAdminMenusAgentHTTPClientImpl{client}
}

func (c *QyAdminMenusAgentHTTPClientImpl) CreateQyAdminMenusAgent(ctx context.Context, in *CreateQyAdminMenusAgentRequest, opts ...http.CallOption) (*CreateQyAdminMenusAgentReply, error) {
	var out CreateQyAdminMenusAgentReply
	pattern := "/api/admin/v1/menusAgent"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAgentCreateQyAdminMenusAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAgentHTTPClientImpl) DeleteQyAdminMenusAgents(ctx context.Context, in *DeleteQyAdminMenusAgentRequest, opts ...http.CallOption) (*DeleteQyAdminMenusAgentReply, error) {
	var out DeleteQyAdminMenusAgentReply
	pattern := "/api/admin/v1/menusAgent"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAgentDeleteQyAdminMenusAgents))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAgentHTTPClientImpl) ListQyAdminMenusAgent(ctx context.Context, in *ListQyAdminMenusAgentRequest, opts ...http.CallOption) (*ListQyAdminMenusAgentReply, error) {
	var out ListQyAdminMenusAgentReply
	pattern := "/api/admin/v1/menusAgent"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQyAdminMenusAgentListQyAdminMenusAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QyAdminMenusAgentHTTPClientImpl) UpdateQyAdminMenusAgent(ctx context.Context, in *UpdateQyAdminMenusAgentRequest, opts ...http.CallOption) (*UpdateQyAdminMenusAgentReply, error) {
	var out UpdateQyAdminMenusAgentReply
	pattern := "/api/admin/v1/menusAgent/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQyAdminMenusAgentUpdateQyAdminMenusAgent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
