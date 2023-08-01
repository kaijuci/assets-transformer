FROM kaijuci/gomagick:2.0 AS base
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

WORKDIR /app
COPY go.mod go.sum ./

# fetch dependencies
RUN go mod download -x

FROM base AS go-build

COPY . .

# todo: build statically linked binary w/ imagemagick

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build \
    -ldflags "-X github.com/kaijuci/transformer/vars.Version=$(git rev-parse --short HEAD)" \
    -o xfrmr \
    cmd/main.go

# echo
RUN /app/xfrmr version

# todo: add another stage and copy built executable from go-build to that stage

CMD ["/app/xfrmr"]