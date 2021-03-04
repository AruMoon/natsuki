#!/bin/bash
wget https://dl.google.com/go/go$GOVERSION.linux-amd64.tar.gz
tar -zxvf go$GOVERSION.linux-amd64.tar.gz
export GOPATH=$HOME/project/go
export GOROOT=$GOPATH/bin
export PATH=$GOROOT:$PATH