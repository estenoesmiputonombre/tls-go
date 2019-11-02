package main

import (
    "strings"
)

//Cryptography ... This is the interface that we will use to crypto methods
type Crytography interface {
    Crypto()
    DesCrypto() []string
}

//Cipher ... it manages the structure of the request cipher
type Cipher struct {
    Amount int `json:"sift"`
    Text string `json:"text"`
}

const (
    //MaxCipher ... Max number that we will use to make the mod
    MaxCipher = 26
)

//Crypto ... Method to encrypt the text into a ciphertext
func (c *Cipher) Crypto() {
    var localAmount = Abs(c.Amount)
    if localAmount > MaxCipher {
        localAmount %= MaxCipher
    }
    c.Text = strings.Map(func(r rune) rune {
        switch {
            case r >= 'a' && r <= 'z':
                return 'a' + (r-'a'+int32(localAmount))%MaxCipher
            case r >= 'A' && r <= 'Z':
                return 'A' + (r-'A'+int32(localAmount))%MaxCipher
        }
        return r
    }, c.Text)
}

//DesCrypto ... Uses a brute force to desencrypt the text
func (c *Cipher) DesCrypto() []Cipher {
    ciphers := make([]Cipher, MaxCipher-1)
    for i := 1; i < MaxCipher; i++ {
         ciphers[i-1] = Cipher{
            Text: strings.Map(func(r rune) rune {
                switch {
                case r >= 'a' && r <= 'z':
                    return 'a'+(r-'a'+int32(i))%MaxCipher
                case r >= 'A' && r <= 'Z':
                    return 'A'+(r-'A'+int32(i))%MaxCipher
                }
                return r
            }, c.Text),
        Amount: i,
        }
    }
    return ciphers
}

