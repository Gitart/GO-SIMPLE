#!/bin/bash

clear
NOWF=$(date +"%d-%m-%Y %T")
echo Time $NOWF
echo "Current directory " $PWD
echo Time $NOWF>>log.txt

export GOPATH=$PWD
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export CGO_ENABLED=1
export GOCACHE="/home/airpc/.cache/go-build"
export GOHOSTARCH="amd64"
export GOHOSTOS="linux"
export GOOS="linux"

#echo $PATH
# go get github.com/fatih/color

#go get -v -u github.com/ethereum/go-ethereum/common
#go get -v -u github.com/ethereum/go-ethereum/ethclient
#go test -v . | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''


#echo Start GO programm
go build -o zorg
./zorg -s true
              
