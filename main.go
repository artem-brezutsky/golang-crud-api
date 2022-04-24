package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbm     string    `json:"isbm"`
	Title    string    `json:"title"`
	Director *Director `json:"directors"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movieId := strconv.Itoa(rand.Intn(1000000))
	movie = Movie{
		ID:    movieId,
		Isbm:  strconv.Itoa(rand.Intn(1000000)),
		Title: "Title_" + movieId,
		Director: &Director{
			Firstname: "John_" + movieId,
			Lastname:  "Doe_" + movieId,
		},
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedMovie Movie
	for index, item := range movies {
		if item.ID == params["id"] {
			movies[index] = Movie{
				ID:    params["id"],
				Isbm:  "12355",
				Title: "New Title",
				Director: &Director{
					Firstname: "John_" + params["id"],
					Lastname:  "Doe_" + params["id"],
				},
			}

			updatedMovie = movies[index]
			break
		}
	}
	json.NewEncoder(w).Encode(updatedMovie)
	return
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
	return
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbm:  "432475",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbm:  "43244325",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "Adam",
			Lastname:  "Smith",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
