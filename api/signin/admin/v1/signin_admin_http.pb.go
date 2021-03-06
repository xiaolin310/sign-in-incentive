// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

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

type SignInAdminServiceHTTPServer interface {
	GetSignInRecord(context.Context, *GetSignInRecordRequest) (*GetSignInRecordReply, error)
}

func RegisterSignInAdminServiceHTTPServer(s *http.Server, srv SignInAdminServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/admin/signin/v1/getSignInRecord", _SignInAdminService_GetSignInRecord0_HTTP_Handler(srv))
}

func _SignInAdminService_GetSignInRecord0_HTTP_Handler(srv SignInAdminServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSignInRecordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.admin.v1.SignInAdminService/GetSignInRecord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSignInRecord(ctx, req.(*GetSignInRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSignInRecordReply)
		return ctx.Result(200, reply)
	}
}

type SignInAdminServiceHTTPClient interface {
	GetSignInRecord(ctx context.Context, req *GetSignInRecordRequest, opts ...http.CallOption) (rsp *GetSignInRecordReply, err error)
}

type SignInAdminServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSignInAdminServiceHTTPClient(client *http.Client) SignInAdminServiceHTTPClient {
	return &SignInAdminServiceHTTPClientImpl{client}
}

func (c *SignInAdminServiceHTTPClientImpl) GetSignInRecord(ctx context.Context, in *GetSignInRecordRequest, opts ...http.CallOption) (*GetSignInRecordReply, error) {
	var out GetSignInRecordReply
	pattern := "/api/admin/signin/v1/getSignInRecord"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.signin.admin.v1.SignInAdminService/GetSignInRecord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
