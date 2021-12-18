# Read about multistage images
FROM golang:1.17-alpine as builder

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum
COPY vendor/ vendor/

COPY cmd/ cmd/
COPY api/ api/
COPY config/ config/
COPY logger/ logger/
COPY migrations/ migrations/
COPY model/ model/
COPY storage/ storage/

# build stage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o app cmd/app/main.go

# actual image
FROM alpine:3.14

LABEL GROUP Lv-644.Golang

RUN apk add --update --no-cache ca-certificates
WORKDIR /usr/lib/edriverspace
COPY --from=builder /src/app /usr/lib/edriver-space/app
RUN chmod +x /usr/lib/edriver-space/app

ENTRYPOINT [ "/usr/lib/edriver-space/app" ]

USER app
EXPOSE 5050
