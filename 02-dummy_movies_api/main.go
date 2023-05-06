package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie



func main() {
	r := mux.NewRouter()
	// seed data
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448743", Title: "Movie Two", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "3", Isbn: "448743", Title: "Movie Three", Director: &Director{FirstName: "John", LastName: "Doe"}})

    r.HandleFunc("/movies", getMovies).Methods("GET")
    r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
    r.HandleFunc("/movies/{id}", updateMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
    w.WriteHeader(http.StatusOK)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if id < 1 || id > len(movies) {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies[id-1])
    w.WriteHeader(http.StatusOK)
}


func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var movie Movie
    json.NewDecoder(r.Body).Decode(&movie)
    movies = append(movies, movie)
    w.WriteHeader(http.StatusCreated)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    if id < 1 || id > len(movies) {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    var movie Movie
    json.NewDecoder(r.Body).Decode(&movie)
    movies[id-1] = movie
    w.WriteHeader(http.StatusOK)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            break
        }
    }
	w.WriteHeader(http.StatusOK)
}


