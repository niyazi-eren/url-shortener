#!/bin/bash
snap install go --classic
export GOROOT=/usr/local/go
export GOPATH=$HOME
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
git clone https://github.com/niyazi-eren/url-shortener.git
cd url-shortener
go build -buildvcs=false
./url-shortener