package main

import (
	"crypto/tls"
	"flag"
	"github.com/valyala/fasthttp"
	"log"
	"net"
)

var (
	addr        = flag.String("addr", "127.0.0.1:8080", "TCP address to listen to for http")
	tlsAddr     = flag.String("tlsAddr", "", "TCP address to listen to for https")
	tlsCertFile = flag.String("tlsCertFile", "", "Path to TLS certificate file")
	tlsKeyFile  = flag.String("tlsKeyFile", "", "Path to TLS key file")
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

	// Читаем TLS сертификат из файла
	cert, err := tls.LoadX509KeyPair(*tlsCertFile, *tlsKeyFile)
	if err != nil {
		log.Fatalf("cannot load cert for -tlsCertFile=%q, -tlsKeyFile=%q: %s",
		*tlsCertFile, *tlsKeyFile, err)
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
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		Certificates: []tls.Certificate{cert},
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
