#!/bin/bash
snap install go --classic
export GOROOT=/usr/local/go
export GOPATH=$HOME
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export PUBLIC_DNS=$(curl http://169.254.169.254/latest/meta-data/public-hostname)
git clone https://github.com/niyazi-eren/url-shortener.git
cd url-shortener
sudo go build -buildvcs=false
./url-shortener