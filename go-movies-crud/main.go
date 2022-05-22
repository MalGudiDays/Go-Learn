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
	ID       string    "json:'id'"
	Isbn     string    "json:'isbn'"
	Title    string    "json:'title'"
	Director *Director "json:;director'"
}

type Director struct {
	firstName string "json:'firstname'"
	lastNmae  string "json:'lastname'"
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, elem := range movies {
		if elem.ID == params["id"] {
			//inefficient method
			//movies = append(movies[:index], movies[index + 1:])
			//efficient one without copying all elems one by one order doesn't matter
			movies[index] = movies[len(movies)-1]
			movies = movies[:len(movies)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, elem := range movies {
		if elem.ID == params["id"] {
			movies[index] = movies[len(movies)-1]
			movies = movies[:len(movies)-1]

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, elem := range movies {
		if elem.ID == params["id"] {
			json.NewEncoder(w).Encode(elem)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "41", Title: "bhoolbhulaiya", Director: &Director{"aaa", "bbb"}})
	movies = append(movies, Movie{ID: "2", Isbn: "42", Title: "bhoolbhulaiya", Director: &Director{"aba", "bcb"}})
	movies = append(movies, Movie{ID: "3", Isbn: "43", Title: "bhoolbhulaiya", Director: &Director{"aca", "bdb"}})
	movies = append(movies, Movie{ID: "4", Isbn: "44", Title: "bhoolbhulaiya", Director: &Director{"ada", "beb"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8000", r))

}
