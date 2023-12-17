package models



type Song struct {
	Id int64 `json:"id"`
	ArtistName string `json:"artistName"`
	Title string `json:"title"`
	DurationInMillis int64 `json:"durationInMillis"`
}