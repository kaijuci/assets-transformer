FROM kaijuci/gomagick:1.0 AS base
LABEL org.opencontainers.image.authors="sumeet@kaiju.ci"

WORKDIR /app
COPY go.mod go.sum ./

# fetch dependencies
RUN go mod download

FROM base AS go-build

COPY . .

RUN tree .

# build
RUN CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build ./...

# test
CMD CGO_CFLAGS_ALLOW='-Xpreprocessor' \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go test -v -cover ./...
