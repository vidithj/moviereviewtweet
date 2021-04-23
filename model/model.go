package model

// Movie struct for movie json file
type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year,omitempty"`
}

// Review struct for review json file
type Review struct {
	Title  string `json:"title"`
	Review string `json:"review"`
	Score  string `json:"score"`
}
