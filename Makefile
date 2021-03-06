.PHONY: all build run deps test-cov format

all: build

build:
	go build -o aut

run:
	go run -race main.go

mod:
	go mod download

deps: mod
	( cd /tmp; go get \
		github.com/golang/mock/mockgen \
		)

test-cov:
	go test -coverprofile=cover.txt ./... && go tool cover -html=cover.txt -o cover.html

format:
	go fmt ./...

docker-build:
	 docker build -t authorizer .

docker-run:
	 docker run -i authorizer
