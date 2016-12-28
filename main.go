package main

import (
    "MicroService"
    "log"
)


func main() {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
       MicroService.Start()
    }