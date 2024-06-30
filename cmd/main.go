package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teris-io/shortid"
)

func main() {
	fmt.Println("hello world!")

	urlMappings := make(map[string]string)

	router := http.NewServeMux()

	router.HandleFunc("GET /short-url/{id}", func(w http.ResponseWriter, r *http.Request) {
		for k := range urlMappings {
			log.Print(k)
		}

		id := r.PathValue("id")
		url, ok := urlMappings[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("short url with id %s not found", id)
			return
		}
		http.Redirect(w, r, url, http.StatusFound)
	})

	router.HandleFunc("POST /short-url/", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		url := r.Form.Get("url")

		if url != "" {
			newId, err := shortid.Generate()
			if err != nil {
				log.Fatalln(err)
			}
			urlMappings[newId] = string(url)
			w.Write([]byte(newId))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request"))
		}
	})

	fs := http.FileServer(http.Dir("./web"))

	router.Handle("/", fs)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
