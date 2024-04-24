package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			UnmarshalOptions: protojson.UnmarshalOptions{},
			MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: true},
		},
	))
	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	}{
		{
			name:         "Auth",
			addr:         ":8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         ":8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}
	for _, s := range serverConfig {
		err := s.registerFunc(c, mux, s.addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatalf("cannot register auth:v=%v", err)
		}
	}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("http to %v", err)
	}
}
