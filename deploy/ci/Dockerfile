FROM golang:1.20.6-alpine AS base
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# https://pkgs.alpinelinux.org/package/edge/community/x86/imagemagick

WORKDIR /app
COPY go.mod go.sum ./

# fetch dependencies
RUN go get -u -t -v all

FROM base AS go-imagemagick

# install imagemagick and others
RUN apk update && apk -U --no-cache add \
    build-base \
    tree \
    imagemagick imagemagick-dev

# verify imagemagick
RUN convert --version

# check imagemagick installation
RUN pkg-config --cflags --libs MagickWand

FROM go-imagemagick AS client

COPY . .

RUN tree .

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build -v ./...

# test
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v ./...

CMD ["/app/main"]