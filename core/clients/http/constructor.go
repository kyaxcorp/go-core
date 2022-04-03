package http

import (
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
)

func New() *resty.Client {
	return resty.New()
}

func NewWithClient(hc *http.Client) *resty.Client {
	return resty.NewWithClient(hc)
}

func NewWithLocalAddr(localAddr net.Addr) *resty.Client {
	return resty.NewWithLocalAddr(localAddr)
}
