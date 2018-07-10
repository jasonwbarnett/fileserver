package main

import (
	"log"
	"net/http"

	jsonfs_http "github.com/jasonwbarnett/jsonfs/net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", jsonfs_http.FileServer(jsonfs_http.Dir("/usr/share/doc"))))
}
