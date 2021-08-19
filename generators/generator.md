# Samples 


```go
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const scriptContent = `
document.addEventListener('DOMContentLoaded', function () {
   var updateButton = document.getElementById("textUpdate");
   updateButton.addEventListener("click", function() {
      var p = document.getElementById("content");
      var message = document.getElementById("textInput").value;
      p.innerHTML = message;
   });
};
`

const htmlContent = `
<html>
   <head>
      <script src="script.js" nonce="{{ . }}"></script>
   </head>
   <body>
       <p id="content"></p>

       <div class="input-group mb-3">
         <input type="text" class="form-control" id="textInput">
         <div class="input-group-append">
           <button class="btn btn-outline-secondary" type="button" id="textUpdate">Update</button>
         </div>
       </div>

       <blockquote class="twitter-tweet" data-lang="en">
         <a href="https://twitter.com/jack/status/20?ref_src=twsrc%5Etfw">March 21, 2006</a>
       </blockquote>
       <script async src="https://platform.twitter.com/widgets.js"
         nonce="{{ . }}" charset="utf-8"></script>
   </body>
</html>
`

func generateNonce() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}

func generateHTML(nonce string) (string, error) {
	var buf bytes.Buffer

	t, err := template.New("htmlContent").Parse(htmlContent)
	if err != nil {
		return "", err
	}

	err = t.Execute(&buf, nonce)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generatePolicy(nonce string) string {
	s := fmt.Sprintf(`'nonce-%v`, nonce) 
	var contentSecurityPolicy = []string{
		`object-src 'none';`,
		fmt.Sprintf(`script-src %v 'strict-dynamic';`, s),
		`base-uri 'none';`,
	}
	return strings.Join(contentSecurityPolicy, " ")
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	nonce, err := generateNonce()
	if err != nil {
		returnError(w)
		return
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Set("Content-Security-Policy", generatePolicy(nonce))

	fmt.Fprintf(w, scriptContent)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	nonce, err := generateNonce()
	if err != nil {
		returnError(w)
		return
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", generatePolicy(nonce))

	htmlContent, err := generateHTML(nonce)
	if err != nil {
      returnError(w)
		return
	}

	fmt.Fprintf(w, htmlContent)
}

func returnError(w http.ResponseWriter) {
         http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}



func NonceHandler(w http.ResponseWriter, r *http.Request) {
       n,_ := generateNonce()
       p := generatePolicy(n)
       fmt.Println("Nonce",n)
       fmt.Println("Police",p)

}



func main() {
	http.HandleFunc("/", NonceHandler)
	http.HandleFunc("/script.js", scriptHandler)
	http.HandleFunc("/test", htmlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
