#protoc -I=. --go-grpc_out=paths=source_relative:gen/go trip.proto
protoc -I=. --go_out=paths=source_relative:gen/go trip.proto
protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/go trip.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto

PBTS_BIN_DIR=../../wx/miniprogram/node_modules/.bin
PBTS_OUT_DIR=../../wx/miniprogram/service/proto_gen
../../wx/miniprogram/node_modules/.bin/pbjs -t static -w es6 trip.proto  --no-create  --no-encde --no-decode --no-verify --no-delimited --force-number -o ../../wx/miniprogram/service/proto_gen/trip_pb.js
../../wx/miniprogram/node_modules/.bin/pbts -o ../../wx/miniprogram/service/proto_gen/trip_pb.d.ts ../../wx/miniprogram/service/proto_gen/trip_pb.js