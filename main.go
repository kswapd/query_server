package main

import (
    "query_server/MicroService"
    "log"
)


func main() {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
       MicroService.Start()
    }