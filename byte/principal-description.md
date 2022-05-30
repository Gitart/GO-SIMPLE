 ## Подробное объяснение Golang bytes.buffer

#### Введение буфера
**Buffer** - это структура типа Buffer в пакете байтов {…}
A buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value for Buffer is an empty buffer ready to use.
(Это буфер переменной длины с методами чтения и записи. Нулевое значение Buffer - пустой буфер, но его можно использовать)
Буфер похож на контейнер-контейнер, вы можете хранить вещи, получать их (данные доступа)

Создать буфер
```
    func main() {
        buf1 := bytes.NewBufferString("hello")
        buf2 := bytes.NewBuffer([]byte("hello"))
        buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
        fmt.Printf("%v,%v,%v\n",buf1,buf2,buf3)
        fmt.Printf("%v,%v,%v\n",buf1.Bytes(),buf2.Bytes(),buf3.Bytes())
     
        buf4 := bytes.NewBufferString("")
        buf5 := bytes.NewBuffer([]byte{})
        fmt.Println(buf4.Bytes(),buf5.Bytes())
    }
```
Выход
```
    hello,hello,hello
    [104 101 108 108 111],[104 101 108 108 111],[104 101 108 108 111]
    [] []
```
#### Запись в буфер
Новый буфер пуст, и его также можно записать напрямую.

Write
```
    func (b *Buffer) Write(p []byte) (n int,err error)

    func main() {
    s := []byte(" world")
    buf := bytes.NewBufferString("hello")
    fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
    buf.Write(s)
    fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
    }
```
результат
```
    hello,[104 101 108 108 111]
    hello world,[104 101 108 108 111 32 119 111 114 108 100]
```
WriteString
```
    func (b *Buffer) WriteString(s string)(n int,err error)

    func main() {
        s := " world"
        buf := bytes.NewBufferString("hello")
        fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
        buf.WriteString(s)
        fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
    }
```
результат
```
    hello,[104 101 108 108 111]
    hello world,[104 101 108 108 111 32 119 111 114 108 100]
```
####  WriteByte
```
    func (b *Buffer) WriteByte(c byte) error

    func main() {
        var s byte = '?'
        buf := bytes.NewBufferString("hello")
        fmt.Println(buf.Bytes()) // [104 101 108 108 111]
        buf.WriteByte(s)
        fmt.Println(buf.Bytes())    // [104 101 108 108 111 63]
    }
```
#### WriteRune
```
    func (b *Buffer) WriteRune(r Rune) (n int,err error)

    func main(){
       var s rune = 'хорошо'
       buf := bytes.NewBufferString("hello")
       fmt.Println(buf.String()) //hello
       buf.WriteRune(s)   
           fmt.Println (buf.String ()) // привет хорошо
    }
```
результат
```
    22909
    [104 101 108 108 111]
    [104 101 108 108 111 229 165 189]
```
#### Записать из буфера
```
    func main() {
        file,_ := os.Create("test.txt")
        buf := bytes.NewBufferString("hello world")
        buf.WriteTo(file)
    }
```
#### Читать буфер
Read
```
    func (b *Buffer) Read(p []byte)(n int,err error)

    func main() {
        s1 := []byte("hello")
        buff :=bytes.NewBuffer(s1)
        s2 := []byte(" world")
        buff.Write(s2)
        fmt.Println(buff.String())  //hello world
     
        s3 := make([]byte,3)
        buff.Read(s3)
             fmt.Println (string (s3)) // Емкость hel, s3 равна 3, только 3 можно прочитать
        fmt.Println(buff.String()) //lo world
     
             buff.Read (s3) // перезапишет s3
        fmt.Println(string(s3))  // lo 
        fmt.Println(buff.String())  // world
    }
```
#### ReadByte
Вернуть первый байт заголовка буфера
```
    func (b *Buffer) ReadByte() (c byte,err error)

    func main() {
        buf := bytes.NewBufferString("hello")
        fmt.Println(buf.String())  // hello
        b,_:= buf.ReadByte()
        fmt.Println(string(b))  //h
        fmt.Println(buf.String())   //ello
    }
```
#### ReadRun
Метод ReadRune, возвращает первую руну заголовка буфера.
```
    func (b *Buffer) ReadRune() (r rune,size int,err error)

    func main() {
     
             buf1: = bytes.NewBufferString ("Привет, xuxiaofeng")
        fmt.Println(buf1.Bytes()) //[228 189 160 229 165 189 120 117 120 105 97 111 102 101 110 103]
        b1,n1,_ := buf1.ReadRune()
             fmt.Println (string (b1)) // вы
        fmt.Println(n1)
        
        
        buf := bytes.NewBufferString("hello")
        fmt.Println(buf.String())  //hello
        b,n,_:= buf.ReadRune()
        fmt.Println(n) // 1
        fmt.Println(string(b))  //h
        fmt.Println(buf.String())   //ello
    }
```
Почему n == 3 и n1 == 1? Давайте посмотрим на исходный код ReadRune
```
    func (b *Buffer) ReadRune() (r rune, size int, err error) {
        if b.empty() {
            // Buffer is empty, reset to recover space.
            b.Reset()
            return 0, 0, io.EOF
        }
        c := b.buf[b.off]
             if c <utf8.RuneSelf {// Здесь можно судить, является ли первый прочитанный символ Rune
            b.off++
            b.lastRead = opReadRune1
            return rune(c), 1, nil
        }
        r, n := utf8.DecodeRune(b.buf[b.off:])
        b.off += n
        b.lastRead = readOp(n)
        return r, n, nil
    }
```
#### ReadBytes
Для метода ReadBytes в качестве разделителя требуется байт. При чтении найдите первый разделитель, который появляется в буфере, и верните байт из начала буфера в разделитель. .
```
    func (b *Buffer) ReadBytes(delim byte) (line []byte,err error)

    func main() {
        var d byte = 'f'
        buf := bytes.NewBufferString("xuxiaofeng")
        fmt.Println(buf.String())
     
        b,_ :=buf.ReadBytes(d)
        fmt.Println(string(b))
        fmt.Println(buf.String())
    }
```
Эквивалент разделителю

#### ReadString
похож на метод readBytes

Читать в буфер
ReadFrom, считывает содержимое r в буфер из r, реализующего интерфейс io.Reader, а n возвращает количество операций чтения.
```
    func (b *Buffer) ReadFrom(r io.Reader) (n int64,err error)

    func main(){
       file, _ := os.Open("text.txt")
       buf := bytes.NewBufferString("bob ")
       buf.ReadFrom(file)
       fmt.Println(buf.String()) //bob hello world
    }
```
#### Удалить из буфера
Метод Next, возвращает первые n байтов (срез), исходный буфер изменяется
```
    func (b *Buffer) Next(n int) []byte

    func main() {
        buf := bytes.NewBufferString("helloworld")
        fmt.Println(buf.String())  // helloworld
        b := buf.Next(2)
        fmt.Println(string(b))  // he
    }
```
## Введение в принцип буфера
Нижний слой байтового буфера перехода хранится в байтовых срезах. Срез имеет длину len и ограничение емкости. Запись в буфер начинается с позиции длины len. Когда len> cap, она автоматически Расширение. Чтение буфера будет начинать чтение с позиции выключения встроенной метки (выключение всегда записывает начальную позицию чтения), когда выключено == len, это указывает, что буфер был полностью прочитан
и сбросить буфер (len = off = 0), кроме того, когда длина содержимого + длина записи (например, len) <= cap / 2, буфер перемещается вперед и перезаписывается Прочитанное содержимое (off = 0, len- = off), чтобы избежать непрерывного расширения буфера
```go
    func main() {
        byteSlice := make([]byte, 20) 
             byteSlice [0] = 1 // Установить первый байт буфера в 1
             byteBuffer: = bytes.NewBuffer (byteSlice) // Создаем 20-байтовый буфер len = 20 off = 0
        c, _ := byteBuffer.ReadByte()                     // off+=1
             fmt.Printf ("len:% d, c =% d \ n", byteBuffer.Len (), c) // len = 20 off = 1 print c = 1
        byteBuffer.Reset()                                // len = 0 off = 0
             fmt.Printf ("len:% d \ n", byteBuffer.Len ()) // выводим len = 0
             byteBuffer.Write ([] byte ("привет, байтовый буфер")) // записываем буфер len + = 17
             fmt.Printf ("len:% d \ n", byteBuffer.Len ()) // выводим len = 17
             byteBuffer.Next (4) // пропустить 4 байта + = 4
             c, _ = byteBuffer.ReadByte () // считываем 5-й байт + = 1
             fmt.Printf ("5-й байт:% d \ n", c) // Печать: 111 (соответствует букве o) len = 17 off = 5
             byteBuffer.Truncate (3) // Установите количество небайтов равным 3 len = off + 3 = 8 off = 5
             fmt.Printf ("len:% d \ n", byteBuffer.Len ()) // вывести len = 3 как количество непрочитанных байтов, выше len = 8 - это длина нижележащего среза
             byteBuffer.WriteByte (96) // len + = 1 = 9 меняем y на A
        byteBuffer.Next(3)                                // len=9 off+=3=8
        c, _ = byteBuffer.ReadByte()                      // off+=1=9    c=96
             fmt.Printf ("9-й байт:% d \ n", c) // печать: 96
    }
    ```
