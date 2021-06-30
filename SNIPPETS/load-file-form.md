## Html form

```html
<html>
<head>
<meta charset="utf-8">
<title>File Upload</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/
form-data">
<label for="file">File:</label>
<input type="file" name="file" id="file">
<input type="submit" name="submit" value="Submit">
</form>
</body>
</html>
```

## Go file
Create upload-file.go, where we will define handlers to render the file upload
template, get the file from the request, process it, and write the response to an
HTTP response stream, as follows:

```go
package main
import
(
"fmt"
"html/template"
"io"
"log"
"net/http"
"os"
)
const
(
CONN_HOST = "localhost"

CONN_PORT = "8080"
)
func fileHandler(w http.ResponseWriter, r *http.Request)
{
file, header, err := r.FormFile("file")
if err != nil
{
log.Printf("error getting a file for the provided form key : ",
err)
return
}
defer file.Close()
out, pathError := os.Create("/tmp/uploadedFile")
if pathError != nil
{
log.Printf("error creating a file for writing : ", pathError)
return
}
defer out.Close()
_, copyFileError := io.Copy(out, file)
if copyFileError != nil
{
log.Printf("error occurred while file copy : ", copyFileError)
}
fmt.Fprintf(w, "File uploaded successfully : "+header.Filename)
}
func index(w http.ResponseWriter, r *http.Request)
{
parsedTemplate, _ := template.ParseFiles("templates/
upload-file.html")
parsedTemplate.Execute(w, nil)
}
func main()
{
http.HandleFunc("/", index)
http.HandleFunc("/upload", fileHandler)
err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
if err != nil
{
log.Fatal("error starting http server : ", err)
return
}
}
```
