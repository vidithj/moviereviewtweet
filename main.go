package main

import (
	"99designs-coding-test/api"
	"fmt"
	"os"
	"strings"
)

const (
	assetspath = "assets/"
)

func main() {
	//read command line arg to get movie and review json
	arg := os.Args
	reviews, err := api.ReadReviews(assetspath + arg[1])
	if err != nil {
		fmt.Println("Error while fetching movies : ", err)
	}
	movies, err := api.ReadMoviesInfo(assetspath + arg[2])
	if err != nil {
		fmt.Println("Error while fetching reviews : ", err)
	}
	//create tweets
	tweets := api.MovieReview(movies, reviews)
	//format and print tweets
	fmt.Println("Tweets:")
	fmt.Println("-------")
	fmt.Println(strings.Join(tweets, "\n"))
}
