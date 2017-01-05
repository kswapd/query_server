package main

import (
	"flag"
	"log"
	"query_server/MicroService"
	"query_server/Monitor"
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	Monitor.CreateInfluxDBClient()
	MicroService.Start()

}
