**Убунту (Ксениал) 16.04        **
Зависимости сборки      
В этом документе по установке предполагается Ubuntu 16.04+ на платформе x86-64.       

**Установить Git**
копия$ sudo apt-get install git             
Установить Go 1.11+           
Загрузите Go 1.10+ с https://golang.org/dl/ .          


копия$ wget https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.11.1.linux-amd64.tar.gz

**Настройка PATH**  
Добавьте путь к вашему ~/.bashrc.   

копияexport PATH=$PATH:/usr/local/go/bin:${HOME}/go/bin

**Источник новой среды**  
копия$ source ~/.bashrc

**Тестирование всего этого**
копия$ go env
$ go version
