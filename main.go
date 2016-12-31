package main

import (
	"flag"
	"log"
	"query_server/MicroService"
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	MicroService.Start()
}
