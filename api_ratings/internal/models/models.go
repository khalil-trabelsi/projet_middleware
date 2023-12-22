package models

type Rating struct {
	Id      int64   `json:"id"`
	UserId  int64   `json:"userId"`
	SongId  int64   `json:"songId"`
	Comment string  `json:"comment"`
	Note    float64 `json:"note"`
}
