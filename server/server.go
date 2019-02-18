package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/bingtianbaihua/hfs/log"
	"github.com/bingtianbaihua/hfs/middleware"
	"github.com/bingtianbaihua/hfs/model"
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

	fileAdapter, err := middleware.NewFileAdapter(&middleware.FileAdapterConfig{Prefix: cfg.Prefix, Dir: cfg.Dir})
	logAdapter, err := middleware.NewLogHandler(new(middleware.LogConfig))
	recoverAdapter := middleware.NewRecoverAdapter()
	if err != nil {
		log.Info("error:%v", err)
		return nil, err
	}

	// build handler chains
	chains := model.Build(fileAdapter.FileHandle(), logAdapter.HandleTask, recoverAdapter.HandleTask)

	server := &http.Server{
		Addr:         net.JoinHostPort(cfg.Host, cfg.Port),
		Handler:      chains,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
		IdleTimeout:  time.Duration(10) * time.Second,
	}
	return &HTTPServer{
		Server: server,
	}, nil
}
