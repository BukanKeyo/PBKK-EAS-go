package entities

import "time"

type Movie struct {
	Id          uint
	Name        string
	Genre       Genre
	Rating      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
