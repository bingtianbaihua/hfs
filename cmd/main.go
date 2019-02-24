package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/bingtianbaihua/hfs/log"

	"github.com/bingtianbaihua/hfs/server"
)

const (
	version = "0.0.1"
)

var (
	host, port, prefix, dir, v, cert, key, tls string
)

func main() {
	flag.StringVar(&v, "v", "", "release version")
	flag.StringVar(&host, "host", "0.0.0.0", "ip address")
	flag.StringVar(&port, "port", "8910", "listen http port")
	flag.StringVar(&prefix, "prefix", "/", "http route path")
	flag.StringVar(&dir, "dir", "./", "http file directory")
	flag.StringVar(&cert, "cert", "", "https cert file")
	flag.StringVar(&key, "key", "", "https key file")
	flag.StringVar(&tls, "tls", "false", "enable https")

	flag.Parse()

	cfg := &server.Config{
		Host:   host,
		Port:   port,
		Prefix: prefix,
		Dir:    dir,
	}

	srv, err := server.NewHTTPServer(cfg)
	if err != nil {
		log.Error("new server error:%v", err)
		return
	}

	b, err := strconv.ParseBool(tls)
	if err != nil {
		b = false
	}

	if b && cert != "" && key != "" {
		log.Info("Starting HTTPS server on: %v", net.JoinHostPort(cfg.Host, cfg.Port))
		log.Warn("HTTPS server error: %v", srv.ListenAndServeTLS(cert, key))
	} else {
		log.Info("Starting HTTP server on: %v", net.JoinHostPort(cfg.Host, cfg.Port))
		log.Warn("http server error: %v", srv.ListenAndServe())
	}
}
