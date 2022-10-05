# Defining and Using a Template Variable in the template.html File in the templates Folder

```go
{{ define "mainTemplate" -}}
 {{ $length := len . }}
 <h1>There are {{ $length }} products in the source data.</h1>
 {{ range getCats . -}}
 <h1>Category: {{ lower . }}</h1>
 {{ end }}
{{- end }}
```

## Defining and Using a Template Variable in the template.html File in the templates Folder

```go
{{ define "mainTemplate" -}}
 <h1>There are {{ len . }} products in the source data.</h1>
 {{- range getCats . -}}
 {{ if ne ($char := slice (lower .) 0 1) "s" }}
 <h1>{{$char}}: {{.}}</h1>
 {{- end }}
 {{- end }}
{{- end }}
```


## Templ
```go
{{ define "mainTemplate" -}}
 {{ range $key, $value := . -}}
 <h1>{{ $key }}: {{ printf "$%.2f" $value.Price }}</h1>
 {{ end }}
{{- end }}
```
