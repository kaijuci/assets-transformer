FROM golang:1.20.6-alpine
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# https://pkgs.alpinelinux.org/package/edge/community/x86/imagemagick

# install imagemagick and others
RUN apk update && apk -U --no-cache add \
    tree \
    imagemagick libmagickwand-dev

WORKDIR /app
COPY . /app

RUN tree .

# check imagemagick installation
RUN pkg-config --cflags --libs MagickWand

# fetch dependencies
RUN go get -u -t -v all

# test
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v ./...

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build -v cmd/*.go -o main

CMD ["/app/main"]