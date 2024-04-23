# 生成rpc 服务端 和 客户端
protoc -I=/Users/heping/Documents/GitHub/coolcar/server/rental/api --go-grpc_out=paths=source_relative:gen/v1 rental.proto

# 对proto中数据结构生成go 代码
protoc -I=/Users/heping/Documents/GitHub/coolcar/server/rental/api --go_out=paths=source_relative:gen/v1 rental.proto
#protoc -I=/Users/heping/Documents/GitHub/coolcar/server/rental/api --go_out=plugins=grpc,paths=source_relative:gen/v1 rental.proto

# go_gateway使用
protoc -I=/Users/heping/Documents/GitHub/coolcar/server/rental/api --grpc-gateway_out=paths=source_relative,grpc_api_configuration=rental.yaml:gen/v1 rental.proto

# mongod.exe --dbpath=..\data\db