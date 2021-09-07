package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func configureHost(s *HTTPServer) error {
	hostStr, ok := os.LookupEnv("ARTICLE_HOST")
	if !ok {
		return nil
	}

	if ip := net.ParseIP(hostStr); ip == nil {
		return nil
	}

	s.host = hostStr
	return nil
}

func configurePort(s *HTTPServer) error {
	portStr, ok := os.LookupEnv("ARTICLE_PORT")
	if !ok {
		return nil
	}

	port, err := strconv.ParseUint(portStr, 10, 16)

	if err != nil {
		return err
	}

	s.port = uint16(port)
	return nil
}

func main() {
	server, err := NewHTTPServer(configureHost, configurePort)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server dijalankan di %s port %d ...\n", server.host, server.port)

	server.Start()
}
