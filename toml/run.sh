#!/bin/bash

clear
echo Start build process ...
echo "Start :" $(date +"%x %T.%6N")
echo

# echo $HOME
# echo $HOSTNAME
# echo $HOSTTYPE
# echo $OLDPWD
# echo $OSTYPE
# echo $PATH
# echo $PPID
# echo $SECONDS

#export GOPATH=$PWD
#export GOROOT=$HOME/go
export BUILD_NAME=toml
export GOMODCACHE=$PWD/package
export GO111MODULE=auto
go mod tidy

go build -o $BUILD_NAME

./$BUILD_NAME --log=false

# timeout 5m
# sleep 5
# echo Finish

