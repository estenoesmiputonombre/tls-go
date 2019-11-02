package main

import (
    "net/http"
    "strings"
    "log"
)

//SiftCipher ... Controller used to get back the result of this algorithm
func SiftCipher(w http.ResponseWriter, r *http.Request) {
   if strings.EqualFold(r.Method, http.MethodPost) {
        var c Cipher
        if err := UnparsingJSON(w, r, &c); err != nil {
            log.Fatal(err)
        } else {
            if len(c.Text) != 0 {
                c.Crypto()
                ParsingJSON(w, r, c)
            } else {
                w.WriteHeader(http.StatusBadRequest)
                w.Header().Set(http.CanonicalHeaderKey("content-type"), "text/plain")
                w.Write([]byte("You should use a correct JSON"))
            }
        }
   } else {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set(http.CanonicalHeaderKey("content-type"), "text/plain")
        w.Write([]byte("You should use a POST method"))
   }
}

//BruteForceSiftCipher ... Try to resolve the chipertext into a comprehensible text
func BruteForceSiftCipher(w http.ResponseWriter, r *http.Request) {
    if strings.EqualFold(r.Method, http.MethodPost) {
        var c Cipher
        if err := UnparsingJSON(w, r, &c); err != nil {
            log.Fatal(err)
        } else {
            if len(c.Text) != 0 {
                arr := c.DesCrypto()
                ParsingJSON(w, r, arr)
            } else {
                w.WriteHeader(http.StatusBadRequest)
                w.Header().Set(http.CanonicalHeaderKey("content-type"), "text/plain")
                w.Write([]byte("You should use a correct JSON"))
            }
        }
   } else {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set(http.CanonicalHeaderKey("content-type"), "text/plain")
        w.Write([]byte("You should use a POST method"))
   }
}

