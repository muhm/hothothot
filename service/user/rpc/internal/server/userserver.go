// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"hothothot/service/user/rpc/internal/logic"
	"hothothot/service/user/rpc/internal/svc"
	"hothothot/service/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUserById(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UserServer) GetUserByName(ctx context.Context, in *user.NameReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserByNameLogic(ctx, s.svcCtx)
	return l.GetUserByName(in)
}

func (s *UserServer) GetUserByMail(ctx context.Context, in *user.MailReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserByMailLogic(ctx, s.svcCtx)
	return l.GetUserByMail(in)
}

func (s *UserServer) SaveOtpSecret(ctx context.Context, in *user.OtpReq) (*user.OtpReply, error) {
	l := logic.NewSaveOtpSecretLogic(ctx, s.svcCtx)
	return l.SaveOtpSecret(in)
}
