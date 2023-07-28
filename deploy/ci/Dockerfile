FROM golang:1.20.6-alpine
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# https://pkgs.alpinelinux.org/package/edge/community/x86/imagemagick

# install imagemagick and others
RUN apk update && apk -U --no-cache add \
    build-base \
    tree \
    imagemagick imagemagick-dev

WORKDIR /app
COPY . .

RUN tree .

# check imagemagick installation
RUN pkg-config --cflags --libs MagickWand

# fetch dependencies
RUN go get -u -t -v all

# verify imagemagick
RUN convert --version

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