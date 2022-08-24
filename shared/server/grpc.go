package server

import (
	"coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

// RunGRPCServer 启动一个GRPC服务
func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}
	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create auth Interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	//创建一个Server
	s := grpc.NewServer(opts...)
	//设置注册方法
	c.RegisterFunc(s)
	//注册心跳检测
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
