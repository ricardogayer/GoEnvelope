package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type envelope map[string]interface{}

type Movie struct {
	ID  int64    			`json:"ident"`
	CreatedAt time.Time     `json:"data_criacao"`
	Title string            `json:"titulo"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movie",showMovieHandler).GetHandler()
	router.HandleFunc("/movies",showMoviesHandler).GetHandler()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func showMovieHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type","application/json")

		movie := Movie{
			ID:        1000,
			CreatedAt: time.Now(),
			Title:     "Casablanca",
		}

	json.NewEncoder(w).Encode(envelope{"movie": movie})

}

func showMoviesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type","application/json")

	var movies = []Movie{}

	movie1 := Movie{
		ID:        1000,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
	}
	movies = append(movies,movie1)

	movie2 := Movie{
		ID:        2300,
		CreatedAt: time.Now(),
		Title:     "Matrix",
	}

	movies = append(movies,movie2)

	movie3 := Movie{
		ID:        3500,
		CreatedAt: time.Now(),
		Title:     "Senhor dos Anéis",
	}

	movies = append(movies,movie3)


	json.NewEncoder(w).Encode(envelope{"movies": movies})

}
