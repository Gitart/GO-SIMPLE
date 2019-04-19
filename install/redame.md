## Инсталяция или обновление

0. Удалить старую версию если надо
```sh
sudo apt purge golang*
```
1. Скачать свежую версию Go
2. Разархивировать в диреткорию
   /home/user/go  
2. Создать файл profile
```sh
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN
```
3. Запустить файл 
```sh
source ./profile
```
4. Проверить 
```sh
go version
```
5. Проверить 
```sh
go env
```
