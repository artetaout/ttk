#!/bin/bash

set -ex

apt update
apt install -y tmux protobuf-compiler vim make curl

cd ~
apt update && apt install -y wget vim
wget -c https://dl.google.com/go/go1.17.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz 

mkdir -p ~/go_repos

cat >> ~/.bashrc << "EOF"
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=$HOME/go_repos
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOBIN:$PATH
EOF


source ~/.bashrc

go version

cat >> ~/.bashrc << "EOF"
export GO111MODULE=on
export GOSUMDB="sum.golang.google.cn"
EOF

source ~/.bashrc
go env -w GOPROXY=https://goproxy.cn,direct
go install github.com/go-delve/delve/cmd/dlv@v1.8.0
go install github.com/cweill/gotests/gotests@v1.6.0
go install -v github.com/fatih/gomodifytags@v1.16.0
go install -v github.com/josharian/impl@v1.1.0
go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0
# go install -v honnef.co/go/tools/cmd/staticcheck@latest
go install -v golang.org/x/tools/gopls@latest
