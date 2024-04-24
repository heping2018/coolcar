package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"log"

	serve "coolcar/shared/sever"

	"google.golang.org/grpc"
)

// 验证token
func main() {
	logger, err := serve.NewZapLogger()
	if err != nil {
		log.Fatal("cannot", err)
	}
	err = serve.RunGRPCServer(&serve.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
	})
	if err != nil {
		log.Fatal("start fail", err)
	}
	// lis, err := net.Listen("tcp", ":8082")
	// if err != nil {
	// 	log.Fatal("cannot connection mongodb", err)
	// }
	// // type UnaryServerInterceptor func(ctx context.Context, req any, info *UnaryServerInfo, handler UnaryHandler) (resp any, err error)
	// in, err := auth.Interceptor("shared/auth/public.key")
	// if err != nil {
	// 	logger.Fatal("cannot create interceptor")
	// }
	// // 注册拦截器
	// s := grpc.NewServer(grpc.UnaryInterceptor(in))
	// rentalpb.RegisterTripServiceServer(s, &trip.Service{
	// 	Logger: logger,
	// })
	// err = s.Serve(lis)
	// logger.Fatal("coa", zap.Error(err))
}
