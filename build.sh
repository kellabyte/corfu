#!/usr/bin/env bash

export GOPATH=$(pwd)

go get github.com/Sirupsen/logrus
go get github.com/boltdb/bolt
go get github.com/msgpack-rpc/msgpack-rpc-go/rpc
go get github.com/paulbellamy/ratecounter

go build github.com/kellabyte/corfu/cmd/corfu-store
go build github.com/kellabyte/corfu/cmd/corfu-token
