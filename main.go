package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // fetched all url params
	for _, data := range movies {

		if data.ID == params["id"] {
			json.NewEncoder(w).Encode(data)
			return
		}
	}

	json.NewEncoder(w).Encode("No Data Found")
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // fetched all url params
	for index, data := range movies {

		if data.ID == params["id"] {
			json.NewEncoder(w).Encode(data)
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}

	json.NewEncoder(w).Encode("No Data Found")
}

func createMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){

	// set json content Type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over movies and range

	for index, data := range movies {

		// delte the movie with id sent in response 
		if data.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}

	// add a new movie given from request body 
	var movie Movie

	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)


}

func main() {
	
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "RRR", Director: &Director{FirstName: "RajMouli", LastName: "Sir"}})
	movies = append(movies, Movie{ID: "2", Isbn: "124", Title: "Pushpa", Director: &Director{FirstName: "Karansrei", LastName: "Pans"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}