# Upload file from web browser to server

Recently, I came across a question on Go programming language Facebook 
group on how to upload file with http://golang.org/*
In this tutorial, we will use Go and HTML to achieve the following :
A simple HTML file upload form to upload file to server
Go program on server side to receive the uploaded file
Verify the uploaded file location and content on server.
1) HTML file upload form to upload file to server -   


goupload.html

```html
  <html>
  <title>Go upload</title>
  <body>

  <form action="http://localhost.com:8080/receive" method="post" enctype="multipart/form-data">
  <label for="file">Filename:</label>
  <input type="file" name="file" id="file">
  <input type="submit" name="submit" value="Submit">
  </form>

  </body>
  </html>
```

NOTE : The input id file will be captured by the uploadHandler's FormFile function.


2) Go program on server side to receive the uploaded file - receive.go

```golang
 package main

 import (
 	"fmt"
 	"io"
 	"net/http"
 	"os"
 )

 func uploadHandler(w http.ResponseWriter, r *http.Request) {

 	// the FormFile function takes in the POST input id file
 	file, header, err := r.FormFile("file")

 	if err != nil {
 		fmt.Fprintln(w, err)
 		return
 	}

 	defer file.Close()

 	out, err := os.Create("/tmp/uploadedfile")
 	if err != nil {
 		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
 		return
 	}

 	defer out.Close()

 	// write the content from POST to the file
 	_, err = io.Copy(out, file)
 	if err != nil {
 		fmt.Fprintln(w, err)
 	}

 	fmt.Fprintf(w, "File uploaded successfully : ")
 	fmt.Fprintf(w, header.Filename)
 }

 func main() {
 	http.HandleFunc("/", uploadHandler)
 	http.ListenAndServe(":8080", nil)
 }
 ```
 
the code above uses the FormFile function to process the POST array and focus on file input

run receive.go at the server 
```
> go run receive.go
```


browse goupload.html and upload a file of your choice. The uploaded file will be named uploadedfile

3) Verify the uploaded file location and content on the server.
See if the file is uploaded to the /tmp with ls command and cat to see the content
[Visit](https://www.socketloop.com/tutorials/golang-how-to-verify-uploaded-file-is-image-or-allowed-file-types)
for further information on how to verify or detect the uploaded file type.
