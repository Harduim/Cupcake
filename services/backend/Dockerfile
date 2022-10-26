FROM golang:1.18-alpine3.16

RUN apk add --update bash gcc g++ scons

WORKDIR /go/src

ENTRYPOINT [ "go", "run", "main.go" ]