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

type SingInInterfaceHTTPServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceReply, error)
	GetSignInInfo(context.Context, *GetSignInInfoRequest) (*SignInInfoReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInInfoReply, error)
}

func RegisterSingInInterfaceHTTPServer(s *http.Server, srv SingInInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/signin/v1/register", _SingInInterface_Register0_HTTP_Handler(srv))
	r.POST("/api/signin/v1/login", _SingInInterface_Login0_HTTP_Handler(srv))
	r.GET("/api/signin/v1/signininfo", _SingInInterface_GetSignInInfo0_HTTP_Handler(srv))
	r.GET("/api/signin/v1/signin", _SingInInterface_SignIn0_HTTP_Handler(srv))
	r.GET("/api/signin/v1/getBalance", _SingInInterface_GetBalance0_HTTP_Handler(srv))
}

func _SingInInterface_Register0_HTTP_Handler(srv SingInInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.interface.v1.SingInInterface/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _SingInInterface_Login0_HTTP_Handler(srv SingInInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.interface.v1.SingInInterface/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _SingInInterface_GetSignInInfo0_HTTP_Handler(srv SingInInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSignInInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.interface.v1.SingInInterface/GetSignInInfo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSignInInfo(ctx, req.(*GetSignInInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignInInfoReply)
		return ctx.Result(200, reply)
	}
}

func _SingInInterface_SignIn0_HTTP_Handler(srv SingInInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.interface.v1.SingInInterface/SignIn")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignIn(ctx, req.(*SignInRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignInInfoReply)
		return ctx.Result(200, reply)
	}
}

func _SingInInterface_GetBalance0_HTTP_Handler(srv SingInInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBalanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.signin.interface.v1.SingInInterface/GetBalance")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBalance(ctx, req.(*GetBalanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBalanceReply)
		return ctx.Result(200, reply)
	}
}

type SingInInterfaceHTTPClient interface {
	GetBalance(ctx context.Context, req *GetBalanceRequest, opts ...http.CallOption) (rsp *GetBalanceReply, err error)
	GetSignInInfo(ctx context.Context, req *GetSignInInfoRequest, opts ...http.CallOption) (rsp *SignInInfoReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	SignIn(ctx context.Context, req *SignInRequest, opts ...http.CallOption) (rsp *SignInInfoReply, err error)
}

type SingInInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewSingInInterfaceHTTPClient(client *http.Client) SingInInterfaceHTTPClient {
	return &SingInInterfaceHTTPClientImpl{client}
}

func (c *SingInInterfaceHTTPClientImpl) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...http.CallOption) (*GetBalanceReply, error) {
	var out GetBalanceReply
	pattern := "/api/signin/v1/getBalance"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.signin.interface.v1.SingInInterface/GetBalance"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SingInInterfaceHTTPClientImpl) GetSignInInfo(ctx context.Context, in *GetSignInInfoRequest, opts ...http.CallOption) (*SignInInfoReply, error) {
	var out SignInInfoReply
	pattern := "/api/signin/v1/signininfo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.signin.interface.v1.SingInInterface/GetSignInInfo"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SingInInterfaceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/signin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.signin.interface.v1.SingInInterface/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SingInInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/signin/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.signin.interface.v1.SingInInterface/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SingInInterfaceHTTPClientImpl) SignIn(ctx context.Context, in *SignInRequest, opts ...http.CallOption) (*SignInInfoReply, error) {
	var out SignInInfoReply
	pattern := "/api/signin/v1/signin"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.signin.interface.v1.SingInInterface/SignIn"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
