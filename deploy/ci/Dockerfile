FROM golang:1.20.6-alpine
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# https://pkgs.alpinelinux.org/package/edge/community/x86/imagemagick
ENV IMAGEMAGICK_VERSION=7.1.1.14-r0

RUN apk update && apk -U --no-cache add \
    tree \
    imagemagick 

WORKDIR /work
COPY . /work

RUN tree .

# fetch dependencies
RUN go get -u -t -v all

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v ./...

# test
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v ./...
