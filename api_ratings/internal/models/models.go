package models

import (
	"github.com/gofrs/uuid"
)

type Rating struct {
	Id      int64      `json:"id"`
	UserId  *uuid.UUID `json:"user_id"`
	SongId  *uuid.UUID `json:"song_id"`
	Content string     `json:"Content"`
	Rating  float64    `json:"rating"`
	Date    string     `json:"date"`
}
