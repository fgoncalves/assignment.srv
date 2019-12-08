OUT:="assigner"

GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/assignment/assignment.proto

.PHONY: build
build: proto

	go build -o ${OUT} main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t assignment-srv:latest

.PHONY: clean
clean:
	rm -f ${OUT}

.PHONY: run
run:

	./${OUT}