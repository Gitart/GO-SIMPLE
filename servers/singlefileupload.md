## Load single  file

```html
<!doctype html>
<html>
<head>
  <title>File Upload</title>
</head>

<body>
  <form action="/" method="POST" enctype="multipart/form-data">
    <label for="file">File:</label>
    <input type="file" name="file" id="file">
    <br>
    <button type="submit" name="submit">Submit</button>
  </form>
</body>
</html>
```

```golang
func fileForm(w http.ResponseWriter, r *http.Request) {

if r.Method == "GET" {
    t, _ := template.ParseFiles("file.html")
    t.Execute(w, nil)
} else {

f, h, err := r.FormFile("file")

if err != nil {
  panic(err)

}

defer f.Close()
filename := "/tmp/" + h.Filename
out, err := os.Create(filename)

if err != nil {
  panic(err)
}

defer out.Close()
io.Copy(out, f)
fmt.Fprint(w, "Upload complete")
}
}
```


## Multiple load file

```html
<!doctype html>
<html>
<head>
  <title>File Upload</title>
</head>

<body>
  <form action="/" method="POST" enctype="multipart/form-data">
    <label for="files">File:</label>
    <input type="file" name="files" id="files" multiple>
    <br>
    <button type="submit" name="submit">Submit</button>
  </form>
</body>
</html>
```
## go prog

```golang
func fileForm(w http.ResponseWriter, r *http.Request) {

if r.Method == "GET" {
t, _ := template.ParseFiles("file_multiple.html")
t.Execute(w, nil)
} else {
err := r.ParseMultipartForm(16 << 20)

if err != nil {
  fmt.Fprint(w, err)
  return
}

data := r.MultipartForm
files := data.File["files"]
for _, fh := range files {
f, err := fh.Open()
defer f.Close()

if err != nil {
  fmt.Fprint(w, err)
  return
}

out, err := os.Create("/tmp/" + fh.Filename)
defer out.Close()

if err != nil {
  fmt.Fprint(w, err)
  return
}

_, err = io.Copy(out, f)

if err != nil {
  fmt.Fprintln(w, err)
  return
}

}
  fmt.Fprint(w, "Upload complete")
}
}
```
