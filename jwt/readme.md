 https://jwt.io/
 ## Основные принципы работы JWT  
     * Стандартный метод подписывается с ключом  - "keysecret2"
     * Потом вычитывается с тем же ключом 
     * Если ключ был изменен то при проверке генерируется ощибка  "token signature is invalid"
     * Но при этом все данные зашитые в токене можно прочитать  из структуры в которую были зашиты данные или кастомные или стандартные
     * Просроченный ключ тоже выдает ошибку если его использовать
     * позже указанного срока: "token is expired by 5.044608s"