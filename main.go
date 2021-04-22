package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var port string

func main() {
	port = "5000"

	r := registerRoutes()
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/images/{name}", getImage)
	})
	fs := http.FileServer(http.Dir("/home/jose/Desktop/non-blocking/static-files"))
	r.Handle("/", fs)
	return r
}

func getImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
