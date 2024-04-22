package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/auth/wchart"
	"coolcar/auth/dao"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("cannot", err)
	}
	lis, err := net.Listen("tcp", ":8081")
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("cannot connection mongodb", err)
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
		OpenIDResolver: &wchart.Service{
			AppId:     "wx0a1604bef842ed71",
			AppSecret: "5795fffd8d2c5522e630e0f405ea0718",
		},
		Mongo: dao.NewMongo(mc.Database("coolcar")),
	})
	err = s.Serve(lis)
	logger.Fatal("coa", zap.Error(err))
}
