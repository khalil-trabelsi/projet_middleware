package models

type Rating struct {
	UserId  int64   `json:"userId"`
	SongId  int64   `json:"songId"`
	Comment string  `json:"comment"`
	Note    float64 `json:"note"`
}
