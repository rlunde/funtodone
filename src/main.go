package main

import (
	"fmt"
	//"html"
	"github.com/gorilla/mux"
	"gopkg.in/authboss.v0"
	"log"
	//"gopkg.in/authboss.v0/auth"
	"./model"
	"net/http"
	"os"
)

func main() {
	ab := authboss.New() // Usually store this globally
	ab.MountPath = "/authboss"
	ab.LogWriter = os.Stdout

	xyz := model.String(nil)
	fmt.Println(xyz)

	if err := ab.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}

	// Make sure to put authboss's router somewhere
	http.Handle("/authboss", ab.NewRouter())

	r := mux.NewRouter()
	// r.HandleFunc("/search/{searchTerm}", Search)
	// r.HandleFunc("/load/{dataId}", Load)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/")))
	http.Handle("/", r)
	log.Println("starting http server on http://localhost:8100")
	http.ListenAndServe(":8100", nil)
}
