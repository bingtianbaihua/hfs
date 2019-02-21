package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bingtianbaihua/hfs/server"
)

const (
	version = "0.0.1"
)

var (
	host, port, prefix, dir, v string
)

func main() {
	flag.StringVar(&v, "v", "", "release version")
	flag.StringVar(&host, "host", "0.0.0.0", "ip address")
	flag.StringVar(&port, "port", "8910", "listen http port")
	flag.StringVar(&prefix, "prefix", "/", "http route path")
	flag.StringVar(&dir, "dir", "./", "http file directory")

	flag.Parse()
	fmt.Printf("address: %v, port: %v, url prefix: %v, serve directory: %v\n", host, port, prefix, dir)

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
