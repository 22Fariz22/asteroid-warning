package entity

type NeoCounts struct {
	NeoCounts []Asteroid `json:"neo_counts"`
}

type Asteroid struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
