package main

import (
	"crypto/rand"
	"net/http"
	"testing"

	. "github.com/emicklei/forest"
)

func ExpectErrorCode(t *testing.T, r *http.Response, code float64) bool {
	if got, want := JSONPath(t, r, ".code"), code; got != want {
		msg := JSONPath(t, r, ".message")
		Errorf(t, "got %v want %v got message: %v", got, want, msg)
		return false
	}
	return true
}

//create random strings
func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
