.PHONY: proto
proto:
	protoc -I proto proto/auth/v1/*.proto --go_out=. --go-grpc_out=. 