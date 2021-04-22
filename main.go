package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string
var dir = "/home/jose/Desktop/non-blocking/static-files"

func main() {
	port = "5000"
	r := mux.NewRouter()
	r.HandleFunc("/api/images/{name}", getImage)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	r.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getImage(w http.ResponseWriter, r *http.Request) {
	if _, exists := mux.Vars(r)["name"]; !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	name := mux.Vars(r)["name"]
	bytes, err := ioutil.ReadFile(dir + "/images/" + name + ".jpg")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	image := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
	w.Write([]byte(image))
}
