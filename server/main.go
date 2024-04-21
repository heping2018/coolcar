package main

import (
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("fail to %v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}
<<<<<<< HEAD
func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		UnmarshalOptions: protojson.UnmarshalOptions{},
		MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: true},
	}))
	defer cancel()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("mux to %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("http to %v", err)
	}
}
=======
>>>>>>> parent of 2e91dab (step-3)
