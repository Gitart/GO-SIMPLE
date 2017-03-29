## Testing HTTP Calls

Another important thing to talk about is how to test HTTP calls. Go is widely used to build web applications and due 
to that, it has got the httptest package.
This package allows you to test both HTTP calls and fake HTTP responses by providing a simple server implementation 
and other utilities such as a ResponseRecorder.

Let’s start by testing an HTTP Client’s behavior. This client is responsible for sending a POST request to a given 
URL with CustomHeader: iLoveBacon in this request’s headers. The method which does that is called BaconPost.

Here is the code for our BaconClient:

```golang
package bacon

import (
    "net/http"
)

// Every BaconClient needs an http.Client in order to make its requests
type BaconClient struct {
    httpClient *http.Client
}

func (client *BaconClient) BaconPost(address string) {
    // Here we're creating a "POST" request and adding a custom header to it
    req, _ := http.NewRequest("POST", address, nil)
    req.Header.Add("CustomHeader", "iLoveBacon")

    // This sends our request
    client.httpClient.Do(req)
}
```

In order to test this, we’re going to create a test server and do assertions on the request object
we get when our client does its HTTP call.

```golang
package bacon

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCustomHeader(t *testing.T) {
    // This is what will be called when the request arrives
    testHandler := func(w http.ResponseWriter, req *http.Request) {
        // We don't even need to answer the request, we just need to assert the request has the data we want
        assert.Equal(t, req.Method, "POST")
        assert.Equal(t, req.Header.Get("CustomHeader"), "iLoveBacon")
    }

    // Here we're effectively creating a server and passing our `testHandler` to it
    testServer := httptest.NewServer(http.HandlerFunc(testHandler))
    defer testServer.Close()

    // Now let's instantiate a client and tell it to do its request to our fake server
    httpClient := &http.Client{}
    client := BaconClient{httpClient}

    // This sends the POST request with our custom header
    client.BaconPost(testServer.URL)
}
```

### This is what happens when running the test above:

We create testHandler, which is the function that will be called by our server when a request arrives.   
We create testServer and pass testHandler to it.    
We create a BaconClient.    
We tell the BaconClient to execute is BaconPost method, which does an HTTP request    
Our testServer receives this request and calls testHandler   
The testHandler function is run, executing the assertions on the req object passed to it    
Testing HTTP servers, however, is a bit different from testing HTTP clients.    
Basically what you’ve gotta do in order to test these servers is test their handlers,     
which are responsible for processing the input and writing to an output.   

To ease this task, the httptest package has got a struct called ResponseRecorder, which is an implementation    
of the http.ResponseWriter that is passed to the handler functions. This other implementation,   
however, records its mutation for later inspection.

To demonstrate this let’s say you’ve got a server which repeats the word sent to it 1 to 10 times and adds    
a header (RepeatHeader) to the response indicating how many times it was repeated.  
This is its code:

```golang
package randmult

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func randomBetween(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func RandomHandler(w http.ResponseWriter, req *http.Request) {
    randomFactor := randomBetween(1, 10)
    w.Header().Add("RepeatHeader", strconv.FormatInt(int64(randomFactor), 10))

    reqBody, _ := ioutil.ReadAll(req.Body)

    fmt.Fprint(w, strings.Repeat(string(reqBody), randomFactor))
}
```

In order to test this we need to create a http.HandlerFunc and then we will be able to pass a   
ResponseRecorder to it by calling ServeHTTP. We then do our assertions using the data on that ResponseRecorder.  

For example:

```golang
package randmult

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)

func TestRandomHandler(t *testing.T) {
    // Let's create a HandlerFunc using our handler as an argument
    handler := http.HandlerFunc(RandomHandler)

    // Here we're creating a io.Reader to be sent as our request's body
    buf := new(bytes.Buffer)
    fmt.Fprint(buf, "word")

    // Create a request object
    req, _ := http.NewRequest("POST", "/", buf)

    // Call our handler function passing our ResponseRecorder and request to it
    recorder := httptest.NewRecorder()
    handler(recorder, req)

    // Let's see if the result matches the multiplier and the word we've sent
    multiplier, _ := strconv.Atoi(recorder.HeaderMap.Get("RepeatHeader"))
    expected := strings.Repeat("word", multiplier)
    response, _ := ioutil.ReadAll(recorder.Body)

    // Now we do assertions on our recorder
    assert.Equal(t, string(response), expected)
}
```

### This is what is happening in the example above:    

We create a HandlerFunc wrapping our RandomHandler     
We create an http.Request object with a word in its body    
We create a ResponseRecorder, which is just like an http.ResponseWriter, but records modifications made to it    
We call our HandlerFunc and pass the ResponseRecorder and the http.Request we have created to it   
We calculate what should be the expected output based on the RepeatHeader and word    
We check the response sent to see if it matches what we expected     
