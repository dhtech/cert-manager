all: dns.pb.go

dns.pb.go: dns.proto
	protoc --go_out=plugins=grpc:. $?
