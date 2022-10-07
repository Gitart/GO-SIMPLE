# Uses gorilla sessions

```go
import (
    "github.com/gorilla/sessions"
    "net/http"
)

// Authorization Key
var authKey = []byte("somesecret")

// Encryption Key
var encKey = []byte("someothersecret")

var store = sessions.NewCookieStore(authKey, encKey)

func initSession(r *http.Request) *sessions.Session {
    session, _ := store.Get(r, "my_cookie") // Don't ignore the error in real code
    if session.IsNew { //Set some cookie options
        session.Options.Domain = "example.org"
        session.Options.MaxAge = 0
        session.Options.HttpOnly = false
        session.Options.Secure = true
    }
    return session
}
```

## Then, in your handlers:

```
func ViewPageHandler(w http.ResponseWriter, r *http.Request) {
    session := initSession(r)
    session.Values["page"] = "view"
    session.Save(r, w)
....
```
