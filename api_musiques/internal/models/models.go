package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id *uuid.UUID `json:"id"`
	ArtistName string `json:"artistName"`
	Title string `json:"title"`
	DurationInMillis int `json:"durationInMillis"`
}