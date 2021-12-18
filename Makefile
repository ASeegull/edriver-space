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
	go test ./...

test-coverage:
	- rm -rf *.out  # Remove all coverage files if exists
	go test -race -failfast -tags=integration -coverprofile=coverage-all.out ./...
