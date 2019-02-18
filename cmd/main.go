package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bingtianbaihua/hfs/server"
)

const (
	helpStr = `
fs version: fs/0.01
Usage: fs [-?h] [-host ip address] [-port port] [-path http handler path] [-dir file directory]

Options:
  -?,-h         : this help
  -v            : show version and exit
  -V            : show version and configure options then exit
  -host         : fs work work address, default: 0.0.0.0
  -port			: fs listen http port
  -prefix       : http handler path
  -dir			: fileserver directory
`
)

var (
	host, port, prefix, dir, help string
)

func main() {
	flag.StringVar(&help, "h", "", helpStr)
	flag.StringVar(&host, "host", "0.0.0.0", "ip address")
	flag.StringVar(&port, "port", "8910", "listen http port")
	flag.StringVar(&prefix, "prefix", "/", "http route path")
	flag.StringVar(&dir, "dir", "./", "http file directory")

	flag.Parse()
	fmt.Println(host, port, prefix, dir)

	cfg := &server.Config{
		Host:   host,
		Port:   port,
		Prefix: prefix,
		Dir:    dir,
	}

	srv, err := server.NewHTTPServer(cfg)
	if err != nil {
		log.Fatalf("new server error:%v", err)
		return
	}

	srv.ListenAndServe()
}
