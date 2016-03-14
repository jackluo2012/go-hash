package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-hash/pwhash"
	"log"
	"net/http"
)

func main() {
	// Create route
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hash/{password}", HashIndex)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HashIndex(w http.ResponseWriter, r *http.Request) {
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

	// Response hash
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(hash); err != nil {
		panic(err)
	}
}
