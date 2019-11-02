package main

import (
    "log"
    "net/http"
    "io/ioutil"
)

func main() {
    if buff, err := sendRequestGet("http://localhost:8080/hello"); err != nil {
        log.Fatal(err)
    } else {
        log.Println(string(buff))
    }
}

func sendRequestGet(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer resp.Body.Close()
    if buff, err := ioutil.ReadAll(resp.Body); err != nil {
        return nil, err
    } else {
        return buff, nil
    }
}
