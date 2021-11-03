# Create go mode


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

go build -o nats
./nats

# timeout 5m
sleep 5

echo Finish        
```

                         
