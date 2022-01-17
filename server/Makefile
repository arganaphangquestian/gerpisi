gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false ./data/data.proto
run:
	@go build -o appserver main.go ; ./app