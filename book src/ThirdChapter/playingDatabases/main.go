package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	//PORT ... Port to connect to golang service
	PORT = os.Getenv("PORT")
)

func main() {
	gorilla := mux.NewRouter()
	gorilla.HandleFunc("/employers/{amount:[0-9]{1,2}}", ServePage)
	http.Handle("/", gorilla)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}

//ServePage ... Serves the page to addr:port/whatever
func ServePage(w http.ResponseWriter, r *http.Request) {
}
