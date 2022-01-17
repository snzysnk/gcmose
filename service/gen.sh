PROTO_PATH=./proto
API_PATH=./auth/api
mkdir -p $API_PATH

#window命令行执行
protoc -I=./proto --go_out=paths=source_relative:./auth/api --go-grpc_out=paths=source_relative:./auth/api wechat.proto
protoc -I=./proto --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./proto/wechat.yaml:./auth/api wechat.proto

protoc -I=./proto --go_out=paths=source_relative:./rental/api --go-grpc_out=paths=source_relative:./rental/api trip.proto
protoc -I=./proto --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./proto/trip.yaml:./rental/api trip.proto

#linux下执行
#protoc -I=$PROTO_PATH --go_out=paths=source_relative:$API_PATH --go-grpc_out=paths=source_relative:$API_PATH wechat.proto
#protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/wechat.yaml:$API_PATH wechat.proto

#function CreateProto(){
#  PROTO_PATH=./proto
#  API_PATH=$1
#  echo $API_PATH
#  mkdir -p $API_PATH
#}
#
#CreateProto ./rental/api
