#!/bin/bash

# Variables for only read
runtime="github.com/yael-castro/jikkosoft/3/internal/runtime"
commit=$(git log --pretty=format:'%h' -n 1 || echo 'unknown')

ldflags=""
options=""

function build() {
    cd "./cmd/$binary" || exit

    if ! go mod tidy
    then
      exit 1
    fi

    if ! go build \
      -o ../../build/ \
      -tags "$tags" \
      -ldflags "$ldflags" \
      "$options"
    then
      exit 1
    fi

    cd ../../

    echo "Commit: '$commit'"
    echo "MD5 checksum: $(md5sum "build/$binary")"
    echo "Success build"
    exit
}


ldflags="-X $runtime.GitCommit=$commit"
binary="http"
tags="http"

printf "\nBuilding API REST in \"build\" directory\n"
CGO_ENABLED=0 build