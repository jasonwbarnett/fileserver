package main

import (
	"log"
	"net/http"

	"github.com/jasonwbarnett/fileserver"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", fileserver.New(http.Dir("/usr/share/doc"))))
}
