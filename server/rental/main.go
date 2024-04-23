package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/auth"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("cannot", err)
	}
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("cannot connection mongodb", err)
	}
	// type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)
	in, err := auth.Interceptor("shared/auth/public.key")
	if err != nil {
		logger.Fatal("cannot create interceptor")
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})
	err = s.Serve(lis)
	logger.Fatal("coa", zap.Error(err))
}
