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

	router.HandleFunc("GET /short-url/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		body, ok := (urlMappings)[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("short url with id %s not found", id)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
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
