package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// NewTelegramProxy returns a reverse-proxy handler that forwards /tg/{rest}
// to baseURL/{rest}, stripping the /tg prefix.
func NewTelegramProxy(baseURL string) (http.Handler, error) {
	target, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	return http.StripPrefix("/tg", proxy), nil
}
