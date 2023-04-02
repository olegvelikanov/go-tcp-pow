generate:
	protoc -I=. --go_out=internal/pkg/contract/pb ./proto/contract.proto