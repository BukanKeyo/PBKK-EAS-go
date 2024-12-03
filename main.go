package main

import (
	"go-movies/config"
	"go-movies/controllers/genrecontroller"
	"go-movies/controllers/homecontroller"
	"go-movies/controllers/moviecontroller"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Genre
	http.HandleFunc("/genres", genrecontroller.Index)
	http.HandleFunc("/genres/add", genrecontroller.Add)
	http.HandleFunc("/genres/edit", genrecontroller.Edit)
	http.HandleFunc("/genres/delete", genrecontroller.Delete)

	// 3. Movies
	http.HandleFunc("/movies", moviecontroller.Index)
	http.HandleFunc("/movies/add", moviecontroller.Add)
	http.HandleFunc("/movies/detail", moviecontroller.Detail)
	http.HandleFunc("/movies/edit", moviecontroller.Edit)
	http.HandleFunc("/movies/delete", moviecontroller.Delete)

	// Run server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
