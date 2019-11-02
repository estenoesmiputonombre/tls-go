package main

import (
    "log"
)

//Handlers ... create the handlers for the server
func Handlers() {
    log.Println("Creating the handlers")
    mux.HandleFunc("/cipher", SiftCipher)
    mux.HandleFunc("/descipher", BruteForceSiftCipher)
}
