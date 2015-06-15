#!/usr/bin/env bash

mkdir tools
cd tools
git clone https://github.com/Masterminds/glide.git
cd glide
make bootstrap
cd ../../

export GOPATH=$(pwd)/_vendor
tools/glide/glide install

go build github.com/kellabyte/corfu/cmd/corfu-store
go build github.com/kellabyte/corfu/cmd/corfu-token
