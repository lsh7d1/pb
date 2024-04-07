# API_PROTO_FILES，查找所有.proto文件
API_PROTO_FILES=$(find api -name "*.proto")

# 使用protoc编译器编译协议缓冲区文件
protoc --proto_path=./api \
        --go_out=paths=source_relative:./api \
        --go-http_out=paths=source_relative:./api \
        --go-grpc_out=paths=source_relative:./api \
        --openapi_out=paths=source_relative:. \
        ${API_PROTO_FILES}
