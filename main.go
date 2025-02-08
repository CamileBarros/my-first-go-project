package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FavoriteSongs struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   string `json:"year"`
	Rate   int    `json:"rate"`
}

var favoriteSongs = []FavoriteSongs{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetFavoriteSongs).Methods("GET")
	r.HandleFunc("/", CreateFavoriteSongs).Methods("POST")
	http.Handle("/", r)
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetFavoriteSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favoriteSongs)
}

func CreateFavoriteSongs(w http.ResponseWriter, r *http.Request) {
	var song FavoriteSongs
	_ = json.NewDecoder(r.Body).Decode(&song)
	favoriteSongs = append(favoriteSongs, song)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favoriteSongs)
}
