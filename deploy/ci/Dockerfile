FROM kaijuci/gomagick:2.0 AS base
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

# install deps
RUN apt-get update \
    && apt-get install -y \
    tree

WORKDIR /app
COPY go.mod go.sum ./

# fetch dependencies
RUN go mod download -x

FROM base AS go-build

COPY . .

RUN tree .

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build \
    -ldflags "-X github.com/kaijuci/assets-transformer/vars.Version=$(git rev-parse --short HEAD)" \
    -o xfrmr \
    cmd/main.go

# test
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v -cover ./... 2>&1 | tee results.out

# print report
RUN cat results.out

# echo version - mostly to satisfy kaniko
ENTRYPOINT ["/app/xfrmr"]
