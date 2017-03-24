## Загрузка файлов с помощью клиента
Загружать файлы через форму без участия человека:

```golang
package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
)

func postFile(filename string, targetUrl string) error {
    bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    // этот шаг очень важен
    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
    if err != nil {
        fmt.Println("ошибка записи в буфер")
        return err
    }

    // процедура открытия файла
    fh, err := os.Open(filename)
    if err != nil {
        fmt.Println("ошибка открытия файла")
        return err
    }

    //iocopy
    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        return err
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, err := http.Post(targetUrl, contentType, bodyBuf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    fmt.Println(resp.Status)
    fmt.Println(string(resp_body))
    return nil
}

// пример использования
func main() {
    target_url := "http://localhost:9090/upload"
    filename := "./astaxie.pdf"
    postFile(filename, target_url)
}
```

Этот пример показывает, как можно использовать клиента для загрузки файлов.     
Он использует multipart.Write для того, чтобы записывать файлы в кэш, и посылает их на сервер посредством метода POST.
Если у Вас есть другие поля, которые нужно писать в данные, такие, как имя пользователя,    
вызывайте по необходимости метод multipart.WriteField.
