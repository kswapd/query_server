package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	listenPort = flag.Uint("listen_port", 8085, "The port listen to.")
	listenIP   = flag.String("listen_ip", "", "The ip listen to.")
	remotePort = flag.Uint("remote_port", 8000, "The port of monitor server.")
	remoteHost = flag.String("remote_host", "54.223.95.178", "The host of monitor server.")
)

type handle struct {
	host string
	port uint
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.URL.String())
	remote, err := url.Parse(fmt.Sprintf("http://%s:%d", this.host, this.port))
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: *remoteHost, port: *remotePort}
	log.Println("listen: ", *listenIP, *listenPort)
	log.Println("backend: ", *remoteHost, *remotePort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *listenPort), h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main3() {
	flag.Parse()
	startServer()
}
