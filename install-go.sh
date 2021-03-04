#!/bin/bash
wget https://dl.google.com/go/go$GOVERSION.linux-amd64.tar.gz
tar -zxvf go$GOVERSION.linux-amd64.tar.gz
export PATH=$HOME/project/go:$PATH