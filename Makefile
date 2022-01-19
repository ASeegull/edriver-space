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

run:
	go run cmd/app/main.go

start:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o app cmd/app/main.go
	docker-compose up -d --build

stop:
	docker-compose stop

restart:
	make stop
	make start

test:
	go test ./...

test-coverage:
	- rm -rf *.out  # Remove all coverage files if exists
	go test -race -failfast -tags=integration -coverprofile=coverage-all.out ./...

gen:
	mockgen -source=service/service.go -destination=service/mocks/mock.go
	mockgen -source=repository/repository.go -destination=repository/mocks/mock.go
