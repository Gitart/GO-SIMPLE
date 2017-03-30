# Serving HTTP and Websocket from different ports in a program example


### Problem:
For some reason you want to split the port use by your program to serve HTTP and Web-socket traffic.     
How to write one program that is able to listen and serve HTTP/Web-socket traffics on different ports?   

### Solution:
Setup a http.NewServeMux() and http.ListenAndServe() for one port and launch it as go-routine.    
For example, one go-routine for handling web socket traffic and main routine in the program to handle HTTP traffic.    
NOTE : Instead of web-socket, can you use for another HTTP server? Yes, of course.


```golang
 package main

 import (
         "fmt"
         "golang.org/x/net/websocket"
         "io"
         "net/http"
 )

 func homeHandler(w http.ResponseWriter, r *http.Request) {
         html := `<!DOCTYPE html>
          <html>
          <head>
                  <title>One HTTP and Websocket serve with different ports example</title>
                  <script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
                  <script>
                  $(function() {
                    var ws = new WebSocket("ws://localhost:8082/"); // use port 8082 instead of 8080
                    var $ul = $('#msg-list');

                   $('#sendBtn').click(function(){
                      var data = $('#name').val();
                      ws.send(data);
                      console.log("Sending data to HTTP server via Websocket via port 8082 :" + data);
                     $('<li>').text(data).appendTo($ul);
                    });
                   });
                  </script>
          </head>
          <body>
                  <input id="name" type="text" />
                  <input type="button" id="sendBtn" value="send"></input>
                  <ul id="msg-list"></ul>
          </body></html>`

         w.Write([]byte(fmt.Sprintf(html)))
 }

 func wsHandler(ws *websocket.Conn) {
         fmt.Println("Receive data from : ", ws.LocalAddr())
         fmt.Println("Sending data to : ", ws.RemoteAddr())

         // echo back
         io.Copy(ws, ws)
 }

 func main() {
         mxWS := http.NewServeMux()
         mxWS.Handle("/", websocket.Handler(wsHandler))
         fmt.Println("Listen and serve websocket on :8082")
         go func() {
                 http.ListenAndServe(":8082", mxWS)
         }()

         mxHTTP := http.NewServeMux()
         mxHTTP.HandleFunc("/", homeHandler)
         fmt.Println("Listen and serve HTTP on :8080")
         http.ListenAndServe(":8080", mxHTTP)
 } 
```

Run the code and point your browser to http://localhost:8080. 
Open up the developer view to see the console messages.

References:
https://www.socketloop.com/tutorials/golang-web-javascript-to-server-side-websocket-example
