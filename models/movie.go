package models

type Movie struct {
	Title   string             `json:"title"`
	Related map[string]float64 `json:"related"`
}
