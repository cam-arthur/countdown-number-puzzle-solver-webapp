package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc(`/getResults`, countdownHandler).Methods("POST")
	r.HandleFunc("/returnLatest", getLatestResultHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err.Error())
	}
}
