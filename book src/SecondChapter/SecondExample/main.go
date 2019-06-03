package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	//PORT ... Port to send an receive information.
	PORT = 8080
)

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", http.StripPrefix("/web/", fs))
	log.Printf("Listening on Port %d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
