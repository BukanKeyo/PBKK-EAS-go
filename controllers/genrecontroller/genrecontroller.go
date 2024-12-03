package genrecontroller

import (
	"go-movies/entities"
	"go-movies/models/genremodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	genres := genremodel.GetAll()
	data := map[string]any{
		"genres": genres,
	}

	temp, err := template.ParseFiles("views/genre/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genre/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var genre entities.Genre

		genre.Name = r.FormValue("name")
		genre.CreatedAt = time.Now()
		genre.UpddatedAt = time.Now()

		ok := genremodel.Create(genre)
		if !ok {
			temp, _ := template.ParseFiles("views/genre/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/genre/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genre := genremodel.Detail(id)
		data := map[string]any{
			"genre": genre,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var genre entities.Genre

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		genre.Name = r.FormValue("name")
		genre.UpddatedAt = time.Now()

		if ok := genremodel.Update(id, genre); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/genres", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := genremodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/genres", http.StatusSeeOther)
}
