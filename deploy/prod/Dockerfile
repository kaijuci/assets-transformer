FROM golang:1.20-alpine

RUN apk add g++ gcc imagemagick-dev

ARG VERSION
ENV VERSION=$VERSION

WORKDIR /app
ADD . .

# todo: build statically linked binary -- requires from source ImageMagick static libs etc.

RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build \
    -ldflags "-X github.com/kaijuci/transformer/vars.Version=$VERSION" \
    -o xfrmr cmd/main.go

CMD ["/app/xfrmr"]