#!/usr/bin/env bash

export GOPATH=$(pwd)
go build github.com/kellabyte/corfu/cmd/corfu-store
go build github.com/kellabyte/corfu/cmd/corfu-token
