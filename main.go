package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	isbn string `json:"isbn"`
	Name string `json:"name"`
	Year string `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	
}

var movies []Movie

// Get all movies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get a movie
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Create a movie
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update a movie
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID: "1", isbn: "448743", Name: "The Shawshank Redemption", Year: "1994", Director: &Director{FirstName: "Frank", LastName: "Darabont"}})
	movies = append(movies, Movie{ID: "2", isbn: "448744", Name: "The Godfather", Year: "1972", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "3", isbn: "448745", Name: "The Dark Knight", Year: "2008", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "4", isbn: "448746", Name: "The Godfather: Part II", Year: "1974", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "5", isbn: "448747", Name: "12 Angry Men", Year: "1957", Director: &Director{FirstName: "Sidney", LastName: "Lumet"}})
	movies = append(movies, Movie{ID: "6", isbn: "448748", Name: "Schindler's List", Year: "1993", Director: &Director{FirstName: "Steven", LastName: "Spielberg"}})
	movies = append(movies, Movie{ID: "7", isbn: "448749", Name: "The Lord of the Rings: The Return of the King", Year: "2003", Director: &Director{FirstName: "Peter", LastName: "Jackson"}})
	movies = append(movies, Movie{ID: "8", isbn: "448750", Name: "Pulp Fiction", Year: "1994", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "9", isbn: "448751", Name: "The Lord of the Rings: The Fellowship of the Ring", Year: "2001", Director: &Director{FirstName: "Peter", LastName: "Jackson"}})
	movies = append(movies, Movie{ID: "10", isbn: "448752", Name: "Forrest Gump", Year: "1994", Director: &Director{FirstName: "Robert", LastName: "Zemeckis"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}