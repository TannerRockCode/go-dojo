.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt: 
	go fmt ./...

vet: fmt 
	go vet ./...

build: vet 
	go build -o /home/tanner/repos/go-dojo/bin/literals_practice

