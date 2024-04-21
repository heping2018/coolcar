package main

import (
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/auth/wchart"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("cannot", err)
	}
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("cannot", err)
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
		OpenIDResolver: &wchart.Service{
			AppId:     "wx0a1604bef842ed71",
			AppSecret: "5795fffd8d2c5522e630e0f405ea0718",
		},
	})
	err = s.Serve(lis)
	logger.Fatal("coa", zap.Error(err))
}
