package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
}
