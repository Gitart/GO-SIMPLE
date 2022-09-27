# ESP8266 NodeMCU WebSocket Server: управляющие выходы (Arduino IDE)

В этом руководстве вы узнаете, как создать веб-сервер с ESP8266, используя протокол связи WebSocket. В качестве примера мы покажем вам, как создать веб-страницу для удаленного управления выходами ESP8266. Состояние вывода отображается на веб-странице и автоматически обновляется во всех клиентах.

![Выходы управления сервером WebSocket ESP8266 NodeMCU Arduino IDE](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2020/10/ESP8266-NodeMCU-WebSocket-Server-Control-Outputs-Arduino-IDE.jpg?resize=828%2C466&quality=100&strip=all&ssl=1)

ESP8266 будет запрограммирован с помощью Arduino IDE и ESPAsyncWebServer. У нас также есть аналогичное [руководство по WebSocket для ESP32](https://randomnerdtutorials.com/esp32-websocket-server-arduino/) .

Если вы следили за некоторыми из наших [предыдущих проектов веб-сервера, таких как этот](https://randomnerdtutorials.com/esp32-async-web-server-espasyncwebserver-library/) , вы могли заметить, что если у вас одновременно открыто несколько вкладок (на одном или разных устройствах), состояние не обновляется во всех вкладки автоматически, если вы не обновите веб-страницу. Чтобы решить эту проблему, мы можем использовать протокол WebSocket — все клиенты могут быть уведомлены, когда происходят изменения, и соответствующим образом обновлять веб-страницу.

Это руководство было основано на проекте, созданном и задокументированном одним из наших читателей (Стефан Кальдерони). Вы можете прочитать его отличный [учебник здесь](https://m1cr0lab-esp32.github.io/remote-control-with-websocket/) .

## Знакомство с веб-сокетом

WebSocket — это постоянное соединение между клиентом и сервером, которое обеспечивает двунаправленную связь между обеими сторонами с использованием TCP-соединения. Это означает, что вы можете отправлять данные с клиента на сервер и с сервера на клиент в любой момент времени.

![Сервер WebSocket ESP32 ESP8266 Как это работает](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2020/10/ESP3-ESP82662-WebSocket-Server-How-it-Works-f.png?resize=727%2C785&quality=100&strip=all&ssl=1)

Клиент устанавливает соединение WebSocket с сервером посредством процесса, известного как *рукопожатие WebSocket* . Рукопожатие начинается с HTTP-запроса/ответа, что позволяет серверам обрабатывать HTTP-соединения, а также соединения WebSocket на одном и том же порту. После установления соединения клиент и сервер могут отправлять данные WebSocket в полнодуплексном режиме.

Используя протокол WebSockets, сервер (плата ESP8266) может отправлять информацию клиенту или всем клиентам без запроса. Это также позволяет нам отправлять информацию в веб-браузер, когда происходит изменение.

Это изменение может происходить на веб-странице (вы нажали кнопку) или на стороне ESP8266, например, при нажатии физической кнопки на схеме.

## Обзор проекта

Вот веб-страница, которую мы создадим для этого проекта.

![ESP32 WebSocket Server Toggle Outputs Обзор проекта](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2020/10/ESP32-WebSocket-Web-Server-Control-Outputs-Web-Page.png?resize=828%2C453&quality=100&strip=all&ssl=1)

*   Веб-сервер ESP8266 отображает веб-страницу с кнопкой для переключения состояния GPIO 2;
*   Для простоты мы управляем GPIO 2 — встроенным светодиодом. Вы можете использовать этот пример для управления любым другим GPIO;
*   Интерфейс показывает текущее состояние GPIO. Всякий раз, когда происходит изменение состояния GPIO, интерфейс обновляется мгновенно;
*   Состояние GPIO автоматически обновляется на всех клиентах. Это означает, что если у вас открыто несколько вкладок веб-браузера на одном устройстве или на разных устройствах, все они обновляются одновременно.

### Как это работает?

На следующем изображении показано, что происходит при нажатии кнопки «Переключить».

![ESP8266 Веб-сервер NodeMCU WebSocket Обновление всех клиентов Как это работает](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2020/10/ESP8266-NodeMCU-WebSocket-Web-Server-Update-All-Clients.png?resize=828%2C752&quality=100&strip=all&ssl=1)

Вот что происходит, когда вы нажимаете кнопку «Переключить»:

1.  Нажмите на кнопку «Переключить»;
2.  Клиент (ваш браузер) отправляет данные по протоколу WebSocket с сообщением «toggle»;
3.  ESP8266 (сервер) получает это сообщение, поэтому он знает, что должен переключать состояние светодиода. Если ранее светодиод был выключен, включите его;
4.  Затем он отправляет данные с новым состоянием светодиода всем клиентам также по протоколу WebSocket;
5.  Клиенты получают сообщение и соответствующим образом обновляют состояние индикатора на веб-странице. Это позволяет нам обновлять всех клиентов почти мгновенно, когда происходит изменение.

## Подготовка Arduino IDE

Мы будем программировать  плату [ESP8266](https://makeradvisor.com/tools/esp8266-esp-12e-nodemcu-wi-fi-development-board/) с помощью Arduino IDE, поэтому убедитесь, что она установлена ​​в вашей Arduino IDE.

*   [Установка платы ESP8266 NodeMCU в Arduino IDE (Windows, Mac OS X, Linux)](https://randomnerdtutorials.com/how-to-install-esp8266-board-arduino-ide/)

### Установка библиотек — асинхронный веб-сервер

Для создания веб-сервера мы будем использовать библиотеку [ESPAsyncWebServer](https://github.com/me-no-dev/ESPAsyncWebServer) . Для правильной работы этой библиотеке требуется библиотека [ESPAsyncTCP](https://github.com/me-no-dev/ESPAsyncTCP) . Щелкните ссылки ниже, чтобы загрузить библиотеки.

*   [ESPAsyncWebServer](https://github.com/me-no-dev/ESPAsyncWebServer/archive/master.zip)
*   [ESPAsyncTCP](https://github.com/me-no-dev/ESPAsyncTCP/archive/master.zip)

Эти библиотеки недоступны для установки через диспетчер библиотек Arduino, поэтому вам необходимо скопировать файлы библиотек в папку установочных библиотек Arduino. В качестве альтернативы в Arduino IDE вы можете перейти в **Sketch** \> **Include Library** > **Add .zip Library**  и выбрать библиотеки, которые вы только что загрузили.

## Код для сервера ESP8266 NodeMCU WebSocket

Скопируйте следующий код в вашу Arduino IDE.

```c
/*********
  Rui Santos
  Complete project details at https://RandomNerdTutorials.com/esp8266-nodemcu-websocket-server-arduino/
  The above copyright notice and this permission notice shall be included in all
  copies or substantial portions of the Software.
*********/

// Import required libraries
#include <ESP8266WiFi.h>
#include <ESPAsyncTCP.h>
#include <ESPAsyncWebServer.h>

// Replace with your network credentials
const char* ssid = "REPLACE_WITH_YOUR_SSID";
const char* password = "REPLACE_WITH_YOUR_PASSWORD";

bool ledState = 0;
const int ledPin = 2;

// Create AsyncWebServer object on port 80
AsyncWebServer server(80);
AsyncWebSocket ws("/ws");

const char index_html[] PROGMEM = R"rawliteral(
<!DOCTYPE HTML><html>
<head>
  <title>ESP Web Server</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
  html {
    font-family: Arial, Helvetica, sans-serif;
    text-align: center;
  }
  h1 {
    font-size: 1.8rem;
    color: white;
  }
  h2{
    font-size: 1.5rem;
    font-weight: bold;
    color: #143642;
  }
  .topnav {
    overflow: hidden;
    background-color: #143642;
  }
  body {
    margin: 0;
  }
  .content {
    padding: 30px;
    max-width: 600px;
    margin: 0 auto;
  }
  .card {
    background-color: #F8F7F9;;
    box-shadow: 2px 2px 12px 1px rgba(140,140,140,.5);
    padding-top:10px;
    padding-bottom:20px;
  }
  .button {
    padding: 15px 50px;
    font-size: 24px;
    text-align: center;
    outline: none;
    color: #fff;
    background-color: #0f8b8d;
    border: none;
    border-radius: 5px;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -khtml-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-tap-highlight-color: rgba(0,0,0,0);
   }
   /*.button:hover {background-color: #0f8b8d}*/
   .button:active {
     background-color: #0f8b8d;
     box-shadow: 2 2px #CDCDCD;
     transform: translateY(2px);
   }
   .state {
     font-size: 1.5rem;
     color:#8c8c8c;
     font-weight: bold;
   }
  </style>
<title>ESP Web Server</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="icon" href="data:,">
</head>
<body>
  <div class="topnav">
    <h1>ESP WebSocket Server</h1>
  </div>
  <div class="content">
    <div class="card">
      <h2>Output - GPIO 2</h2>
      <p class="state">state: <span id="state">%STATE%</span></p>
      <p><button id="button" class="button">Toggle</button></p>
    </div>
  </div>
<script>
  var gateway = `ws://${window.location.hostname}/ws`;
  var websocket;
  window.addEventListener('load', onLoad);
  function initWebSocket() {
    console.log('Trying to open a WebSocket connection...');
    websocket = new WebSocket(gateway);
    websocket.onopen    = onOpen;
    websocket.onclose   = onClose;
    websocket.onmessage = onMessage; // <-- add this line
  }
  function onOpen(event) {
    console.log('Connection opened');
  }
  function onClose(event) {
    console.log('Connection closed');
    setTimeout(initWebSocket, 2000);
  }
  function onMessage(event) {
    var state;
    if (event.data == "1"){
      state = "ON";
    }
    else{
      state = "OFF";
    }
    document.getElementById('state').innerHTML = state;
  }
  function onLoad(event) {
    initWebSocket();
    initButton();
  }
  function initButton() {
    document.getElementById('button').addEventListener('click', toggle);
  }
  function toggle(){
    websocket.send('toggle');
  }
</script>
</body>
</html>
)rawliteral";

void notifyClients() {
  ws.textAll(String(ledState));
}

void handleWebSocketMessage(void *arg, uint8_t *data, size_t len) {
  AwsFrameInfo *info = (AwsFrameInfo*)arg;
  if (info->final && info->index == 0 && info->len == len && info->opcode == WS_TEXT) {
    data[len] = 0;
    if (strcmp((char*)data, "toggle") == 0) {
      ledState = !ledState;
      notifyClients();
    }
  }
}

void onEvent(AsyncWebSocket *server, AsyncWebSocketClient *client, AwsEventType type,
             void *arg, uint8_t *data, size_t len) {
    switch (type) {
      case WS_EVT_CONNECT:
        Serial.printf("WebSocket client #%u connected from %s\n", client->id(), client->remoteIP().toString().c_str());
        break;
      case WS_EVT_DISCONNECT:
        Serial.printf("WebSocket client #%u disconnected\n", client->id());
        break;
      case WS_EVT_DATA:
        handleWebSocketMessage(arg, data, len);
        break;
      case WS_EVT_PONG:
      case WS_EVT_ERROR:
        break;
  }
}

void initWebSocket() {
  ws.onEvent(onEvent);
  server.addHandler(&ws);
}

String processor(const String& var){
  Serial.println(var);
  if(var == "STATE"){
    if (ledState){
      return "ON";
    }
    else{
      return "OFF";
    }
  }
  return String();
}

void setup(){
  // Serial port for debugging purposes
  Serial.begin(115200);

  pinMode(ledPin, OUTPUT);
  digitalWrite(ledPin, LOW);

  // Connect to Wi-Fi
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.println("Connecting to WiFi..");
  }

  // Print ESP Local IP Address
  Serial.println(WiFi.localIP());

  initWebSocket();

  // Route for root / web page
  server.on("/", HTTP_GET, [](AsyncWebServerRequest *request){
    request->send_P(200, "text/html", index_html, processor);
  });

  // Start server
  server.begin();
}

void loop() {
  ws.cleanupClients();
  digitalWrite(ledPin, ledState);
}

```

[Просмотр необработанного кода](https://github.com/RuiSantosdotme/Random-Nerd-Tutorials/raw/master/Projects/ESP8266/ESP8266_WebSocket_Server.ino)

Вставьте свои сетевые учетные данные в следующие переменные, и код сразу заработает.

```c
const char* ssid = "REPLACE_WITH_YOUR_SSID";
const char* password = "REPLACE_WITH_YOUR_PASSWORD";
```

## Как работает код

Продолжайте читать, чтобы узнать, как работает код, или перейдите к разделу « [Демонстрация](https://randomnerdtutorials.com/esp8266-nodemcu-websocket-server-arduino/#1 "#1") ».

### Импорт библиотек

Импортируйте необходимые библиотеки для сборки веб-сервера.

```c
#include <ESP8266WiFi.h>
#include <ESPAsyncTCP.h>
#include <ESPAsyncWebServer.h>
```

### Сетевые учетные данные

Вставьте свои сетевые учетные данные в следующие переменные:

```c
const char* ssid = "REPLACE_WITH_YOUR_SSID";
const char* password = "REPLACE_WITH_YOUR_PASSWORD";
```

### Выход GPIO

Создайте переменную с именем ledState для хранения состояния GPIO и переменной с именем светодиодный контакт который относится к GPIO, которым вы хотите управлять. В этом случае мы будем управлять встроенным светодиодом (который подключен к GPIO 2 ).

```c
bool ledState = 0;
const int ledPin = 2;
```

### Асинквебсервер и асинквебсокет

Создать АсинкВеб-Сервер объект на порту 80.

```c
AsyncWebServer server(80);
```

The ESPAsyncWebServer включает подключаемый модуль WebSocket, упрощающий работу с соединениями WebSocket. Создать Асинквебсокет объект называется ws для обработки соединений на /ws дорожка.

```c
AsyncWebSocket ws("/ws");
```

### Создание веб-страницы

The index\_html переменная содержит HTML, CSS и JavaScript, необходимые для создания и оформления веб-страницы и обработки взаимодействия клиент-сервер с использованием протокола WebSocket.

**Примечание:** мы размещаем все необходимое для создания веб-страницы на index\_html переменная, которую мы используем в скетче Arduino. Обратите внимание, что может быть более практичным иметь отдельные файлы HTML, CSS и JavaScript, которые затем вы загружаете в файловую систему ESP8266 и ссылаетесь на них в коде.

Рекомендуемая литература: [Веб-сервер ESP8266 с использованием SPIFFS (файловая система SPI Flash)](https://randomnerdtutorials.com/esp8266-web-server-spiffs-nodemcu/)

Вот содержание index\_html переменная:

```html
<!DOCTYPE HTML>
<html>
<head>
  <title>ESP Web Server</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="icon" href="data:,">
  <style>
  html {
    font-family: Arial, Helvetica, sans-serif;
    text-align: center;
  }
  h1 {
    font-size: 1.8rem;
    color: white;
  }
  h2{
    font-size: 1.5rem;
    font-weight: bold;
    color: #143642;
  }
  .topnav {
    overflow: hidden;
    background-color: #143642;
  }
  body {
    margin: 0;
  }
  .content {
    padding: 30px;
    max-width: 600px;
    margin: 0 auto;
  }
  .card {
    background-color: #F8F7F9;;
    box-shadow: 2px 2px 12px 1px rgba(140,140,140,.5);
    padding-top:10px;
    padding-bottom:20px;
  }
  .button {
    padding: 15px 50px;
    font-size: 24px;
    text-align: center;
    outline: none;
    color: #fff;
    background-color: #0f8b8d;
    border: none;
    border-radius: 5px;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -khtml-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-tap-highlight-color: rgba(0,0,0,0);
   }
   .button:active {
     background-color: #0f8b8d;
     box-shadow: 2 2px #CDCDCD;
     transform: translateY(2px);
   }
   .state {
     font-size: 1.5rem;
     color:#8c8c8c;
     font-weight: bold;
   }
  </style>
<title>ESP Web Server</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="icon" href="data:,">
</head>
<body>
  <div class="topnav">
    <h1>ESP WebSocket Server</h1>
  </div>
  <div class="content">
    <div class="card">
      <h2>Output - GPIO 2</h2>
      <p class="state">state: <span id="state">%STATE%</span></p>
      <p><button id="button" class="button">Toggle</button></p>
    </div>
  </div>
<script>
  var gateway = `ws://${window.location.hostname}/ws`;
  var websocket;
  function initWebSocket() {
    console.log('Trying to open a WebSocket connection...');
    websocket = new WebSocket(gateway);
    websocket.onopen    = onOpen;
    websocket.onclose   = onClose;
    websocket.onmessage = onMessage; // <-- add this line
  }
  function onOpen(event) {
    console.log('Connection opened');
  }

  function onClose(event) {
    console.log('Connection closed');
    setTimeout(initWebSocket, 2000);
  }
  function onMessage(event) {
    var state;
    if (event.data == "1"){
      state = "ON";
    }
    else{
      state = "OFF";
    }
    document.getElementById('state').innerHTML = state;
  }
  window.addEventListener('load', onLoad);
  function onLoad(event) {
    initWebSocket();
    initButton();
  }

  function initButton() {
    document.getElementById('button').addEventListener('click', toggle);
  }
  function toggle(){
    websocket.send('toggle');
  }
</script>
</body>
</html>
```

### CSS

Между <стиль> </стиль> теги мы включаем стили для оформления веб-страницы с помощью CSS. Не стесняйтесь изменять его, чтобы веб-страница выглядела так, как вы хотите. Мы не будем объяснять, как работает CSS для этой веб-страницы, потому что это не относится к этому руководству по WebSocket.

```html
<style>
  html {
    font-family: Arial, Helvetica, sans-serif;
    text-align: center;
  }
  h1 {
    font-size: 1.8rem;
    color: white;
  }
  h2 {
    font-size: 1.5rem;
    font-weight: bold;
    color: #143642;
  }
  .topnav {
    overflow: hidden;
    background-color: #143642;
  }
  body {
    margin: 0;
  }
  .content {
    padding: 30px;
    max-width: 600px;
    margin: 0 auto;
  }
  .card {
    background-color: #F8F7F9;;
    box-shadow: 2px 2px 12px 1px rgba(140,140,140,.5);
    padding-top:10px;
    padding-bottom:20px;
  }
  .button {
    padding: 15px 50px;
    font-size: 24px;
    text-align: center;
    outline: none;
    color: #fff;
    background-color: #0f8b8d;
    border: none;
    border-radius: 5px;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -khtml-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-tap-highlight-color: rgba(0,0,0,0);
   }
   .button:active {
     background-color: #0f8b8d;
     box-shadow: 2 2px #CDCDCD;
     transform: translateY(2px);
   }
   .state {
     font-size: 1.5rem;
     color:#8c8c8c;
     font-weight: bold;
   }
 </style>
```

### HTML

Между <тело> </тело> теги мы добавляем содержимое веб-страницы, которое видно пользователю.

```html
<div class="topnav">
  <h1>ESP WebSocket Server</h1>
</div>
<div class="content">
  <div class="card">
    <h2>Output - GPIO 2</h2>
    <p class="state">state: <span id="state">%STATE%</span></p>
    <p><button id="button" class="button">Toggle</button></p>
  </div>
</div>
```

Там есть заголовок 1 с текстом «ESP WebSocket Server». Не стесняйтесь изменять этот текст.

```html
<h1>ESP WebSocket Server</h1>
```

Затем идет заголовок 2 с текстом «Выход — GPIO 2».

```c
<h2>Output - GPIO 2</h2>
```

После этого у нас идет абзац, отображающий текущее состояние GPIO.

```html
<p class="state">state: <span id="state">%STATE%</span></p>
```

The %ГОСУДАРСТВО% является заполнителем для состояния GPIO. Он будет заменен текущим значением ESP8266 во время отправки веб-страницы. Заполнители в тексте HTML должны располагаться между % знаки. Это означает, что это %ГОСУДАРСТВО% text похож на переменную, которая затем будет заменена фактическим значением.

После отправки веб-страницы клиенту состояние должно динамически изменяться при каждом изменении состояния GPIO. Мы получим эту информацию по протоколу WebSocket. Затем JavaScript обрабатывает полученную информацию, чтобы соответствующим образом обновить состояние. Чтобы иметь возможность обрабатывать этот текст с помощью JavaScript, текст должен иметь идентификатор, на который мы можем ссылаться. В этом случае идентификатор государство ( <span id="состояние"> ).

Наконец, есть абзац с кнопкой для переключения состояния GPIO.

```html
<p><button id="button" class="button">Toggle</button></p>
```

Обратите внимание, что мы дали идентификатор кнопке ( идентификатор = «кнопка» ).

### JavaScript — обработка веб-сокетов

JavaScript находится между <скрипт> </скрипт> теги. Он отвечает за инициализацию соединения WebSocket с сервером, как только веб-интерфейс полностью загружается в браузере, и за обработку обмена данными через WebSockets.

```html
<script>
  var gateway = `ws://${window.location.hostname}/ws`;
  var websocket;
  function initWebSocket() {
    console.log('Trying to open a WebSocket connection...');
    websocket = new WebSocket(gateway);
    websocket.onopen    = onOpen;
    websocket.onclose   = onClose;
    websocket.onmessage = onMessage; // <-- add this line
  }
  function onOpen(event) {
    console.log('Connection opened');
  }

  function onClose(event) {
    console.log('Connection closed');
    setTimeout(initWebSocket, 2000);
  }
  function onMessage(event) {
    var state;
    if (event.data == "1"){
      state = "ON";
    }
    else{
      state = "OFF";
    }
    document.getElementById('state').innerHTML = state;
  }

  window.addEventListener('load', onLoad);

  function onLoad(event) {
    initWebSocket();
    initButton();
  }

  function initButton() {
    document.getElementById('button').addEventListener('click', toggle);
  }

  function toggle(){
    websocket.send('toggle');
  }
</script>
```

Давайте посмотрим, как это работает.

Шлюз — это точка входа в интерфейс WebSocket.

```javascript
var gateway = `ws://${window.location.hostname}/ws`;
```

окно.местоположение.имя хоста получает текущий адрес страницы (IP-адрес веб-сервера).

Создайте новую глобальную переменную с именем веб-сокет .

```javascript
var websocket;
```

Добавьте прослушиватель событий, который будет вызывать в процессе работать при загрузке веб-страницы.

```javascript
window.addEventListener('load', onload);
```

The в процессе() функция вызывает инитвебсокет() функция для инициализации соединения WebSocket с сервером и кнопка инициализации() функция для добавления прослушивателей событий к кнопкам.

```javascript
function onload(event) {
  initWebSocket();
  initButton();
}
```

The инитвебсокет() Функция инициализирует соединение WebSocket на шлюзе, определенном ранее. Мы также назначаем несколько функций обратного вызова для открытия, закрытия соединения WebSocket или получения сообщения.

```javascript
function initWebSocket() {
  console.log('Trying to open a WebSocket connection…');
  websocket = new WebSocket(gateway);
  websocket.onopen    = onOpen;
  websocket.onclose   = onClose;
  websocket.onmessage = onMessage;
}
```

Когда соединение открыто, мы просто печатаем сообщение в консоли и отправляем сообщение «привет». ESP8266 получает это сообщение, поэтому мы знаем, что соединение было инициализировано.

```javascript
function onOpen(event) {
  console.log('Connection opened');
  websocket.send('hi');
}
```

Если по какой-то причине подключение к веб-сокету закрыто, мы вызываем инитвебсокет() функция снова через 2000 миллисекунд (2 секунды).

```javascript
function onClose(event) {
  console.log('Connection closed');
  setTimeout(initWebSocket, 2000);
}
```

The установить время ожидания () метод вызывает функцию или оценивает выражение через указанное количество миллисекунд.

Наконец, нам нужно обработать то, что происходит, когда мы получаем новое сообщение. Сервер (ваша плата ESP) отправит сообщение «1» или «0». В соответствии с полученным сообщением мы хотим отобразить сообщение «ВКЛ» или «ВЫКЛ» в абзаце, отображающем состояние. Помните, что <диапазон> тег с идентификатор = «состояние» ? Мы получим этот элемент и установим для него значение ON или OFF.

```javascript
function onMessage(event) {
  var state;
  if (event.data == "1"){
    state = "ON";
  }
  else{
    state = "OFF";
  }
  document.getElementById('state').innerHTML = state;
}
```

The кнопка инициализации() функция получает кнопку по ее id ( кнопка ) и добавляет прослушиватель событий типа 'щелкнуть' .

```c
function initButton() {
  document.getElementById('button').addEventListener('click', toggle);
}
```

Это означает, что при нажатии на кнопку переключать вызывается функция.

The переключать функция отправляет сообщение, используя соединение WebSocket с 'переключать' текст.

```c
function toggle(){
  websocket.send('toggle');
}
```

Затем ESP8266 должен обработать то, что происходит, когда он получает это сообщение — переключить текущее состояние GPIO.

### Обработка веб-сокетов — сервер

Ранее вы видели, как обрабатывать соединение WebSocket на стороне клиента (в браузере). Теперь давайте посмотрим, как с этим справиться на стороне сервера.

#### Уведомить всех клиентов

The уведомить клиентов () функция уведомляет всех клиентов сообщением, содержащим все, что вы передаете в качестве аргумента. В этом случае мы хотим уведомлять всех клиентов о текущем состоянии светодиода всякий раз, когда происходит изменение.

```c
void notifyClients() {
  ws.textAll(String(ledState));
}
```

The Асинквебсокет класс предоставляет текстВсе() способ отправки одного и того же сообщения всем клиентам, подключенным к серверу одновременно.

#### Обработка сообщений WebSocket

The обработатьВебсокетмессаже() function — это функция обратного вызова, которая будет выполняться всякий раз, когда мы получаем новые данные от клиентов по протоколу WebSocket.

```c
void handleWebSocketMessage(void *arg, uint8_t *data, size_t len) {
  AwsFrameInfo *info = (AwsFrameInfo*)arg;
  if (info->final && info->index == 0 && info->len == len && info->opcode == WS_TEXT) {
    data[len] = 0;
    if (strcmp((char*)data, "toggle") == 0) {
      ledState = !ledState;
      notifyClients();
    }
  }
}
```

Если мы получаем сообщение «переключить», мы переключаем значение ledState переменная. Кроме того, мы уведомляем всех клиентов по телефону уведомить клиентов () функция. Таким образом, все клиенты уведомляются об изменении и соответствующим образом обновляют интерфейс.

```c
if (strcmp((char*)data, "toggle") == 0) {
  ledState = !ledState;
  notifyClients();
}
```

#### Настройте сервер WebSocket

Теперь нам нужно настроить прослушиватель событий для обработки различных асинхронных шагов протокола WebSocket. Этот обработчик событий можно реализовать, определив по событию() следующим образом:

```c
void onEvent(AsyncWebSocket *server, AsyncWebSocketClient *client, AwsEventType type,
 void *arg, uint8_t *data, size_t len) {
  switch (type) {
    case WS_EVT_CONNECT:
      Serial.printf("WebSocket client #%u connected from %s\n", client->id(), client->remoteIP().toString().c_str());
      break;
    case WS_EVT_DISCONNECT:
      Serial.printf("WebSocket client #%u disconnected\n", client->id());
      break;
    case WS_EVT_DATA:
      handleWebSocketMessage(arg, data, len);
      break;
    case WS_EVT_PONG:
    case WS_EVT_ERROR:
      break;
  }
}
```

The тип аргумент представляет происходящее событие. Он может принимать следующие значения:

*   WS\_EVT\_CONNECT когда клиент вошел в систему;
*   WS\_EVT\_DISCONNECT когда клиент вышел из системы;
*   WS\_EVT\_DATA когда пакет данных получен от клиента;
*   WS\_EVT\_PONG в ответ на пинг-запрос;
*   WS\_EVT\_ERROR при получении сообщения об ошибке от клиента.

#### Инициализировать веб-сокет

Наконец, инитвебсокет() функция инициализирует протокол WebSocket.

```c
void initWebSocket() {
  ws.onEvent(onEvent);
  server.addHandler(&ws);
}
```

### процессор()

The процессор() Функция отвечает за поиск заполнителей в тексте HTML и замену их тем, что мы хотим, перед отправкой веб-страницы в браузер. В нашем случае мы заменим %ГОСУДАРСТВО% заполнитель с НА если ledState является 1 . В противном случае замените его на ВЫКЛЮЧЕННЫЙ .

```c
String processor(const String& var){
  Serial.println(var);
  if(var == "STATE"){
    if (ledState){
      return "ON";
    }
    else{
      return "OFF";
    }
  }
}
```

### настраивать()

в настраивать() , инициализируйте Serial Monitor для целей отладки.

```c
Serial.begin(115200);
```

Настройте светодиодный контакт как ВЫХОД и установите его на НИЗКИЙ при первом запуске программы.

```c
pinMode(ledPin, OUTPUT);
digitalWrite(ledPin, LOW);
```

Инициализируйте Wi-Fi и напечатайте IP-адрес ESP8266 на последовательном мониторе.

```c
WiFi.begin(ssid, password);
while (WiFi.status() != WL_CONNECTED) {
  delay(1000);
  Serial.println("Connecting to WiFi..");
}

// Print ESP Local IP Address
Serial.println(WiFi.localIP());
```

Инициализируйте протокол WebSocket, вызвав метод инитвебсокет() функция, созданная ранее.

```c
initWebSocket();
```

### Обработка запросов

Подайте текст, сохраненный на index\_html переменная при получении запроса по корню **/** URL — нужно передать процессор в качестве аргумента для замены заполнителей текущим состоянием GPIO.

```c
server.on("/", HTTP_GET, [](AsyncWebServerRequest *request){
  request->send_P(200, "text/html", index_html, processor);
});
```

Наконец, запустите сервер.

```c
server.begin();
```

### петля()

Светодиод будет физически управляться на петля() .

```c
void loop() {
  ws.cleanupClients();
  digitalWrite(ledPin, ledState);
}
```

Обратите внимание, что все мы называем очисткаклиентов() метод. Вот почему (пояснение со страницы GitHub библиотеки ESPAsyncWebServer):

Браузеры иногда неправильно закрывают соединение WebSocket, даже если Закрыть() функция вызывается в JavaScript. Это в конечном итоге истощит ресурсы веб-сервера и приведет к сбою сервера. Периодически звоню в очисткаклиентов() функция от основного петля() ограничивает количество клиентов, закрывая самый старый клиент, когда превышено максимальное количество клиентов. Это может быть вызвано каждый цикл, однако, если вы хотите использовать меньше энергии, достаточно вызывать так редко, как один раз в секунду.

## Демонстрация

После ввода ваших сетевых учетных данных на ssid а также пароль переменные, вы можете загрузить код на свою доску. Не забудьте проверить, правильно ли выбрана плата и COM-порт.

После загрузки кода откройте Serial Monitor со скоростью 115200 бод и нажмите встроенную кнопку EN/RST. IP-адрес ESP должен быть напечатан.

Откройте браузер в локальной сети и введите IP-адрес ESP8266. Вы должны получить доступ к веб-странице для управления выводом.

![ESP32 WebSocket Server Toggle Outputs Обзор проекта](https://i0.wp.com/randomnerdtutorials.com/wp-content/uploads/2020/10/ESP32-WebSocket-Web-Server-Control-Outputs-Web-Page.png?resize=828%2C453&quality=100&strip=all&ssl=1)

Нажмите на кнопку, чтобы переключить светодиод. Вы можете открыть несколько вкладок веб-браузера одновременно или получить доступ к веб-серверу с разных устройств, и состояние индикатора будет автоматически обновляться на всех клиентах при каждом изменении.

## Подведение итогов

В этом руководстве вы узнали, как настроить сервер WebSocket с ESP8266. Протокол WebSocket обеспечивает полнодуплексную связь между клиентом и сервером. После инициализации сервер и клиент могут обмениваться данными в любой момент времени.

Это очень полезно, потому что сервер может отправлять данные клиенту всякий раз, когда что-то происходит. Например, вы можете добавить к этой настройке [физическую кнопку](https://randomnerdtutorials.com/esp32-esp8266-web-server-physical-button/) , которая при нажатии уведомляет всех клиентов о необходимости обновления веб-интерфейса.

В этом примере мы показали вам, как управлять одним GPIO ESP8266. Вы можете использовать этот метод для управления большим количеством GPIO. Вы также можете использовать протокол WebSocket для отправки показаний датчиков или уведомлений в любой момент времени.

Мы надеемся, что вы нашли этот урок полезным. Мы намерены создать больше руководств и примеров с использованием протокола WebSocket. Итак, следите за обновлениями.
