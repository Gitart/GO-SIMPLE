## Устанавка Geth и Ethereum Swarm

Загружаем исходный код из репозитория:

```
$ mkdir -p $GOPATH/src/github.com/ethereum
$ cd $GOPATH/src/github.com/ethereum
$ git clone https://github.com/ethereum/go-ethereum
$ cd go-ethereum
$ git checkout master
$ go get github.com/ethereum/go-ethereum
```

## Запускаем компиляцию клиента geth и демона swarm:

```
$ go install -v ./cmd/geth
$ go install -v ./cmd/swarm
```

### Проверяем версию установленного geth и swarm:

```
$ geth version
Geth
Version: 1.8.0-unstable
Architecture: amd64
Protocol Versions: [63 62]
Network Id: 1
Go Version: go1.9.2
Operating System: linux
GOPATH=/home/frolov/go
GOROOT=/usr/local/go

$ swarm version
Swarm
Version: 1.8.0-unstable
Network Id: 0
Go Version: go1.9.2
OS: linux
GOPATH=/home/frolov/go
GOROOT=/usr/local/go
```
