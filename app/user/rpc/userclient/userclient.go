// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"github.com/wslynn/wechat-gozero/proto/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest         = user.LoginRequest
	LoginResponse        = user.LoginResponse
	PersonalInfoRequest  = user.PersonalInfoRequest
	PersonalInfoResponse = user.PersonalInfoResponse
	RegisterRequest      = user.RegisterRequest
	RegisterResponse     = user.RegisterResponse

	UserClient interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		PersonalInfo(ctx context.Context, in *PersonalInfoRequest, opts ...grpc.CallOption) (*PersonalInfoResponse, error)
	}

	defaultUserClient struct {
		cli zrpc.Client
	}
)

func NewUserClient(cli zrpc.Client) UserClient {
	return &defaultUserClient{
		cli: cli,
	}
}

func (m *defaultUserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserClient) PersonalInfo(ctx context.Context, in *PersonalInfoRequest, opts ...grpc.CallOption) (*PersonalInfoResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.PersonalInfo(ctx, in, opts...)
}
