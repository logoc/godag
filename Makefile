# Makefile
GOPATH:=$(shell go env GOPATH)

# proto1 相关指令
.PHONY: proto1
path1:
	protoc --go_out=. proto1/greeter/greeter.proto
path2:
	protoc -I. --go_out=. proto1/greeter/greeter.proto # 点号表示当前路径，注意-I参数没有等于号
path3:
	protoc --proto_path=. --go_out=. proto1/greeter/greeter.proto
