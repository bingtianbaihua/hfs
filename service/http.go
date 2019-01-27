package service

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

type Config struct {
	Host string
	Port string

	Prefix string
	Dir    string
}

type HTTPServer struct {
	*http.Server
}

func (srv *HTTPServer) ListenAndServe() error {
	return srv.Server.ListenAndServe()
}

func (srv *HTTPServer) ListenAndServeTLS(certFile, keyFile string) error {
	return srv.Server.ListenAndServeTLS(certFile, keyFile)
}

func NewHTTPServer(cfg *Config) (Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config can not be empty")
	}

	handler := http.StripPrefix(cfg.Prefix, http.FileServer(http.Dir(cfg.Dir)))
	server := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
	return &HTTPServer{
		Server: server,
	}, nil
}
