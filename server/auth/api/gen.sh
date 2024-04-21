# 生成rpc 服务端 和 客户端
protoc -I=D:\GitHub\project\coolcar\server\auth\api --go-grpc_out=paths=source_relative:gen/v1 auth.proto

# 对proto中数据结构生成go 代码
protoc -I=D:\GitHub\project\coolcar\server\auth\api --go_out=paths=source_relative:gen/v1 auth.proto
#protoc -I=D:\GitHub\project\coolcar\server\auth\api --go_out=plugins=grpc,paths=source_relative:gen/v1 auth.proto

# go_gateway使用
protoc -I=D:\GitHub\project\coolcar\server\auth\api --grpc-gateway_out=paths=source_relative,grpc_api_configuration=auth.yaml:gen/v1 auth.proto

# mongod.exe --dbpath=..\data\db