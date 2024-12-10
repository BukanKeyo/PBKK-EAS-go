package controllers

import (
	"go-movies/entities"
	"go-movies/models"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func IndexMovie(w http.ResponseWriter, r *http.Request) {
	movies := models.GetAllMovie()
	data := map[string]any{
		"movies": movies,
	}

	temp, err := template.ParseFiles("views/movie/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/movie/create.html")
		if err != nil {
			panic(err)
		}

		genres := models.GetAllGenre()
		data := map[string]any{
			"genres": genres,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var movie entities.Movie

		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
		if err != nil {
			panic(err)
		}

		rating, err := strconv.Atoi(r.FormValue("rating"))
		if err != nil {
			panic(err)
		}

		movie.Name = r.FormValue("name")
		movie.Genre.Id = uint(genreId)
		movie.Rating = int64(rating)
		movie.Description = r.FormValue("description")
		movie.CreatedAt = time.Now()
		movie.UpdatedAt = time.Now()

		if ok := models.CreateMovie(movie); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/movies", http.StatusSeeOther)
	}
}

func DetailMovie(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	movie := models.DetailMovie(id)
	data := map[string]any{
		"movie": movie,
	}

	temp, err := template.ParseFiles("views/movie/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func EditMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/movie/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		movie := models.DetailMovie(id)
		genres := models.GetAllGenre()

		data := map[string]any{
			"movie":  movie,
			"genres": genres,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var movie entities.Movie

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genreId, err := strconv.Atoi(r.FormValue("genre_id"))
		if err != nil {
			panic(err)
		}

		rating, err := strconv.Atoi(r.FormValue("rating"))
		if err != nil {
			panic(err)
		}

		movie.Name = r.FormValue("name")
		movie.Genre.Id = uint(genreId)
		movie.Rating = int64(rating)
		movie.Description = r.FormValue("description")
		movie.UpdatedAt = time.Now()

		if ok := models.Update(id, movie); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/movies", http.StatusSeeOther)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/movies", http.StatusSeeOther)
}
