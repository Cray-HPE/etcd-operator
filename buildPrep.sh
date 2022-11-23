#!/bin/bash

set -e

GO_VERSION="1.18.8"
export PATH=$HOME/sdk/go${GO_VERSION}/bin:$HOME/go/bin/go/bin:$PATH
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
  wget https://dl.google.com/go/go1.18.8.linux-amd64.tar.gz
  mkdir -p /usr/local/go${GO_VERSION}
  tar -C /usr/local/go${GO_VERSION} -xzf go1.18.8.linux-amd64.tar.gz
fi

INSTALLED_GO_VERSION=$(go version | awk '{print $3}')

echo "INSTALLED_GO_VERSION: $INSTALLED_GO_VERSION"
echo "GOPATH: $GOPATH"

if [[ "go${GO_VERSION}" !=  $INSTALLED_GO_VERSION ]]; then
    echo "Attempting to switch go from version ${INSTALLED_GO_VERSION} to ${GO_VERSION}"
    go get golang.org/dl/go$GO_VERSION || true
    /home/jenkins/go/bin/go$GO_VERSION download || true
fi

GO_EXEC=$(which go)
export GOPATH=$HOME/sdk/go$GO_VERSION
echo "GO_EXEC: $GO_EXEC"
echo "GOPATH: $GOPATH"

mkdir -p $GOPATH/bin
mkdir -p $GOPATH/src
mkdir -p $GOPATH/pkg

if which dep; then
  echo "dep is installed."
else
  echo "dep wasn't installed, trying to install"
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi

ORIGINAL_DIR=$PWD
mkdir -p $GOPATH/src/github.com/coreos/etcd-operator
$(cp -r . $GOPATH/src/github.com/coreos/etcd-operator | true)
cd $GOPATH/src/github.com/coreos/etcd-operator

$GO_EXEC version
$GO_EXEC get -v -t -d ./...
$GO_EXEC mod vendor
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/operator/build
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/backup-operator/build
$GOPATH/src/github.com/coreos/etcd-operator/hack/build/restore-operator/build
cp -r $GOPATH/src/github.com/coreos/etcd-operator/_output $ORIGINAL_DIR
