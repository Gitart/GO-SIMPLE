
Теперь при указании параметров -tlsAddr, -tlsCertFile и -tlsKeyFile сервер будет 
принимать https-запросы на -tlsAddr дополнительно к http-запросам на -addr. И снова никаких nginx’ов c 
apache’ами и openssl’ами не нужно. Скорость обработки https-трафика сервером на Go сравнима со скоростью
nginx, поэтому перед ним не нужно ставить TLS termination proxy.
