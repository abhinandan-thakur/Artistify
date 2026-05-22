package models

import (
	"time"
)

type Albums struct {
	ID        int       `json:"id"`
	AlbumName string    `json:"album_name"`
	Artist    string    `json:"artist"`
	Sales     int       `json:"sales"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
