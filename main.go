package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello world"))
	})
	log.Println("starting on port 3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
