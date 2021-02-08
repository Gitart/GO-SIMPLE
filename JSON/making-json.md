# Making HTTP requests in Go

September 14, 2020 8 min read 2435

![Making HTTP requests in Go](https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2020/09/makinghttprequestsingo.png?fit=730%2C412&ssl=1)

HTTP requests are a very fundamental part of the web as a whole. They are used to access resources hosted on a server (which could be remote).

HTTP is an acronym for hypertext transfer protocol, a communication protocol that ensures the transfer of data between a client and a server. A perfect instance of an HTTP client\-server interaction is when you open your browser and type in a URL. Your browser acts as a client and fetches resources from a server which it then displays.

In web development, cases where we need to fetch resources, are very common. You might be making a weather application and need to fetch the weather data from an API. In such a case, using your browser as a client would no longer be possible from within your application. So you have to set up an HTTP client within your application to handle the making of these requests.

Most programming languages have various structures in place for setting up HTTP clients for making requests. In the coming sections, we will take a hands\-on approach in exploring how you can make HTTP requests in Golang or Go, as I will refer to the language for the rest of the article.

## Prerequisites

To follow this article you will need:

*   [Go (version 1.14 or higher)](https://golang.org/dl/)
*   A text editor of your choice
*   Basic knowledge of Go

## Making HTTP requests in Go

### GET request

The first request we will be making is a GET request. The HTTP GET method is used for requesting data from a specified source or server. The GET method is mostly used when data needs to be fetched.

For the sake of clarity, it is important to note that the HTTP methods, as seen in this article, are always capitalized.

For our example, we will be fetching some example JSON data from [https://jsonplaceholder.typicode.com/posts](https://jsonplaceholder.typicode.com/posts) using the GET method.

The first step in making an HTTP request with Go is to import the `net/http` package from the standard library. This package provides us with all the utilities we need to make HTTP requests with ease. We can import the `net/http` package and other packages we will need by adding the following lines of code to a `main.go` file that we create:

import  (  "io/ioutil"  "log"  "net/http"  )

The `net/http` package we imported has a Get function used for making GET requests. The Get function takes in a URL and returns a response of type pointer to a struct and an error. When the error is `nil`, the response returned will contain a response body and vice versa:

[

## We made a custom demo for .
No really. Click here to check it out.

![](https://blog.logrocket.com/making-http-requests-in-go/)

![](https://blog.logrocket.com/making-http-requests-in-go/)

Click here to see the full demo with network requests

](https://blog.logrocket.com/making-http-requests-in-go/)

resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")  if err !=  nil  { log.Fatalln(err)  }

To make the request, we invoke the Get function, passing in a URL string ([https://jsonplaceholder.typicode.com/posts](https://jsonplaceholder.typicode.com/posts)) as seen above. The values returned from the invocation of this function are stored in two variables typically called resp and err. Although the variable resp contains our response, if we print it out we would get a load of incoherent data which includes the header and properties of the request made. To get the response we are interested in, we have to access the `Body` property on the response struct and read it before finally printing it out to the terminal. We can read the response body using the `ioutil.ReadMe` function.

Similar to the `Get` function, the `ioutil.ReadMe` function returns a body and an error. It is important to note that the response `Body` should be closed after we are done reading from it to prevent memory leaks.

The defer keyword which executes `resp.Body.Close()` at the end of the function is used to close the response body. We can then go ahead and print out the value of the response to the terminal. As good programmers it is important to handle possible errors, so we use an if statement to check for any errors and log the error if it exists:

package main import  (  "io/ioutil"  "log"  "net/http"  ) func main()  { resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")  if err !=  nil  { log.Fatalln(err)  }  //We Read the response body on the line below. body, err := ioutil.ReadAll(resp.Body)  if err !=  nil  { log.Fatalln(err)  }  //Convert the body to type string sb :=  string(body) log.Printf(sb)  }

At this point, we are all set and can execute the file containing our code. If everything went well you will notice some JSON data similar to the image below gets printed to the terminal:

![JSON data printed in terminal including user ID, id, title, and body](https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2020/09/jsondatainterminal.png?resize=680%2C233&ssl=1)

![JSON data printed in terminal including user ID, id, title, and body](https://i0.wp.com/blog.logrocket.com/wp-content/uploads/2020/09/jsondatainterminal.png?resize=680%2C233&ssl=1)

Congratulations, you have just made your first HTTP request with Go. Now that we have seen how we can fetch resources from a server using the HTTP GET method, we will look at how to post resources to a server next.

### POST request

The HTTP POST method is used to make requests that usually contain a body. It Is used to send data to a server, the data sent is usually used for creating or updating resources.

A clear instance where a POST request is used is when a user tries to create a social media account, the user is required to provide their data (name, email, and password). This data is then parsed and sent as a POST request to a server which then creates and saves the user. Just like for the GET method seen above, Go’s `net/http` package also provides functionality for making POST requests through the Post function. The Post function takes three parameters.

1.  The URL address of the server
2.  The content type of the body as a string
3.  The request body that is to be sent using the POST method of type `io.Reader`

The Post function returns a response and an error. For us to invoke the Post function we have to convert our request body to the accepted type. For this example, we will make a post request to [https://postman\-echo.com/post](https://postman-echo.com/post) and pass in JSON data containing a name and an email. To get started we convert our JSON data to a type that implements the Io.Reader interface the Post function expects, this is a two\-way step:

*   The first step is to encode our JSON data so it can return data in byte format, to do this we use the [Marshall function](https://golang.org/pkg/encoding/json/#Marshal) Go’s Json package provides
*   Next, we convert the encoded JSON data to a type implemented by the `io.Reader` interface, we simply use the `NewBuffer` function for this, passing in the encoded JSON data as an argument. The `NewBuffer` function returns a value of type buffer which we can then pass unto the Post function

postBody, \_ := json.Marshal(map\[string\]string{  "name":  "Toby",  "email":  "Toby@example.com",  }) responseBody := bytes.NewBuffer(postBody)

Now that we have all the arguments the Post function requires, we can go ahead and invoke it, passing in [https://postman\-echo.com/post](https://postman-echo.com/post) as the URL string, application/JSON as the content type, and the request body returned by the `NewBuffer` function as the body. The values returned by the `Post` function is then assigned to resp and err representing the response and error, respectively. After handling the error, we read and print in the response body as we did for the Get function in the previous section. At this point, your file should look like this:

import  (  "bytes"  "encoding/json"  "io/ioutil"  "log"  "net/http"  ) func main()  {  //Encode the data postBody, \_ := json.Marshal(map\[string\]string{  "name":  "Toby",  "email":  "Toby@example.com",  }) responseBody := bytes.NewBuffer(postBody)  //Leverage Go's HTTP Post function to make request resp, err := http.Post("https://postman\-echo.com/post",  "application/json", responseBody)  //Handle Error  if err !=  nil  { log.Fatalf("An Error Occured %v", err)  } defer resp.Body.Close()  //Read the response body body, err := ioutil.ReadAll(resp.Body)  if err !=  nil  { log.Fatalln(err)  } sb :=  string(body) log.Printf(sb)  }

When the file is executed, if everything works well we should have the response printed out. Amazing, right? We just made a post request with Go using the `net/http` package which provides functionality that makes HTTP requests easier. In the next section, we will work on a project, to help us see HTTP requests being used in a real\-life scenario.

## HTTP requests in action

In this section, we will be building a cryptocurrency price checker CLI tool! This exercise aims to enable you to see a real\-life use case of HTTP requests. The tool we are building will check the price of whatever cryptocurrency as specified by the user in the specified fiat currency. We will use the crypto market cap and pricing data provided by Nomics to get the price of the cryptocurrencies in real\-time! To get started, create the needed files and folders to match the tree structure below:

├── model/  │  ├── crypto\-model.go ├── client/  │  ├── crypto\-client.go └── main.go

*   The crypto\-client file will house the code that fetches the cryptocurrency data from the API
*   The crypto\-model file houses a couple of utility functions necessary for our application
*   The main file is the central engine of the application, it will merge all the parts of the application to make it functional

In the crypto\-model file, we create a struct that models the data received from the API, this struct includes only the specific data we need/intend to work with. Next, we create a function called `TextOutput` which is a receiver that belongs to the `Cryptoresponse` struct we created up above. The purpose of the `TextOutput` function is to format the data gotten from the API to plain text which is easier to read than JSON(which we receive from the server). We use the `fmt.Sprintf` function to format the data:

package model import  (  "fmt"  )  // Cryptoresponse is exported, it models the data we receive. type Cryptoresponse  \[\]struct  {  Name  string  \`json:"name"\`  Price  string  \`json:"price"\`  Rank  string  \`json:"rank"\`  High  string  \`json:"high"\`  CirculatingSupply  string  \`json:"circulating\_supply"\`  }  //TextOutput is exported,it formats the data to plain text. func (c Cryptoresponse)  TextOutput()  string  { p := fmt.Sprintf(  "Name: %s\\nPrice : %s\\nRank: %s\\nHigh: %s\\nCirculatingSupply: %s\\n", c\[0\].Name, c\[0\].Price, c\[0\].Rank, c\[0\].High, c\[0\].CirculatingSupply)  return p }

Now that the `crypto-model` file is ready, we can move on to the `crypto-client` file, which is the most relevant to us. In the `crypto-client` file, we create a `FetchCrypto` function that takes in the name of the cryptocurrency and fiat currency as parameters.

> Note that we capitalize the first letter of the function name, this is to ensure it is exported.

In the `FetchCrypto` function, we create a variable called URL, the variable is a concatenation of the URL string provided by the [Nomics API](https://api.nomics.com/v1/) and the various variables that will be passed into our application. Remember our application takes in the name of the desired cryptocurrency and the preferred fiat currency? These are the variables that are then used to build our URL string. Our URL string would look like this.

URL :=  "...currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids="+crypto+"&convert="+fiat

After setting up the URL, we can go ahead and use the Get function we saw up above to make a request. The Get function returns the response and we handle the error elegantly. To get the data we want, in the format we want, we have to decode it! To do so, we use the `Json.NewDecoder` function that takes in the response body and a decode function which takes in a variable of type cryptoresponse which we created in the `crypto-model` file. Lastly, we invoke the `TextOutput` function, on the decoded data to enable us to get our result in plain text:

package client import  (  "encoding/json"  "fmt"  "log"  "net/http"  "github.com/Path/to/model"  )  //Fetch is exported ... func FetchCrypto(fiat string  , crypto string)  (string, error)  {  //Build The URL string URL :=  "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids="+crypto+"&convert="+fiat //We make HTTP request using the Get function resp, err := http.Get(URL)  if err !=  nil  { log.Fatal("ooopsss an error occurred, please try again")  } defer resp.Body.Close()  //Create a variable of the same type as our model  var cResp model.Cryptoresponse  //Decode the data  if err := json.NewDecoder(resp.Body).Decode(&cResp); err !=  nil  { log.Fatal("ooopsss! an error occurred, please try again")  }  //Invoke the text output function & return it with nil as the error value  return cResp.TextOutput(),  nil  }

From what we have above, the application is coming together nicely. However, if you try to run the file above, you will encounter a couple of errors, this is because we are not invoking the `FetchCrypto` function and so the value of the fiat and crypto parameters are not provided. We will put all the various parts of our application together in the `main.go` file we created. Since our application is a command\-line tool, users will have to pass in data through the terminal, We will handle that using Go’s flag package.

In the main function, we create two variables `fiatcurrency` and `nameofcrypto`. These variables both invoke the `flag.string` function, passing in:

*   The name of the commands as the first argument
*   The fallback values as the second
*   The information on how to use the command as the third argument

Next, we invoke the `FetchCrypto` function we defined in the `crypto-client` file and pass in the `fiatcurrency` and `nameofcrypto` variables. We can then go ahead and print the result of the call to `FetchCrypto`:

package main import  (  "flag"  "fmt"  "log"  "github.com/path/to/client"  ) func main()  { fiatCurrency := flag.String(  "fiat",  "USD",  "The name of the fiat currency you would like to know the price of your crypto in",  ) nameOfCrypto := flag.String(  "crypto",  "BTC",  "Input the name of the CryptoCurrency you would like to know the price of",  ) flag.Parse() crypto, err := client.FetchCrypto(\*fiatCurrency,  \*nameOfCrypto)  if err !=  nil  { log.Println(err)  } fmt.Println(crypto)  }

At this point, we are good to go, if we run the command `go run main.go -fiat=EUR -crypto=ETH` we would get an output similar to the image below:

![Name: Ethereum, Price: 362.34252819, Rank: 2, High: 1146.26224974, CirculatingSupply: 112196536](https://i2.wp.com/blog.logrocket.com/wp-content/uploads/2020/09/output1.png?resize=680%2C170&ssl=1)

![Name: Ethereum, Price: 362.34252819, Rank: 2, High: 1146.26224974, CirculatingSupply: 112196536](https://i2.wp.com/blog.logrocket.com/wp-content/uploads/2020/09/output1.png?resize=680%2C170&ssl=1)

This shows our application is working fine which is pretty awesome. We have an application that fetches data from a remote server using the HTTP protocol.

## Conclusion

In this article, we discussed how to make HTTP requests in Go, and we built a CLI tool for checking the prices of cryptocurrencies. I highly recommend checking out the [source code](https://golang.org/src/net/http/client.go?s=16191:16251#L460) and [documentation](https://golang.org/pkg/net/http/) of the `net/http` package to explore the other amazing functionalities it provides.

## [LogRocket](https://logrocket.com/signup/): Full visibility into your web apps

[![LogRocket Dashboard Free Trial Banner](https://i1.wp.com/blog.logrocket.com/wp-content/uploads/2017/03/1d0cd-1s_rmyo6nbrasp-xtvbaxfg.png?resize=1200%2C677&ssl=1)

![LogRocket Dashboard Free Trial Banner](https://i1.wp.com/blog.logrocket.com/wp-content/uploads/2017/03/1d0cd-1s_rmyo6nbrasp-xtvbaxfg.png?resize=1200%2C677&ssl=1)

](https://logrocket.com/signup/)

[LogRocket](https://logrocket.com/signup/) is a frontend application monitoring solution that lets you replay problems as if they happened in your own browser. Instead of guessing why errors happen, or asking users for screenshots and log dumps, LogRocket lets you replay the session to quickly understand what went wrong. It works perfectly with any app, regardless of framework, and has plugins to log additional context from Redux, Vuex, and @ngrx/store.

In addition to logging Redux actions and state, LogRocket records console logs, JavaScript errors, stacktraces, network requests/responses with headers + bodies, browser metadata, and custom logs. It also instruments the DOM to record the HTML and CSS on the page, recreating pixel\-perfect videos of even the most complex single\-page apps.
