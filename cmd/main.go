package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/post/{id}", Post)

	http.ListenAndServe(":8080", r)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index Page."))

}

func Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	w.Write([]byte("Post Page."))

}
