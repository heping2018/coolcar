package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	opneId, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve open id  %v", err)
	}
	s.Logger.Info("receved codde", zap.String("code", req.Code))
	accountId, err := s.Mongo.ResolveAccountID(c, opneId)
	if err != nil {
		s.Logger.Error("cannot resolve id")
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{AccessToken: "token for open idd" + accountId, ExpiresIn: 7200}, nil
}
