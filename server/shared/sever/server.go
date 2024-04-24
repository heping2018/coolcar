package sever

import (
	"coolcar/shared/auth"
	"log"
	"net"

	"github.com/toolkits/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Logger            *zap.Logger
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Fatal("cannot connection mongodb", nameField, err)
	}
	var opts []grpc.ServerOption
	// type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			logger.Fatal("cannot create interceptor", nameField, err)
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	// 注册拦截器
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	err = s.Serve(lis)
	logger.Fatal("coa", zap.Error(err))
	return err
}
