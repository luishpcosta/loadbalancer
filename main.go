package main

import (
	"flag"
	"log"

	"github.com/luishpcosta/loadbalancer/application"
)

func main() {
	var serverList string
	var port int
	flag.StringVar(&serverList, "backends", "", "Load balanced backends, use commas to separate")
	flag.IntVar(&port, "port", 80, "Port to serve")
	flag.Parse()

	if len(serverList) == 0 {
		log.Fatal("Please provide one or more backends to load balance")
	}

	application.Start(serverList, port)
}
