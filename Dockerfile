FROM golang:1.22.2

RUN apt update \
  && apt install -y vim

  WORKDIR /go/src/work