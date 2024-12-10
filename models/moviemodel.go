package models

import (
	"go-movies/config"
	"go-movies/entities"
)

func GetAllMovie() []entities.Movie {
	rows, err := config.DB.Query(`
		SELECT 
			movies.id, 
			movies.name, 
			genres.name as genre_name,
			movies.rating, 
			movies.description, 
			movies.created_at, 
			movies.updated_at FROM movies
		JOIN genres ON movies.genre_id = genres.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var movies []entities.Movie

	for rows.Next() {
		var movie entities.Movie
		if err := rows.Scan(
			&movie.Id,
			&movie.Name,
			&movie.Genre.Name,
			&movie.Rating,
			&movie.Description,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		); err != nil {
			panic(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

func CreateMovie(movie entities.Movie) bool {
	result, err := config.DB.Exec(`
		INSERT INTO movies(
			name, genre_id, rating, description, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?)`,
		movie.Name,
		movie.Genre.Id,
		movie.Rating,
		movie.Description,
		movie.CreatedAt,
		movie.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func DetailMovie(id int) entities.Movie {
	row := config.DB.QueryRow(`
		SELECT 
			movies.id, 
			movies.name, 
			genres.name as genre_name,
			movies.rating, 
			movies.description, 
			movies.created_at, 
			movies.updated_at FROM movies
		JOIN genres ON movies.genre_id = genres.id
		WHERE movies.id = ?
	`, id)

	var movie entities.Movie

	err := row.Scan(
		&movie.Id,
		&movie.Name,
		&movie.Genre.Name,
		&movie.Rating,
		&movie.Description,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return movie
}

func Update(id int, movie entities.Movie) bool {
	query, err := config.DB.Exec(`
		UPDATE movies SET 
			name = ?, 
			genre_id = ?,
			rating = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?`,
		movie.Name,
		movie.Genre.Id,
		movie.Rating,
		movie.Description,
		movie.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM movies WHERE id = ?", id)
	return err
}
