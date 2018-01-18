FROM golang:1.9.2-alpine
MAINTAINER Karl Isenberg <karlkfi@yahoo.com>

RUN apk update && \
    apk upgrade && \
    apk add git \
            build-base

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/mitchellh/gox

VOLUME ${GOPATH}/src/github.com/karlkfi/probe

WORKDIR ${GOPATH}/src/github.com/karlkfi/probe

CMD ["make"]
