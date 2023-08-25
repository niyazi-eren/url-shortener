#!/bin/bash
snap install go --classic
export GOROOT=/usr/local/go
export GOPATH=$HOME
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
# get public host name from route53
export PUBLIC_DNS=$(curl http://169.254.169.254/latest/meta-data/public-hostname)
git clone https://github.com/niyazi-eren/url-shortener.git
cd url-shortener
sudo go build -buildvcs=false
./url-shortener &

# Install npm & node
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash
apt-get install -y nodejs

cd web/app

# create env file
touch .env
chmod u+w .env
echo "VITE_PUBLIC_DNS=\"$PUBLIC_DNS\"" > .env
echo "VITE_PORT=\":8080\"" >> .env

npm i vite
npm i
npm i -D @zerodevx/svelte-toast
npm i -g http-server
npm run build

cd dist

http-server -p 80
