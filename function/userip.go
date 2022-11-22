// Пакет userip предоставляет функции для извлечения IP-адреса пользователя 
// из запроса и связывания его с Context.
package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// FromRequest извлекает пользовательский IP адрес из req, если присуствует.
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

// Тип ключа не экспортируется для предотвращения конфликтов 
// с ключами контекста, определенными в других пакетах.
type key int

// userIPkey - это контекстный ключ для IP-адреса пользователя. 
// Его нулевое значение произвольно. 
// Если этот пакет определит другие контекстные ключи, 
// они будут иметь разные целочисленные значения.
const userIPKey key = 0

// NewContext возвращает новый Context, который несет предоставленное значение userIP.
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// FromContext извлекает пользовательский IP адрес из ctx, если присуствует.
func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
