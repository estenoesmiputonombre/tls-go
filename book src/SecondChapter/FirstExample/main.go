package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	//PORT ... Port to send and receive information
	PORT = 8080
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("The id {%s} is up", vars["id"])
	log.Printf("./files/%s.html", vars["id"])
	fileName := fmt.Sprintf("files/%s.html", vars["id"])
	_, err := os.Stat(fileName)
	if err != nil {
		fileName = fmt.Sprintf("files/%d.html", 404)
	}
	http.ServeFile(w, r, fileName)
}
