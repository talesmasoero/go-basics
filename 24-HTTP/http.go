package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, t *http.Request) {
		w.Write([]byte("HOME PAGE"))
	})

	http.HandleFunc("/user", userHandler)

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("USER PAGE"))
}
