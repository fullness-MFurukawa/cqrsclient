FROM golang:1.21.3-alpine3.17
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/client
WORKDIR /go/src/client
ADD . /go/src/client