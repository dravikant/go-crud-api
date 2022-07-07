package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, movie := range movies {

		if movie.ID == params["id"] {

			json.NewEncoder(w).Encode(movie)
			break
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

//update is just deletion followed by addition
func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, movie := range movies {

		if movie.ID == params["id"] {

			movies = append(movies[:index], movies[index+1:]...)

			var newMovie Movie

			_ = json.NewDecoder(r.Body).Decode(&newMovie)

			newMovie.ID = params["id"]
			movies = append(movies, newMovie)
			json.NewEncoder(w).Encode(newMovie)
			return
		}
	}
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "First Movie", Director: &Director{FirstName: "abc", LastName: "xyz"}})
	movies = append(movies, Movie{ID: "2", Isbn: "14234", Title: "Second Movie", Director: &Director{FirstName: "abc", LastName: "jkl"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
