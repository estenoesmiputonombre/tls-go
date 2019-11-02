package main

import (
    "log"
    "net/http"
)

var (
    mux = http.NewServeMux()
)

func main() {
   log.Println("Server listen on http://localhost:8080")
   Handlers()
   server := &http.Server {
        Addr: "127.0.0.1:8080",
        Handler: mux,
   }
   server.ListenAndServe()
}
