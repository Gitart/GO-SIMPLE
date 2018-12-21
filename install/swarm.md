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


# Запуск демона swarm

Здесь вам потребуется адрес аккаунта, созданного на этапе инициализации узла. Если вы его не сохранили, ничего страшного. Просто зайтите в консоль geth, открытую с помощью скрипта attach_node.sh, и выдайте там следующую команду:

```
> web3.eth.accounts
["0xcd9fcb450c858d1a7678a2bccf36ea5decd2b09b"]
```

Команда покажет адрес созданных учетных записей. Сразу после инициализации там будет только один адрес.

Для запуска демона Ethereum Swarm в режиме единственного узла (Singleton) подготовьте командный файл swarm_start.sh:

```
swarm --bzzaccount cd9fcb450c858d1a7678a2bccf36ea5decd2b09b --datadir "/home/ethertest/data/node1" --maxpeers 0 -nodiscover --verbosity 4 --ens-api /home/ethertest/data/node1/geth.ipc
```

Укажите в нем адрес аккаунта, созданного на вашем узле, без «0x».

При запуске демона вам будет нужно ввести пароль от созданного ранее аккаунта:
```
$ sh swarm_start.sh
Unlocking swarm account 0xCD9Fcb450C858D1A7678a2bCCf36EA5decd2B09B [1/3]
Passphrase:
```

## Загружаем файл в Ethereum Swarm

Проще всего загрузить файл с помощью команды swarm с параметром up. Дополнительно нужно указать путь к загружаемому файлу:

```
$ swarm up start_node.sh
f2073b8f0cf0cfe1e165060882da71a37bb6fd97bdec6be71b4f66ebcf0aba9f
```

Данная команда вернет хеш загруженного файла. Хеш можно использовать для чтения файла. Вы можете это сделать с помощью команд wget или curl:

```
$ wget http://localhost:8500/bzz:/f2073b8f0cf0cfe1e165060882da71a37bb6fd97bdec6be71b4f66ebcf0aba9f/
$ curl http://localhost:8500/bzz:/f2073b8f0cf0cfe1e165060882da71a37bb6fd97bdec6be71b4f66ebcf0aba9f/
```


Команда wget позволяет сохранить содержимое файла на локальном диске. Используйте параметр -O, чтобы задать имя файла. Команда curl выведет содержимое файла на консоль, поэтому в таком виде ее не следует использовать для просмотра содержимого бинарных файлов. В конце URL необходим слеш, иначе произойдет редирект.

Когда файл загружается в Ethereum Swarm описанным выше образом, для него создается и сохраняется так называемый манифест. Это заголовок, описывающий содержимое, доступное в хранилище по заданному идентификатору.

Ниже мы загрузили файл Net-Ethereum-0.28.tar.gz с помощью простой команды swarm up:

```
$ swarm up Net-Ethereum-0.28.tar.gz
8da3713d49c62740f5ab594b06173975ac97cb3dd3848ae996484ec264a10e2f
```

Теперь, указав протокол bzz-list, мы можем просмотреть манифест:

```
$ curl http://localhost:8500/bzz-list:/8da3713d49c62740f5ab594b06173975ac97cb3dd3848ae996484ec264a10e2f/
{"entries":[{"hash":"543ee6e744f93de76ac132b8ab71982e32beaf90d1005e771dde003b2a4a54c3","path":"/","contentType":"application/gzip","mode":420,"size":12403,"mod_time":"2018-01-13T14:57:54+03:00"}]}
```

## Манифест будет показан в формате JSON.
В манифесте хранится пусть к файлу (имя файла), его размер, тип (Content Type), дата и время модификации, а также хеш файла.
Чтобы извлечь содержимое файла по его идентификатору и сохранить под именем t.tar.gz, сделайте так:

```
$ wget -O t.tar.gz http://localhost:8500/bzz:/8da3713d49c62740f5ab594b06173975ac97cb3dd3848ae996484ec264a10e2f/ 
```



## Загрузка каталогов с подкаталогами

Для рекурсивной загрузки каталога вместе с его содержимым в хранилище Ethereum Swarm укажите параметр --recursive:

```
$ swarm --recursive up Net-Ethereum/
4fb1f2270381c022461037151f70ce081082f0ae1a2a23d8c7ea602da69b4115
```

В манифесте будет показана информация обо всех файлах загруженного подкаталога:

Link
https://habr.com/post/346542/
