package main

import (
	"99designs-coding-test/api"
	"fmt"
	"os"
)

const (
	assetspath = "assets/"
)

func main() {
	//read command line arg
	arg := os.Args
	movies, err := api.ReadMoviesInfo(assetspath + arg[1])
	if err != nil {
		fmt.Println("Error while fetching movies : ", err)
	}
	reviews, err := api.ReadReviews(assetspath + arg[2])
	if err != nil {
		fmt.Println("Error while fetching reviews : ", err)
	}
	api.MovieReview(movies, reviews)
}
