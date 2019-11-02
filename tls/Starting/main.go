package main

import (
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set(http.CanonicalHeaderKey("content-type"), "text/plain")
        w.Write([]byte("Hello world"))
    })

    log.Println("Starting the server at http://localhost:8080")
    http.ListenAndServe(":8080", nil)

}
