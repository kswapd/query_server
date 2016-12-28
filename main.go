package main

import (
    "MicroSettrvice"
    "log"
)


func main() {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
       MicroService.Start()
    }