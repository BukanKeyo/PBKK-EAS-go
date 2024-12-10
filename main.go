package main

import (
	"go-movies/config"
	"go-movies/controllers"
	"log"
	"net/http"
)

func main() {
	// 	Database connection
	config.ConnectDB()

	// Routes
	// 1.Homepage
	http.HandleFunc("/", controllers.Welcome)

	// 2. Genre
	http.HandleFunc("/genres", controllers.IndexGenre)
	http.HandleFunc("/genres/add", controllers.AddGenre)
	http.HandleFunc("/genres/edit", controllers.EditGenre)
	http.HandleFunc("/genres/delete", controllers.DeleteGenre)

	// 3. Movies
	http.HandleFunc("/movies", controllers.IndexMovie)
	http.HandleFunc("/movies/add", controllers.AddMovie)
	http.HandleFunc("/movies/detail", controllers.DetailMovie)
	http.HandleFunc("/movies/edit", controllers.EditMovie)
	http.HandleFunc("/movies/delete", controllers.DeleteMovie)

	// Run server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
