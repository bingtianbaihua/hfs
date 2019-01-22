package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Path string `json:"path"`
	Dir  string `json:"dir"`
}

var (
	addr, port, path, dir, help string

	helpStr = `
fs version: fs/0.01
Usage: fs [-?h] [-addr ip address] [-port port] [-path http handler path] [-dir file directory]

Options:
  -?,-h         : this help
  -v            : show version and exit
  -V            : show version and configure options then exit
  -addr			: fs work work address, default: 0.0.0.0
  -port			: fs listen http port
  -path			: http handler path
  -dir			: fileserver directory
`
)

func NewFileServerHandler(cfg *Config) http.Handler {
	if cfg == nil {
		return nil
	}

	handler := http.FileServer(http.Dir(cfg.Dir))
	server := &http.Server{
		Addr:        cfg.Addr + ":" + cfg.Port,
		Handler:     handler,
		ReadTimeout: time.Duration(10) * time.Second,
	}
	server.ListenAndServe()

	return nil
}

func main() {
	flag.StringVar(&help, "h", "", helpStr)
	flag.StringVar(&addr, "addr", "0.0.0.0", "ip address")
	flag.StringVar(&port, "port", "8910", "listen http port")
	flag.StringVar(&path, "path", "/", "http route path")
	flag.StringVar(&dir, "dir", "./", "http file directory")

	flag.Parse()
	fmt.Println(addr, port, path, dir)

	cfg := &Config{
		Addr: addr,
		Port: port,
		Path: path,
		Dir:  dir,
	}

	NewFileServerHandler(cfg)
}
