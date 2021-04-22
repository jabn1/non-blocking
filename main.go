package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var port string
var dir = "/home/jose/Desktop/non-blocking/static-files"
var isSync bool
var mutex = sync.Mutex{}

func main() {
	a := flag.Bool("s", false, "")
	flag.Parse()
	isSync = *a
	if isSync {
		fmt.Println("Started in synchronous mode.")
	} else {
		fmt.Println("Started in asynchronous mode.")
	}
	port = "5000"
	r := mux.NewRouter()
	r.HandleFunc("/api/images/{name}", getImage)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	r.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getImage(w http.ResponseWriter, r *http.Request) {
	if isSync {
		mutex.Lock()
	}
	if _, exists := mux.Vars(r)["name"]; !exists {
		w.WriteHeader(http.StatusNotFound)
		if isSync {
			mutex.Unlock()
		}
		return
	}
	name := mux.Vars(r)["name"]
	bytes, err := ioutil.ReadFile(dir + "/images/" + name + ".jpg")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if isSync {
			mutex.Unlock()
		}
		return
	}

	image := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
	w.Write([]byte(image))
	if isSync {
		mutex.Unlock()
	}
}
