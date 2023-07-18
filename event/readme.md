# Events 
### Send & Subscribe 

**Links**  
github.com/nuttech/bell   
github.com/RichardKnop/machinery   
github.com/go-co-op/gocron    



## Events


```js
 <script>
        $(document).ready(function () {
          var docid = {{.Docid}}
            const first = document.querySelector('#number1');

            first.onchange = function() {
            }

            const evtSource = new EventSource('event/info');
            console.log(evtSource.withCredentials);
            console.log(evtSource.readyState);
            console.log(evtSource.url);
              // $(first).val("sss");

            evtSource.onmessage = function(event) {
                // $(first).val(event.data);
                $("#infosys").html(event.data)
                // console.log('Received SSE event:', event.data);
            };

            evtSource.onopen = function() {
                console.log("Connection to server opened.");
            };

        });

    </script>
```


# GO
```go
package events

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func SseHandler(c echo.Context) error {
	res := c.Response()
	req := c.Request()

	// Set response headers for SSE
	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")

	// Set response headers for CORS (if needed)
	res.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new channel for SSE event data
	events := make(chan string)

	// Notify SSE events
	go notifySSEEvents(events)

	// Close the SSE connection when the client closes the connection
	closed := req.Context().Done()
	go func() {
		<-closed
		close(events)
	}()

	// Write SSE events to the response writer
	for eventData := range events {
		fmt.Fprintf(res, "data: %s\n\n", eventData)
		res.Flush()
	}

	return nil
}

func notifySSEEvents(events chan<- string) {
	for {
		// Simulate generating SSE event data
		eventData := "Im ready " + time.Now().Format(time.RFC3339)

		// Send SSE event data to the channel
		events <- eventData

		// Sleep for a second before sending the next event
		time.Sleep(time.Second)
	}
}
```






