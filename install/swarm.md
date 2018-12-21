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


## Подготовка приватного блокчейна для запуска Ethereum Swarm

Прежде всего, создайте в домашнем каталоге пользователя файл genesis.json:
```json
{
  "config": {
     "chainId": 1907,
     "homesteadBlock": 0,
     "eip155Block": 0,
     "eip158Block": 0
  },
  "difficulty": "40",
  "gasLimit": "5100000",
  "alloc": {}
}
```

#### Далее создайте в домашнем каталоге подкаталог node1:

```
$ mkdir node1
```

Инициализацию узла можно сделать при помощи пакетного файла init_node.sh:

```
geth --datadir node1 account new
geth --datadir node1 init genesis.json
```

При запуске этого файла будет создан аккаунт и запрошен пароль, который вам необходимо сохранить в безопасном месте.

Создайте файл start_node.sh для запуска узла:
```
geth --datadir node1 --nodiscover --mine --minerthreads 1 --maxpeers 0 --verbosity 3 --networkid 98765 --rpc --rpcapi="db,eth,net,web3,personal,web3" console
```

Запустите этот файл и дождитесь завершения генерации DAG.

С помощью файла attach_node.sh вы сможете открыть консоль geth и подключиться к приватному узлу:
```
geth --datadir node1 --networkid 98765 attach ipc://home/frolov/node1/geth.ipc
```

