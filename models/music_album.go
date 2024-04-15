package models

import "time"

type MusicAlbum struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name" binding:"required,min=5"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Genre       string    `json:"genre"`
	Price       float64   `json:"price" binding:"required,min=100,max=1000"`
	Description string    `json:"description"`
}
