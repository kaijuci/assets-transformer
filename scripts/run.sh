#!/usr/bin/env bash

go run -ldflags "-X github.com/kaijuci/assets-transformer/vars.Version=$(git rev-parse --short HEAD)" \
   cmd/main.go "$@"