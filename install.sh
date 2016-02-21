#!/bin/bash
# set up GOPATH and GOBIN for this project
export GOPATH=$(pwd -P)
export GOBIN=$GOPATH/bin
mkdir -p $GOBIN
export PATH=$GOBIN:$PATH
# now build and install the funtodone executable in our bin directory
go install -v github.com/rlunde/funtodone
