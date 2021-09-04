package main

import (
	"log"
)

func main() {
	server, err := NewHTTPServer()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server dijalankan di %s port %d ...\n", server.host, server.port)

	server.Start()
}
