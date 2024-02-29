## If we want to access both the index and value of our loop:
```go
{{range $key, $value := .Materials }} 
{{ $key }}: {{ $value.Name }}
{{ $key }}: {{ $value.Year }}
//Since books is an array, we create a new loop inside this loop.
  {{range $keybooks, $valuebooks := $value.Books }}
    {{ $keybooks }}: {{ $valuebooks.Author }}
    {{ $keybooks }}: {{ $valuebooks.Title }}
  {{end}}
{{end}}
```


## Part
```go
{{if .Name}} 
  //You can call the data you want to appear here. Example:
  {{ .Name}}
{{else}}
  <p>Dont find data.</p>
{{end}}
```
## T

```go
{{if .Name || .Address}}
{{if or .Name .Address}}
{{ if not ..Address }}
<h1>Not Address in</h1>
{{ end }}
{{with . }}
//If we have library data, it will go here.
 {{end}}
```
## Example
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Name}} - Library</title>
</head>
<body>
    <h1>{{.Name}}</h1>
    <p><strong>Address:</strong> {{.Address}}</p>
    <h2>Materials</h2>
    <ul>
        {{range .Materials}}
        <li>
            <strong>ID:</strong> {{.ID}}, <strong>Name:</strong> {{.Name}}, <strong>Year:</strong> {{.Year}},
            <strong>Category:</strong> {{.Category}}
            <ul>
                {{range .Books}}
                <li>
                    <strong>Book Information:</strong> ISBN: {{.ISBN}}, Title: {{.Title}}, Author: {{.Author}}
                </li>
                {{end}}
                {{range .Magazines}}
                <li>
                    <strong>Magazine Information:</strong> ISSN: {{.ISSN}}, Title: {{.Title}}, Publisher: {{.Publisher}}
                </li>
                {{end}}
            </ul>
        </li>
        {{end}}
    </ul>
</body>
</html>
```
