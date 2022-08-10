package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id 				string 		`json:"id"`
	Title 		string 		`json:"title"`
	Director 	*Director `json:"director"`
}

type Director struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	return
}

func main() {
	router := mux.NewRouter()

	// Mock Data
	movies = append(movies, Movie{Id: "1", Title: "Se7en", Director: &Director{ Id: "1", Name: "David Fincher" }})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	log.Fatal(http.ListenAndServe(":1337", router))
}