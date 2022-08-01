FROM golang:alpine3.16

WORKDIR /src

COPY . /src

RUN go mod download
