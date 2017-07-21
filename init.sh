#!/bin/bash
go get -u github.com/golang/dep/cmd/dep
cd $GOPATH/src/server
dep ensure
cd $GOPATH/src/validator
dep ensure
cd $GOPATH/src/headless-operator
dep ensure
