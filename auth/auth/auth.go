package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	OpenIDResolver                        OpenIDResolver
	Mongo                                 *dao.Mongo
	TokenGenerator                        TokenGenerator
	TokenExpire                           time.Duration
	Logger                                *zap.Logger
	authpb.UnimplementedAuthServiceServer //pb为了保证向前兼容，必须内嵌 UnimplementedAuthServiceServer 结构体
}

// OpenIDResolver resolves an authorization code
// to an open id.
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// TokenGenerator generates a token for the specified account.
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolver openid:%v", err)
	}
	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	//创建token
	tkn, err := s.TokenGenerator.GenerateToken(accountID.String(), s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
