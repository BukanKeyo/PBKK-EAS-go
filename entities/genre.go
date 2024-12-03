package entities

import "time"

type Genre struct {
	Id         uint
	Name       string
	CreatedAt  time.Time
	UpddatedAt time.Time
}
