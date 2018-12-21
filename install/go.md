# Установка Go

Очень важно перед началом работ установить Go новой версии, не ниже 1.9.2. 
В репозитории Debian и Ubuntu могут быть старые версии Go, поэтому устанавливаем из исходников.

Скачиваем и распаковываем исходники под непривилегированным пользователем:

```sh
$ curl -O https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
```

### Создаем у пользователя каталог go и устанавливаем переменные окружения:

```sh
$ mkdir -p ~/go; echo "export GOPATH=$HOME/go" >> ~/.bashrc
$ echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc
$ source ~/.bashrc
```

#### Проверка переменных окружения :

```
$ printenv | grep go
````

Responce
```
PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games:/home/frolov/go/bin:/usr/local/go/bin
GOPATH=/home/frolov/go
```
