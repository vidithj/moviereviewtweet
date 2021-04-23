package api

import (
	"99designs-coding-test/model"
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadMoviesInfo(moviepath string) (movie []model.Movie, err error) {
	movieFile, err := os.Open(moviepath)
	if err != nil {
		return []model.Movie{}, err
	}
	defer movieFile.Close()
	movieValue, _ := ioutil.ReadAll(movieFile)
	json.Unmarshal(movieValue, &movie)
	return movie, nil
}

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

func MovieReview(movie []model.Movie, review []model.Review) {

	// create tweets

}
