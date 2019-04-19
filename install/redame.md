
0. Удалить старую версию если надо
   ```
     sudo apt purge golang*
   ```
1. Скачать свежую версию Go
2. Разархивировать в диреткорию
   /home/user/go  
2. Создать файл profile
```sh
     export GOROOT=$HOME/go
     export PATH=$PATH:$GOROOT/bin  
```
3. Запустить файл 
```
source ./profile
```
4. Проверить 
```
go version
```
5. Проверить 
```
go env
```
