package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://localhost:8080/estenoesmiputonombre")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Body: %s", string(body))
}
