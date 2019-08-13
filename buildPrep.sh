#!/bin/bash

set -e

yum install -y go

GO_VERSION="1.11.5"
INSTALLED_GO_VERSION=$(go version | awk '{print $3}')

if [[ "go${GO_VERSION}" !=  $INSTALLED_GO_VERSION ]]; then
    echo "Attempting to upgrade go from version ${INSTALLED_GO_VERSION} to ${GO_VERSION}"
    go get golang.org/dl/go$GO_VERSION || true
    $GOPATH/bin/go$GO_VERSION download || true
    GO_EXEC=$(which go)
    rm -f $GO_EXEC
    cp $GOPATH/bin/go$GO_VERSION $GO_EXEC
fi

mkdir -p $GOPATH/bin
mkdir -p $GOPATH/src
mkdir -p $GOPATH/pkg


hack/build/build