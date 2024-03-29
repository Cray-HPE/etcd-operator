#!/bin/bash

set -e

export PATH=$PATH:$HOME/go/bin/go/bin
: "${GOPATH:=$HOME/go}"

if which git ; then
  echo "git is installed."
else
  echo "git wasn't installed, trying to install"
  yum install -y git
fi


if which go ; then
  echo "Go is installed."
else
  echo "Go wasn't installed, trying to install"
  wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz
  tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
fi

GO_VERSION="1.15.6"
INSTALLED_GO_VERSION=$(go version | awk '{print $3}')

if [[ "go${GO_VERSION}" !=  $INSTALLED_GO_VERSION ]]; then
    echo "Attempting to switch go from version ${INSTALLED_GO_VERSION} to ${GO_VERSION}"
    go get golang.org/dl/go$GO_VERSION || true
    $GOPATH/bin/go$GO_VERSION download || true
    GO_EXEC=$(which go)
fi

mkdir -p $GOPATH/bin
mkdir -p $GOPATH/src
mkdir -p $GOPATH/pkg

if which dep; then
  echo "dep is installed."
else
  echo "dep wasn't installed, trying to install"
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | DEP_RELEASE_TAG=v0.5.4 sh
fi

ORIGINAL_DIR=$PWD
mkdir -p $GOPATH/src/github.com/coreos/etcd-operator
$(cp -r . $GOPATH/src/github.com/coreos/etcd-operator | true)
cd $GOPATH/src/github.com/coreos/etcd-operator

go version
go get -v -t -d ./...
go mod vendor
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/operator/build
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/backup-operator/build
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/restore-operator/build
cp -r $GOPATH/src/github.com/coreos/etcd-operator/_output $ORIGINAL_DIR
