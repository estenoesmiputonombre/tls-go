package main

import (
	"fmt"
	"log"
	"net/http"
)

//HandleHelloWorld ... Handle hello being polite with the client
func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, fmt.Sprintf("Hello %s from golang", r.URL))
	fmt.Fprintf(w, "Hello %s from golang", r.URL)
}

//CreateServer ... Creates the server
func CreateServer() {
	http.HandleFunc("/", HandleHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	CreateServer()
}
