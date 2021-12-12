IMG ?= edriverapp:develop

all: build

build:
	go build cmd/app/main.go

docker-build:
	docker build . -t ${IMG} 

staticcheck: 
	go get -u honnef.co/go/tools/cmd/staticcheck
	
lint:
	staticcheck ./...

test:
	go test -race -failfast ./...
