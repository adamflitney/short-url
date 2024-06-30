package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world!")

	router := http.NewServeMux()

	router.HandleFunc("GET /short-url/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from get"))
	})

	router.HandleFunc("POST /short-url/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from post"))
	})

	fs := http.FileServer(http.Dir("./web"))

	router.Handle("/", fs)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
