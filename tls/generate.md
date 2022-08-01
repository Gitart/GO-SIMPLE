## Автоматизируем бесплатное получение и обновление TLS-сертификатов

Многие уже слышали про прекрасный сервис letsencrypt.org, который выдает всем желающим бесплатные TLS-сертификаты. 
И эти сертификаты признаются всеми современными браузерами. Ниже показано, насколько просто добавить поддержку 
автоматического получения и обновления TLS-сертификатов letsencrypt.org в наш сервер на Go:

```go
package main

import (
	"crypto/tls"
	"flag"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net"
)

var (
	addr        = flag.String("addr", "127.0.0.1:8080", "TCP address to listen to for http")
	tlsAddr     = flag.String("tlsAddr", "", "TCP address to listen to for https")
	tlsCertFile = flag.String("tlsCertFile", "", "Path to TLS certificate file. "+
		"The certificate is automatically generated and put "+
		"to -autocertCacheDir if empty")
	tlsKeyFile = flag.String("tlsKeyFile", "", "Path to TLS key file. "+
		"The key is automatically generated and put "+
		"to -autocertCacheDir if empty")
	autocertCacheDir = flag.String("autocertCacheDir", "autocert-cache",
		"Path to the directory where letsencrypt certs are cached")
)

func main() {
	flag.Parse()

	// Пытаемся запустить https сервер
	startTLS()

	// Запускаем http сервер
	log.Printf("Serving http on -addr=%q", *addr)
	err := fasthttp.ListenAndServe(*addr, handler)
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func startTLS() {
	if len(*tlsAddr) == 0 {
		log.Printf("-tlsAddr is empty, so skip serving https")
		return
	}

	// Создаем net.Listener'а, который принимает подключения по -tlsAddr.
	ln, err := net.Listen("tcp4", *tlsAddr)
	if err != nil {
		log.Fatalf("cannot listen for -tlsAddr=%q: %s", *tlsAddr, err)
	}

	// Создаем требуемую конфигурацию tls.
	// См. https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/ .
	tlsConfig := tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
	}

	if len(*tlsCertFile) > 0 {
		// Читаем TLS сертификат из файла
		cert, err := tls.LoadX509KeyPair(*tlsCertFile, *tlsKeyFile)
		if err != nil {
			log.Fatalf("cannot load cert for -tlsCertFile=%q, -tlsKeyFile=%q: %s", *tlsCertFile, *tlsKeyFile, err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	} else {
		// Настраиваем автоматическое создание и обновление сертификатов.
		m := autocert.Manager{
			Prompt: autocert.AcceptTOS,

			// Сертификаты будут кэшироваться в -autocertCacheDir,
			// чтобы при рестарте сервера не приходилось
			// пересоздавать их снова.
			Cache: autocert.DirCache(*autocertCacheDir),
		}
		tlsConfig.GetCertificate = m.GetCertificate
	}

	// Создаем net.Listener'а для tls подключений поверх созданного
	// выше net.Listener'а
	tlsLn := tls.NewListener(ln, &tlsConfig)

	// запускаем https сервер в отдельном потоке
	log.Printf("Serving https on -tlsAddr=%q", *tlsAddr)
	go fasthttp.Serve(tlsLn, handler)
}

func handler(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello, world!\n")
}
```


Перед компиляцией данного кода понадобится скачать еще один сторонний пакет — golang.org/x/crypto/acme/autocert, который отвечает 
за автоматическое создание и обновление TLS-сертификатов:

```
$ go get -u golang.org/x/crypto/acme/autocert
```

Теперь сервер будет автоматически создавать и обновлять TLS-сертификаты для всех hostname’ов, 
запрошенных по https адресу -tlsAddr, если не указан -tlsCertFile. 
Выписанные сертификаты будут кэшироваться в каталоге -autocertCacheDir.

