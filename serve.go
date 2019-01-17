package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.FileServer(http.Dir("."))
	http.Handle("/", srv)

	err := http.ListenAndServe(":8080", nil)
	if (err != nil) {
		log.Fatalf("Could not start http server %s", err)
	}
}