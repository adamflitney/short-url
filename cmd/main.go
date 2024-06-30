package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/teris-io/shortid"
)

func main() {
	fmt.Println("hello world!")

	urlMappings := make(map[string]string)

	router := http.NewServeMux()

	router.HandleFunc("GET /short-url/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from get"))
	})

	router.HandleFunc("POST /short-url/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		newId, err := shortid.Generate()
		if err != nil {
			log.Fatalln(err)
		}
		urlMappings[newId] = string(body)

		w.Write([]byte(newId))
	})

	fs := http.FileServer(http.Dir("./web"))

	router.Handle("/", fs)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
