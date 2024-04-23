package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/auth/token"
	"coolcar/auth/auth/wchart"
	"coolcar/auth/dao"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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
	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot bytes private key", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
		OpenIDResolver: &wchart.Service{
			AppId:     "wx0a1604bef842ed71",
			AppSecret: "5795fffd8d2c5522e630e0f405ea0718",
		},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJwtTokenGen("coolcar/auth", privatekey),
	})
	err = s.Serve(lis)
	logger.Fatal("coa", zap.Error(err))
}
