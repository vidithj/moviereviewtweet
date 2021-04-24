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
	Score  int    `json:"score"`
}

//Tweet struct for holding information for creating tweet
type Tweet struct {
	Title  string
	Review string
	Score  int
	Year   int
}
