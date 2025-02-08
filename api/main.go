package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handler() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão se não especificada
	}

	r := mux.NewRouter()
	r.HandleFunc("/my-first-go-project", HomeHandler).Methods("GET")
	http.Handle("/", r)
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Autorização necessária"))
		return
	}

	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		w.Write([]byte("Tipo de autorização: Bearer Token"))
	} else if len(authHeader) > 5 && authHeader[:5] == "Basic" {
		w.Write([]byte("Tipo de autorização: Basic Auth"))
	} else {
		w.Write([]byte("Tipo de autorização desconhecido"))
	}
}
