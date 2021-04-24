package api

import (
	"99designs-coding-test/model"
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

//ReadMoviesInfo Read movies json file
func ReadMoviesInfo(moviepath string) (map[string]model.Movie, error) {
	movies := make(map[string]model.Movie, 0)
	var movie []model.Movie
	movieFile, err := os.Open(moviepath)
	if err != nil {
		return movies, err
	}
	defer movieFile.Close()
	movieValue, _ := ioutil.ReadAll(movieFile)
	json.Unmarshal(movieValue, &movie)
	for _, movieInfo := range movie {
		movies[movieInfo.Title] = movieInfo
	}
	return movies, nil
}

//ReadReviews Read reviews json file
func ReadReviews(reviewpath string) (review []model.Review, err error) {
	reviewFile, err := os.Open(reviewpath)
	if err != nil {
		return []model.Review{}, err
	}
	defer reviewFile.Close()
	reviewValue, _ := ioutil.ReadAll(reviewFile)
	json.Unmarshal(reviewValue, &review)
	return review, nil
}

//createStars create stars based on review score
func createStars(score int) string {
	new_score := math.Round(float64(score) / 10)
	full_stars := ""
	i := 0
	for i < int(new_score/2) {
		full_stars += "★"
		i++
	}
	half_star := ""
	if int(new_score)%2 != 0 {
		half_star = "½"
	}
	stars := full_stars + half_star
	return stars
}

//createTweet create tweet by appending title,year,review and stars
func createTweet(tweetModel model.Tweet) string {
	stars := createStars(tweetModel.Score)
	yearStr := ""
	if tweetModel.Year != 0 {
		yearStr = "(" + strconv.Itoa(tweetModel.Year) + ")"
	}
	return tweetModel.Title + yearStr + ": " + tweetModel.Review + " " + stars
}

//MovieReview create tweets array containing tweet for each review
func MovieReview(movies map[string]model.Movie, reviews []model.Review) []string {
	tweets := make([]string, 0)
	for _, review := range reviews {
		movie := movies[review.Title]
		tweetModel := model.Tweet{Title: review.Title, Review: review.Review, Score: review.Score, Year: movie.Year}
		tweet := createTweet(tweetModel)
		if len(tweet) > 140 {
			if len(tweetModel.Title) > 25 {
				tweetModel.Title = tweetModel.Title[0:25]
			}
			tweet = createTweet(tweetModel)
		}
		if len(tweet) > 140 {
			//if tweet more than 140 , get stars and create tweet of len 140
			stars := createStars(tweetModel.Score)
			if len(stars) > 1 {
				// if stars, then create a tweet of length 140-len(stars)-1 and then add stars for total tweet len 140
				tweet = tweet[0:(140-len(stars)-1)] + " " + stars
			} else {
				// if no stars, then directly make the len as 140
				tweet = tweet[0:140]
			}
		}
		tweets = append(tweets, tweet)
	}
	return tweets
}
