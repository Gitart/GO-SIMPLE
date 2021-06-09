## Цели

К концу этого урока вы сможете…

1.  Зашифруйте текст с помощью AES - Advanced Encryption Standard в Go
2.  Затем мы рассмотрим запись этого зашифрованного сообщения в файл.
3.  Наконец, мы рассмотрим, как мы можем расшифровать это сообщение, используя общий секрет.

Исходя из этого, вы сможете создавать свои собственные простые системы шифрования, которые могут выполнять различные операции, например, шифровать файлы в вашей файловой системе и защищать их парольной фразой, которую знаете только вы, или добавлять простое шифрование в различные части систем, с которыми вы работаете. на.

## Вступление

Мы начнем с рассмотрения AES или Advanced Encryption Standard, поскольку это стандарт, который мы будем использовать для шифрования и дешифрования информации в наших программах Go.

Затем мы создадим действительно простую программу шифрования, которая будет принимать парольную фразу из командной строки и использовать ее вместе с AES для шифрования отрывка текста.

Как только это будет сделано, мы создадим программу-аналог, которая расшифрует этот отрывок текста, используя ту же парольную фразу, которую мы использовали для шифрования нашего текста.

## AES - расширенный стандарт шифрования

Итак, AES или Advanced Encryption Standard - это алгоритм шифрования с симметричным ключом, который первоначально был разработан двумя бельгийскими криптографами - Джоан Дэемен и Винсентом Риджменом.

Если вы хотите использовать шифрование в какой-либо из своих программ и не совсем уверены в том, чем все они различаются, то AES определенно является самым безопасным вариантом, поскольку он эффективен и прост в использовании.

> **Примечание.** Я расскажу о других методах шифрования в будущих уроках, поэтому следите за мной в твиттере: [@Elliot\_f](https://twitter.com/elliot_f)

### Симметричное шифрование ключа

Если вы не встречали термин «шифрование с симметричным ключом», не бойтесь, это относительно простая концепция, которая позволяет двум сторонам шифровать и расшифровывать информацию с использованием общего секрета.

## Наш клиент шифрования

Хорошо, давайте перейдем в наш любимый редактор кода и начнем писать код!

Мы начнем с создания нового файла с именем, `encrypt.go` который будет содержать кодовую фразу из командной строки и впоследствии использовать ее для шифрования некоторого текста перед записью его в файл.

Начнем с простого шифрования фрагмента текста заранее заданным ключом и распечатки результатов. Как только мы с этим справимся, мы можем добавить больше сложности:

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func main() {
    fmt.Println("Encryption Program v0.01")

    text := []byte("My Super Secret Code Stuff")
    key := []byte("passphrasewhichneedstobe32bytes!")

    // generate a new aes cipher using our 32 byte long key
    c, err := aes.NewCipher(key)
    // if there are any errors, handle them
    if err != nil {
        fmt.Println(err)
    }

    // gcm or Galois/Counter Mode, is a mode of operation
    // for symmetric key cryptographic block ciphers
    // - https://en.wikipedia.org/wiki/Galois/Counter_Mode
    gcm, err := cipher.NewGCM(c)
    // if any error generating new GCM
    // handle them
    if err != nil {
        fmt.Println(err)
    }

    // creates a new byte array the size of the nonce
    // which must be passed to Seal
    nonce := make([]byte, gcm.NonceSize())
    // populates our nonce with a cryptographically secure
    // random sequence
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println(err)
    }

    // here we encrypt our text using the Seal function
    // Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
    // time, for a given key.
    fmt.Println(gcm.Seal(nonce, nonce, text, nil))
}

```

Итак, если мы попытаемся запустить это с помощью, `go run encrypt.go` вы должны увидеть, что он распечатывает как `Hello World` массив байтов, так и наш зашифрованный `text` .

Давайте попробуем записать это в файл, добавив вызов `ioutile.WriteFile` вместо нашего последнего `fmt.Println` оператора:

```go
// the WriteFile method returns an error if unsuccessful
err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
// handle this error
if err != nil {
  // print it out
  fmt.Println(err)
}

```

> Для получения дополнительной информации о чтении и записи файлов в Go я предлагаю вам ознакомиться с моей статьей с метко названным названием « [Чтение и запись в файлы в Go»](https://tutorialedge.net/golang/reading-writing-files-in-go/) .

### Тестирование

Как только мы закончим вносить эти изменения в наш `encrypt.go` файл, мы можем попробовать его протестировать:

```s
$ go run encrypt.go

```

Если это работает успешно, вы должны увидеть новый файл, созданный в каталоге вашего проекта, с именем `myfile.data` . Если вы откроете это, вы должны увидеть результаты своего шифрования!

## Наш клиент дешифрования

Теперь, когда мы рассмотрели шифрование и запись нашего зашифрованного сообщения в файл, давайте посмотрим, как читать из этого файла и пытаться расшифровать его, используя тот же общий ключ.

Мы начнем с использования `ioutil.ReadFile('myfile.data')` , чтобы читать зашифрованный текст как массив байтов. Когда у нас будет этот массив байтов, мы сможем впоследствии выполнить очень похожие шаги, как мы делали для стороны шифрования.

1.  Сначала нам нужно создать новый Cipher с помощью `aes.NewCipher` функции, передав наш общий ключ в качестве основного параметра.
2.  Далее нам нужно сгенерировать наш GCM
3.  После этого нам нужно получить размер нашего одноразового номера, используя `gcm.NonceSize()`
4.  Наконец, мы расшифруем наш зашифрованный код, `ciphertext` используя `gcm.Open()` функцию, которая возвращает как наш, так `plaintext` и / или, `error` если он есть.

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Println("Decryption Program v0.01")

    key := []byte("passphrasewhichneedstobe32bytes!")
    ciphertext, err := ioutil.ReadFile("myfile.data")
    // if our program was unable to read the file
    // print out the reason why it can't
    if err != nil {
        fmt.Println(err)
    }

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        fmt.Println(err)
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        fmt.Println(err)
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(plaintext))
}

```

### Тестирование

Замечательно, теперь, когда мы закончили писать нашу `decrypt.go` программу, мы можем попробовать ее.

```s
$ go run decrypt.go
Decryption Program v0.01
My Super Secret Code Stuff

```

Как видно из выходных данных, мы смогли успешно прочитать зашифрованное содержимое нашего `myfile.data` файла и впоследствии расшифровать его с помощью нашего общего секретного ключа.

## Задача - зашифрованная файловая система

Если вас интересует задача, отличный способ проверить то, что вы узнали в этом руководстве, - это попытаться расширить существующую программу, которую мы создали выше, для шифрования и дешифрования любых файлов, переданных ей с использованием парольной фразы.

Вы могли бы потенциально превратить это в CLI, который принимает флаги и пути к файлам в качестве входных данных и выводит их в зашифрованном виде в ваше текущее местоположение.

*   [Создание интерфейса командной строки в Go](https://tutorialedge.net/golang/building-a-cli-in-go/)

## Заключение

Итак, в этом руководстве мы успешно рассмотрели ряд интересных концепций, таких как алгоритмы симметричного шифрования, а также способы шифрования и дешифрования информации с помощью Advanced Encryption Standard и секретного ключа.

Мне было очень весело писать это, и, надеюсь, вам понравилось! Если да, или если у вас есть отзывы, то я хотел бы услышать их в разделе комментариев ниже!
