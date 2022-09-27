## Example work with MQTT 

![image](https://user-images.githubusercontent.com/3950155/192479182-7a6d98ae-2e40-4835-a9bf-7ab84c80a3ac.png)


---
1. ClientId must be different in other clients (**Important**)
2. Must be run server MQTT (for windows mosquitto.exe)
3. Run **subscribe** client listener publish client
4. Run **publish** client 




## Что такое MQTT?

MQTT **расшифровывается** как **M** essage **Queuing** Telemetry **Transport** .  MQTT — это простой протокол обмена сообщениями, разработанный для ограниченных устройств с низкой пропускной способностью. Таким образом, это идеальное решение для обмена данными между несколькими IoT-устройствами.

Связь MQTT работает как система *публикации* и *подписки* . Устройства *публикуют* сообщения на определенную тему. Все устройства, которые *подписаны* на эту тему, получают сообщение.

![Публикация MQTT Подписаться](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2018/09/Publish-subscribe-MQTT.png?resize=268%2C149&quality=100&strip=all&ssl=1)

Его основные приложения включают в себя отправку сообщений для управления выходами, чтение и публикацию данных с сенсорных узлов и многое другое.

## Основные понятия MQTT

В MQTT есть несколько основных понятий, которые вам необходимо понять:

*   [Опубликовать/подписаться](https://randomnerdtutorials.com/what-is-mqtt-and-how-it-works/#publish-subscribe)
*   [Сообщения](https://randomnerdtutorials.com/what-is-mqtt-and-how-it-works/#messages)
*   [Темы](https://randomnerdtutorials.com/what-is-mqtt-and-how-it-works/#topics)
*   [Маклер](https://randomnerdtutorials.com/what-is-mqtt-and-how-it-works/#broker)

## MQTT — публикация/подписка

Первая концепция — это система *публикации и подписки* . В системе публикации и подписки устройство может публиковать сообщения в теме или подписываться на определенную тему для получения сообщений.

![Опубликовать Подписаться Тема MQTT](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2018/09/Publish-Subscribe-Topic-MQTT.png?resize=743%2C125&quality=100&strip=all&ssl=1)

*   Например , **Устройство 1** публикует в теме.
*   **Устройство 2** подписано на ту же тему, в которой публикуется **устройство 1 .**
*   Итак, **устройство 2** получает сообщение.

## MQTT — Сообщения

Сообщения — это информация, которой вы хотите обмениваться между вашими устройствами. Это может быть сообщение, например команда, или данные, например, показания датчика.

## MQTT – Темы

Еще одно важное понятие — *темы* . Темы — это то, как вы регистрируете интерес к входящим сообщениям или как вы указываете, где вы хотите опубликовать сообщение.

Темы представлены строками, разделенными косой чертой. Каждая косая черта указывает на уровень темы. Вот пример того, как вы могли бы создать тему для лампы в вашем домашнем офисе:

![](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/5mqtt-topics.jpg?resize=442%2C246&quality=100&strip=all&ssl=1)

**Примечание.** Темы чувствительны к регистру, что отличает эти две темы:

![](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/7case-sensitive-300x188.jpg?resize=300%2C188&quality=100&strip=all&ssl=1)

Если вы хотите включить лампу в своем домашнем офисе с помощью MQTT, вы можете представить себе следующий сценарий:

![Пример публикации подписки MQTT](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2018/09/MQTT-Publish-Subscribe-Example.png?resize=732%2C384&quality=100&strip=all&ssl=1)

1.  Устройство публикует сообщения «включено» и «выключено» в разделе « **дом/офис/лампа** » .
2.  У вас есть устройство, которое управляет лампой (это может быть ESP32, ESP8266 или любая другая плата или устройство). ESP32, который управляет вашей лампой, подписан на ту же тему: **home/office/lamp** .
3.  Таким образом, когда в этой теме публикуется новое сообщение, ESP32 получает сообщения «включено» или «выключено» и включает или выключает лампу.

Устройством, которое публикует сообщения, может быть ESP32, ESP8266 или платформа контроллера домашней автоматизации с поддержкой MQTT, например, Node-RED, Home Assistant, Domoticz или OpenHAB.

![mqtt-устройство](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/device1.png?resize=555%2C247&quality=100&strip=all&ssl=1)

## MQTT — Брокер

Наконец, еще одно важное понятие — *брокер* .

Брокер **MQTT** отвечает за **получение** всех сообщений, **фильтрацию** сообщений, **определение того** , кто в них заинтересован, а затем  за **публикацию**  сообщения всем подписавшимся клиентам.

![mqtt-брокер](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/mqtt_broker.png?resize=750%2C303&quality=100&strip=all&ssl=1)

Есть несколько брокеров, которые вы можете использовать. В проектах домашней автоматизации мы используем **[Mosquitto Broker](https://mosquitto.org/)** , установленный на Raspberry Pi. Вы также можете установить брокера Mosquitto на свой компьютер (что не так удобно, как использование платы Raspberry Pi, потому что вам нужно постоянно держать компьютер включенным, чтобы поддерживать соединение MQTT между вашими устройствами).

![](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/mosquitto-broker.png?resize=200%2C197&quality=100&strip=all&ssl=1)

Установка брокера Mosquitto на Raspberry Pi в вашей локальной сети позволяет вам обмениваться данными между вашими IoT-устройствами, которые также подключены к той же сети.

Чтобы установить брокера Mosquitto на Raspberry Pi, следуйте нашему руководству:

*   [Установите Mosquitto Broker на Raspberry Pi](https://randomnerdtutorials.com/how-to-install-mosquitto-broker-on-raspberry-pi/)

Вы также можете запустить брокера Mosquitto MQTT в облаке. Запуск MQTT Mosquitto Broker в облаке позволяет подключать несколько IoT-устройств из любого места, используя разные сети, если у них есть подключение к Интернету. Проверьте учебник ниже:

*   [Запустите свой облачный брокер MQTT Mosquitto (доступ из любого места с помощью Digital Ocean)](https://randomnerdtutorials.com/cloud-mqtt-mosquitto-broker-access-anywhere-digital-ocean/)

## Как использовать MQTT в проектах домашней автоматизации и IoT

MQTT отлично подходит для проектов домашней автоматизации и Интернета вещей. Вот пример того, как его можно использовать в системе [домашней автоматизации,](https://randomnerdtutorials.com/build-a-home-automation-system-for-100/) созданной с помощью недорогих плат разработки, таких как Raspberry Pi, ESP32, ESP8266 и Arduino.

![Пример системы домашней автоматизации](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2017/01/home-automation-mqtt-example.png?resize=750%2C505&quality=100&strip=all&ssl=1)

*   Raspberry Pi запускает брокера Mosquitto, который необходим для протокола MQTT.

*   Тот же Raspberry Pi работает под управлением Node-RED, платформы домашней автоматизации с поддержкой MQTT — это означает, что он может подписываться на темы для получения сообщений от других устройств IoT и публиковать сообщения в определенных темах для отправки сообщений на другие устройства.

*   Node-RED также позволяет создавать пользовательский интерфейс с кнопками для управления выходными данными и диаграммами для отображения показаний датчиков.

*   Arduino, ESP32 и ESP8266 могут действовать как клиенты MQTT, которые публикуют темы и подписываются на них.

*   Эти платы подключены к исполнительным механизмам, таким как светодиоды или реле, и датчикам, таким как температура, влажность, датчики дыма и т. д.

*   Эти кабаны могут публиковать данные о состоянии сенсора по определенной теме, на которую Node-RED также подписан. Таким образом, Node-RED получает показания датчиков и может отображать их в пользовательском интерфейсе.

*   С другой стороны, Node-RED может публиковать данные по определенной теме для управления выводом при использовании кнопок на интерфейсе. Другие доски также подписаны на эту тему и соответственно контролируют выходы.

На следующем изображении показан пример пользовательского интерфейса Node-RED, который позволяет вам управлять одним выходом и отображать показания температуры и влажности:

![](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2018/05/node-red-ui-output-temperature-humidity.png?resize=517%2C699&quality=100&strip=all&ssl=1)

Вот краткое изложение шагов, которые вы должны выполнить, чтобы построить что-то, как описано ранее:

1.  Настройте свой Raspberry Pi. Следуйте нашему [руководству по началу работы с Raspberry Pi](https://randomnerdtutorials.com/getting-started-with-raspberry-pi/) .
2.  [Включите и подключите Raspberry Pi с помощью SSH](https://randomnerdtutorials.com/installing-raspbian-lite-enabling-and-connecting-with-ssh/) .
3.  Вам нужно [, чтобы Node-RED был установлен на вашем Pi](https://randomnerdtutorials.com/getting-started-with-node-red-on-raspberry-pi/) и [Node-RED Dashboard](https://randomnerdtutorials.com/getting-started-with-node-red-dashboard/) .
4.  Установите [брокера Mosquitto на Raspberry Pi](https://randomnerdtutorials.com/how-to-install-mosquitto-broker-on-raspberry-pi/) .
5.  Добавьте в эту систему ESP8266 или ESP32. Вы можете следовать следующим руководствам по MQTT:

*   [**ESP32** и Node-RED с MQTT — публикация и подписка](https://randomnerdtutorials.com/esp32-mqtt-publish-subscribe-arduino-ide/)
*   [**ESP8266** и Node-RED с MQTT](https://randomnerdtutorials.com/esp8266-and-node-red-with-mqtt/) [—](https://randomnerdtutorials.com/esp32-mqtt-publish-subscribe-arduino-ide/) [публикация и подписка](https://randomnerdtutorials.com/esp8266-and-node-red-with-mqtt/)

Если вы хотите узнать больше об этих предметах, у нас есть специальный курс по созданию собственной системы домашней автоматизации с использованием Raspberry Pi, ESP8266, Arduino и Node-RED. Просто нажмите на следующую ссылку.

[\>> **Зарегистрируйтесь в программе «Создание системы домашней автоматизации» за 100 долларов США <<**](https://randomnerdtutorials.com/build-a-home-automation-system-for-100/)

## Подведение итогов

MQTT — это протокол связи, основанный на системе публикации и подписки. Устройства могут подписываться на тему или публиковать данные по теме. Устройства получают сообщения, опубликованные в темах, на которые они подписаны.

MQTT прост в использовании и отлично подходит для проектов Интернета вещей и домашней автоматизации. Вы можете ознакомиться со всеми нашими [руководствами, связанными с MQTT, здесь](https://randomnerdtutorials.com/?s=mqtt) .

Мы надеемся, что этот учебник был вам полезен, и теперь вы понимаете, что такое MQTT и как он работает.

Спасибо за чтение. Если вам понравилась эта статья, пожалуйста, поддержите нашу работу, [подписавшись на мой блог](https://randomnerdtutorials.com/download) .


![image](https://user-images.githubusercontent.com/3950155/192477476-0b3a5374-8c37-40ed-89e5-9dc42345b97e.png)
## Useful links
* Github driver  github.com/eclipse/paho.mqtt.golang   
* Usage protocol https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/docker/subscriber/main.go   
* Sample usage https://randomnerdtutorials.com/esp8266-nodemcu-websocket-server-arduino/   


