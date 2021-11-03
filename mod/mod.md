# Create go mode

## Basic steps for used go mod

```txt
1. Cretae directory 
2. go mod init nameyorapp
3. export GOMODCACHE=$PWD/package
4. go mod tidy
5. go build -o appname
```


### runmode
```sh
#!/bin/bash

export GOMODCACHE=$PWD/package
export GO111MODULE=auto

go mod init statistics
go mod tidy
```


## run

```sh

clear
echo start build process ...

export GOMODCACHE=$PWD/package
export GO111MODULE=auto
go mod tidy 

go build -o nats
./nats

# timeout 5m
sleep 5

echo Finish        
```

                         
