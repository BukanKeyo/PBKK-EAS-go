package moviecontroller

import (
	"go-movies/entities"
	"go-movies/models/genremodel"
	"go-movies/models/moviemodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	movies := moviemodel.Getall()
	data := map[string]any{
		"movies": movies,
	}

	temp, err := template.ParseFiles("views/movie/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/movie/create.html")
		if err != nil {
			panic(err)
		}

		genres := genremodel.GetAll()
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

		if ok := moviemodel.Create(movie); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/movies", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	movie := moviemodel.Detail(id)
	data := map[string]any{
		"movie": movie,
	}

	temp, err := template.ParseFiles("views/movie/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
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

		movie := moviemodel.Detail(id)
		genres := genremodel.GetAll()

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

		if ok := moviemodel.Update(id, movie); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/movies", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := moviemodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/movies", http.StatusSeeOther)
}
