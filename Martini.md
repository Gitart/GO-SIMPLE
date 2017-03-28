## Пример работы

```golang
package main

import (
    "log"
    "net/http"

    "github.com/go-martini/martini"
    "github.com/martini-contrib/secure"
)

func main() {
    martini.Env = martini.Prod

    m := martini.New()
    m.Use(martini.Logger())
    m.Use(martini.Recovery())
    m.Use(martini.Static("public"))

    r := martini.NewRouter()
    m.MapTo(r, (*martini.Routes)(nil))
    m.Action(r.Handle)

    r.Get("/", func() string {
        return "Hello world!"
    })

    m.Use(secure.Secure(secure.Options{
    SSLRedirect:  true,
    SSLHost:      "localhost:8443",  // This is optional in production. The default behavior is to just redirect the request to the https protocol. Example: http://github.com/some_page would be redirected to https://github.com/some_page.
    }))

    // HTTP
    go func() {
        if err := http.ListenAndServe(":8080", m); err != nil {
            log.Fatal(err)
        }
    }()

    // HTTPS
    // To generate a development cert and key, run the following from your *nix terminal:
    // go run $GOROOT/src/pkg/crypto/tls/generate_cert.go --host="localhost"
    if err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", m); err != nil {
        log.Fatal(err)
    }
}
```
