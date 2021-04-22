package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string
var dir = "/home/jose/Desktop/non-blocking/static-files"

func main() {
	port = "5000"
	r := mux.NewRouter()
	r.HandleFunc("/api/images/", getImage)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	r.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getImage(w http.ResponseWriter, r *http.Request) {
	// name := chi.URLParam(r, "name")
	// bytes, err := ioutil.ReadFile(dir + "/name/" + name + ".jpg")
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// }
	w.Write([]byte("test"))
}
