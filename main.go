package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", r)
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
