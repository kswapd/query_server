package main

import (
	"log"
	"query_server/MicroService"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	MicroService.Start()
}
