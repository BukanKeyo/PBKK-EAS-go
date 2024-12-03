package genremodel

import (
	"go-movies/config"
	"go-movies/entities"
)

func GetAll() []entities.Genre {
	rows, err := config.DB.Query(`SELECT * FROM genres`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var genres []entities.Genre

	for rows.Next() {
		var genre entities.Genre
		if err := rows.Scan(&genre.Id, &genre.Name, &genre.CreatedAt, &genre.UpddatedAt); err != nil {
			panic(err)
		}

		genres = append(genres, genre)
	}

	return genres
}

func Create(genre entities.Genre) bool {
	result, err := config.DB.Exec(`
		INSERT INTO genres (name, created_at, updated_at) 
		VALUE (?, ?, ?)`,
		genre.Name,
		genre.CreatedAt,
		genre.UpddatedAt,
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

func Detail(id int) entities.Genre {
	row := config.DB.QueryRow(`SELECT id, name FROM genres WHERE id = ? `, id)

	var genre entities.Genre

	if err := row.Scan(&genre.Id, &genre.Name); err != nil {
		panic(err.Error())
	}

	return genre
}

func Update(id int, genre entities.Genre) bool {
	query, err := config.DB.Exec(`UPDATE genres SET name = ?, updated_at = ? where id = ?`, genre.Name, genre.UpddatedAt, id)
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
	_, err := config.DB.Exec("DELETE FROM genres WHERE id = ?", id)
	return err
}
