package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer"
)

//定义一个token认证的接口
type tokenVerifier interface {
	Verify(token string) (string, error)
}

//定义一个拦截器结构体
type interceptor struct {
	verifier tokenVerifier
}

// HandleReq 实际处理请求的方法
func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	aid := impersonationFromContext(ctx)
	if aid != "" {
		return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
	}
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err = i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token校验失败:%v", err)
	}
	return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
}

//获取上下文中的token
func tokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	//获取请求的元数据的内容
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}
	tkn := ""
	//遍历所有取到的content值
	for _, v := range m[authorizationHeader] {
		//如果字段以Bearer开头
		if strings.HasPrefix(v, bearerPrefix) {
			//截取字符串
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}

type accountIDKey struct {
}

// ContextWithAccountID 构建带有账户信息的上下文
func ContextWithAccountID(ctx context.Context, accountID id.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, accountID)
}

func impersonationFromContext(ctx context.Context) string {
	//获取请求的元数据
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	//获取请求头
	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}

// Interceptor 创建一个grpc 认证拦截器
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("无法打开公钥文件:%v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("读取公钥文件错误:%v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("无法解析公钥:%v", err)
	}
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}
