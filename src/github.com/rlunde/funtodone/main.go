package main

import (
	//"fmt"
	//"html"
	//"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/search/{searchTerm}", Search)
	// r.HandleFunc("/load/{dataId}", Load)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/")))
	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)
}
