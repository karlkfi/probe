FROM golang:1.6.2-wheezy
MAINTAINER Karl Isenberg <karlkfi@yahoo.com>

RUN echo "deb-src http://httpredir.debian.org/debian wheezy main" >> /etc/apt/sources.list && \
    echo "deb-src http://httpredir.debian.org/debian wheezy-updates main" >> /etc/apt/sources.list && \
    apt-get -yqq update && \
    apt-get -yqq upgrade && \
    DEBIAN_FRONTEND=noninteractive apt-get -yqq install --no-install-recommends \
      build-essential git && \
    apt-get -y clean

RUN go get github.com/tools/godep
RUN go get github.com/mitchellh/gox
ENV GO15VENDOREXPERIMENT 0

VOLUME ${GOPATH}/src/github.com/karlkfi/probe

WORKDIR ${GOPATH}/src/github.com/karlkfi/probe

CMD ["make"]
