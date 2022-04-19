start:
	@go run main.go start --config .env

migrate-create:
	@migrate create -ext sql -dir migration -seq init_schema

migrate-up:
	@migrate -path ./migration -database "mysql://user:password@tcp(localhost:3306)/minipos?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migrate-down:
	@migrate -path ./migration -database "mysql://user:password@tcp(localhost:3306)/minipos?charset=utf8mb4&parseTime=True&loc=Local" -verbose down

startdb:
	@docker container start mysqlc

make genproto:
	@protoc --proto_path=api/proto -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
	--proto_path=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
	--go_out=plugins=grpc:api/pb \
	--grpc-gateway_out=logtostderr=true:api/pb \
	--swagger_out=logtostderr=true:api/swagger \
	./api/proto/*.proto