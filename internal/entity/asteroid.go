package entity

type NeoCounts struct {
	NeoCounts []Asteroid `json:"neo_counts"`
}

type Asteroid struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type NeoList struct {
	Links struct {
		Self, Next, Prev string
	}
	Start            string                `json:"-"`
	End              string                `json:"-"`
	ElementCount     int64                 `json:"element_count"`
	NearEarthObjects map[string][]Asteroid `json:"near_earth_objects"`
}
