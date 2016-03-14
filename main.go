package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wp-password-hash/pwhash"
)

func main() {
	// Create route
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("hash/{password}", HashIndex)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HashIndex(w http.ReponseWriter, r http.Request) {
	// Get password from url
	vars := mux.Vars(r)
	password := vars["password"]

	// Generate random salt
	setting, err := pwhash.GenSaltPrivate(6)

	if err != nil {
		panic(err)
	}

	// Hash the password
	hash := pwhash.CryptPrivate(password, setting)
}
