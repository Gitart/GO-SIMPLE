# Securing Your Go REST APIs With JWTs


> **Note** - The full source code for this tutorial can be found here: [TutorialEdge/go-jwt-tutorial](https://github.com/TutorialEdge/go-jwt-tutorial "TutorialEdge/go-jwt-tutorial")

JWTs, or JSON Web Tokens as they are more formally known, are a compact, URL-safe means of representing claims to be transferred between two parties. This is essentially a confusing way of saying that JWTs allow you to transmit information from a client to the server in a stateless, but secure way.

## Prerequisites

Before you can follow this article, you will need the following:

*   You will need Go version 1.11+ installed on your development machine.

## Introduction

The JWT standard uses either a secret, using the HMAC algorithm, or a public/private key pair using RSA or ECDSA.

> **Note -** If you are interested in the formal definition of what JWTs are, then I recommend checking out the RFC: [RFC-7519](https://tools.ietf.org/html/rfc7519 "RFC-7519")

These are heavily used within Single-Page Applications (SPAs) as a means of secure communications as they allow us to do two key things:

*   **Authentication** - The most commonly used practice. Once a user logs in to your application, or authenticates in some manner, every request that is then sent from the client on behalf of the user will contain the JWT.
*   **Information Exchange** - The second use for JWTs is to securely transmit information between different systems. These JWTs can be signed using public/private key pairs so you can verify each system in this transaction in a secure manner and JWTs contain an anti-tamper mechanism as they are signed based off the header and the payload.

So, if you haven’t guessed by now, in this tutorial, we’ll be looking at exactly what it takes to build a secure Go-based REST API that uses JSON Web Tokens to communicate!

## Video Tutorial

This tutorial is available in a video format, if you want to support me and my work then please feel free to leave a like and subscribe to my channel for my content!

## A Simple REST API

So, we are going to be using the code from one of my other articles, [Creating a simple REST API in Go](https://tutorialedge.net/golang/creating-restful-api-with-golang/ "Creating a simple REST API in Go"), to get us started. This will feature a really simple `Hello World` style endpoint and it will run on port 8081.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}

```

When we run this an attempt to hit our homepage, running on `http://localhost:8081/`, we should see the message `Hello World` in our browser.

## JWT Authentication

So, now that we have a simple API that we can now protect using signed JWT tokens, let’s build a client API that will try to request data from this original API.

To do this, we can use a JWT that has been signed with a secure key that both our client and server will have knowledge off. Let’s walk through how this will work:

1.  Our client will generate a signed JWT based of our shared passphrase.
2.  When our client goes to hit our server API, it will include this JWT as part of the request.
3.  Our server will be able to read this JWT and validate the token using the same passphrase.
4.  If the JWT is valid, it will then return the highly confidential `hello world` message back to the client, otherwise it’ll return `not authorized`.

Our architecture diagram is going to end up looking a little something like this:

![architecture-diagram](https://images.tutorialedge.net/images/golang/go-jwt-tutorial/diagram-01.png)

### Our Server

So, let’s see this in action, let’s create a really simple server:

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
    fmt.Println("Endpoint Hit: homePage")

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("There was an error")
                }
                return mySigningKey, nil
            })

            if err != nil {
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                endpoint(w, r)
            }
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}

func handleRequests() {
    http.Handle("/", isAuthorized(homePage))
    log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
    handleRequests()
}

```

Let’s break this down. We’ve created a really simple API that feature a solitary endpoint, this is protected by our `isAuthorized` middleware decorator. In this `isAuthorized` function, we check to see that the incoming request features the `Token` header in the request and we then subsequently check to see if the token is valid based off our private `mySigningKey`.

If this is a valid token, we then serve the protected endpoint.

> **Note -** This example uses decorators, if you aren’t comfortable with the concept of decorators in Go then I recommend you check out my other article here: [Getting Started with Decorators in Go](https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/ "Getting Started with Decorators in Go")

### Our Client

Now that we have a server that features a JWT secured endpoint, let’s build something that can interact with it.

We’ll be building a simple client application which will attempt to call our `/` endpoint of our server.

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
    validToken, err := GenerateJWT()
    if err != nil {
        fmt.Println("Failed to generate token")
    }

    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
    req.Header.Set("Token", validToken)
    res, err := client.Do(req)
    if err != nil {
        fmt.Fprintf(w, "Error: %s", err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, string(body))
}

func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = "Elliot Forbes"
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

func handleRequests() {
    http.HandleFunc("/", homePage)

    log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
    handleRequests()
}

```

Let’s break down what’s happening in the above code. Again, we’ve defined a really simple API that features a single endpoint. This endpoint, when triggered, generates a new JWT using our secure `mySigningKey`, it then creates a new http client and sets the `Token` header equal to the JWT string that we have just generated.

It then attempts to hit our `server` application which is running on `http://localhost:9000` using this signed JWT token. Our server then validates the token we’ve generated in the client and proceeds to serve us our super secret `Hello World` message.

## Conclusion

Hopefully, this tutorial helped to demystify the art of securing your Go applications and REST APIs using JSON Web Tokens. This was a lot of fun writing this article, and I hope it has helped you in your Go development travels.

If you enjoyed this tutorial, then please let me know in the comments section below, or give this article a share across social media, it really helps me and my site!

> **Note -** If you want to keep track of when new Go articles are posted to the site, then please feel free to follow me on twitter for all the latest news: [@Elliot\_F](https://twitter.com/elliot_f "@Elliot_F").

### Further Reading

If you fancy reading up more on JSON Web Tokens and how they are used then I can thoroughly recommend the following articles:

*   **Go Decorators Tutorial** - [Go Decorators](https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/ "Go Decorators")
*   **RFC-7519** - [https://tools.ietf.org/html/rfc7519](https://tools.ietf.org/html/rfc7519 "https://tools.ietf.org/html/rfc7519")
*   **JWT Introduction** - [https://jwt.io/introduction/](https://jwt.io/introduction/ "https://jwt.io/introduction/")
