package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
)

//ParsingJSON ... Parses JSON 
func ParsingJSON(w http.ResponseWriter, r *http.Request, a interface{}) ([]byte, error) {
    var (
        buff []byte
        err error
    )
    w.Header().Set(http.CanonicalHeaderKey("content-type"), "application/json")
    if buff, err = json.Marshal(a); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{ "Error": "Problems trying to unmarshal the buffer" }`))
        return nil, err
    }
    w.Write(buff)
    return buff, nil
}

func UnparsingJSON(w http.ResponseWriter, r *http.Request, a interface{}) error {
    defer r.Body.Close()
    buff, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set(http.CanonicalHeaderKey("content-type"), "application/json")
        w.Write([]byte(`{ "Error": "Problems trying to unmarshal the buffer" }`))
        return err
    }
    if err := json.Unmarshal(buff, &a); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set(http.CanonicalHeaderKey("content-type"), "application/json")
        w.Write([]byte(`{ "Error": "Problems trying to unmarshal the buffer" }`))
        return err
    } else {
        return nil
    }
}

//Abs ... Return the positive value of the integer
func Abs(a int) int {
    if a < 0 {
        return a*-1
    }
    return a
}
