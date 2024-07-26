package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func deleteMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func createMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Int(100000000))
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

// func updateMovies(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(req)
// 	for _, item := range movies {
// 		if item.ID == params["id"] {

// 		}
// 	}
// }

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "3493", Title: "Movie1", Director: &Director{Firstname: "Paul", Lastname: "Molly"}})
	movies = append(movies, Movie{ID: "2", Isbn: "3444", Title: "Movie2", Director: &Director{Firstname: "Kelly", Lastname: "White"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
