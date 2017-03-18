## Инсталяция RethinkDB

### Выполнить команду копированием и вставкой в Puty одной коммандой
https://rethinkdb.com/docs/install/ubuntu/ 

```sh
source /etc/lsb-release && echo "deb http://download.rethinkdb.com/apt $DISTRIB_CODENAME main" | sudo tee /etc/apt/sources.list.d/rethinkdb.list
wget -qO- https://download.rethinkdb.com/apt/pubkey.gpg | sudo apt-key add -
```

### 2. Обновить все

```sh
sudo apt-get update
sudo apt-get install rethinkdb
```

### 3. Создание конфигурационного файла :   
https://rethinkdb.com/docs/start-on-startup/ 

```sh
sudo cp /etc/rethinkdb/default.conf.sample /etc/rethinkdb/instances.d/instance1.conf
sudo vim /etc/rethinkdb/instances.d/instance1.conf
```

### 4. Старт сервера  

```sh
sudo /etc/init.d/rethinkdb restart
```

### 5. Изменить конфиг - для того что-бы увидеть базу по HTTP: 
```sh
sudo vim /etc/rethinkdb/instances.d/instance1.conf
```

### 6. Добавить :
Для доступа к базе по HTTP

```sh
bind=all
```

### 7. Выйти из vim (нажатие shift + :)
```sh
:wq to write and quit (think write and quit)
```



