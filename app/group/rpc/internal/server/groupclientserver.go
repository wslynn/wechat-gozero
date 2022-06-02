// Code generated by goctl. DO NOT EDIT!
// Source: group.proto

package server

import (
	"context"

	"wechat-gozero/app/group/rpc/internal/logic"
	"wechat-gozero/app/group/rpc/internal/svc"
	"wechat-gozero/app/group/rpc/proto"
)

type GroupClientServer struct {
	svcCtx *svc.ServiceContext
	proto.UnimplementedGroupClientServer
}

func NewGroupClientServer(svcCtx *svc.ServiceContext) *GroupClientServer {
	return &GroupClientServer{
		svcCtx: svcCtx,
	}
}

func (s *GroupClientServer) AddFriend(ctx context.Context, in *proto.AddFriendRequest) (*proto.AddFriendResponse, error) {
	l := logic.NewAddFriendLogic(ctx, s.svcCtx)
	return l.AddFriend(in)
}

func (s *GroupClientServer) HandleFriend(ctx context.Context, in *proto.HandleFriendRequest) (*proto.HandleFriendResponse, error) {
	l := logic.NewHandleFriendLogic(ctx, s.svcCtx)
	return l.HandleFriend(in)
}

func (s *GroupClientServer) GroupUserList(ctx context.Context, in *proto.GroupUserListRequest) (*proto.GroupUserListResponse, error) {
	l := logic.NewGroupUserListLogic(ctx, s.svcCtx)
	return l.GroupUserList(in)
}

func (s *GroupClientServer) MessageGroupInfoList(ctx context.Context, in *proto.MessageGroupInfoListRequest) (*proto.MessageGroupInfoListResponse, error) {
	l := logic.NewMessageGroupInfoListLogic(ctx, s.svcCtx)
	return l.MessageGroupInfoList(in)
}
