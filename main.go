package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./data")))

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
