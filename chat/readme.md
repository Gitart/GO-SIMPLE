channel to signal the client that no more messages will be sent to the client.

The hub handles messages by looping over the registered clients and sending the
message to the client's `send` channel. If the client's `send` buffer is full,
then the hub assumes that the client is dead or stuck. In this case, the hub
unregisters the client and closes the websocket.

### Client

The code for the `Client` type is in [client.go](https://github.com/gorilla/websocket/blob/master/examples/chat/client.go).

The `wsHandler` function is registered by the application's `main` function as
an HTTP handler. The handler upgrades the HTTP connection to the WebSocket
protocol, creates a client, registers the client with the hub and schedules the
client to be unregistered using a defer statement.

Next, the HTTP handler starts the client's `writePump` method as a goroutine.
This method transfers messages from the client's send channel to the websocket
connection. The writer method exits when the channel is closed by the hub or
there's an error writing to the websocket connection.

Finally, the HTTP handler calls the client's `readPump` method. This method
transfers inbound messages from the websocket to the hub.

WebSocket connections [support one concurrent reader and one concurrent
writer](https://godoc.org/github.com/gorilla/websocket#hdr-Concurrency). The
application ensures that these concurrency requirements are met by executing
all reads from the `readPump` goroutine and all writes from the `writePump`
goroutine.

To improve efficiency under high load, the `writePump` function coalesces
pending chat messages in the `send` channel to a single WebSocket message. This
reduces the number of system calls and the amount of data sent over the
network.

## Frontend

The frontend code is in [home.html](https://github.com/gorilla/websocket/blob/master/examples/chat/home.html).

On document load, the script checks for websocket functionality in the browser.
If websocket functionality is available, then the script opens a connection to
the server and registers a callback to handle messages from the server. The
callback appends the message to the chat log using the appendLog function.

To allow the user to manually scroll through the chat log without interruption
from new messages, the `appendLog` function checks the scroll position before
adding new content. If the chat log is scrolled to the bottom, then the
function scrolls new content into view after adding the content. Otherwise, the
scroll position is not changed.

The form handler writes the user input to the websocket and clears the input
field.
