// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/wslynn/wechat-gozero/app/msg/rpc/proto"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/logic"
	"github.com/wslynn/wechat-gozero/app/user/rpc/internal/svc"
)

type UserClientServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserClientServer
}

func NewUserClientServer(svcCtx *svc.ServiceContext) *UserClientServer {
	return &UserClientServer{
		svcCtx: svcCtx,
	}
}

func (s *UserClientServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserClientServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserClientServer) PersonalInfo(ctx context.Context, in *user.PersonalInfoRequest) (*user.PersonalInfoResponse, error) {
	l := logic.NewPersonalInfoLogic(ctx, s.svcCtx)
	return l.PersonalInfo(in)
}
