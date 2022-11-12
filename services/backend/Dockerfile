FROM golang:1.18-bullseye


# TODO: Multi-step build


RUN apt update && apt install scons build-essential -y

WORKDIR /go/src

ENTRYPOINT [ "go", "run", "main.go" ]