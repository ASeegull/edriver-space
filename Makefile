IMG ?= edriverapp:develop


all: build

build:
	go build cmd/app/main.go

docker-build:
	docker build . -t ${IMG} 