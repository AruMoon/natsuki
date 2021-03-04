#!/bin/bash
wget https://dl.google.com/go/go$GOVERSION.linux-amd64.tar.gz
tar -zxvf go$GOVERSION.linux-amd64.tar.gz
export GOROOT=$HOME/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH