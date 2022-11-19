FROM golang:1.18-bullseye


RUN addgroup --gid 1000 "cupcake"
RUN adduser --quiet "cupcake" --uid 1000 --gid 1000
RUN mkdir /opt/cupcake
RUN chown cupcake:cupcake /opt/cupcake

RUN apt update && apt install --no-install-recommends tini scons build-essential -y

USER "cupcake"
ENV TZ="America/Sao_Paulo"
WORKDIR /opt/cupcake
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["sleep", "infinity"]
