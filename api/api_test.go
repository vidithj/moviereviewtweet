package api

import (
	"99designs-coding-test/model"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ReadMoviesInfo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]model.Movie
	}{
		{
			name:  "Success Path",
			input: "../assets/movies.json",
			expected: map[string]model.Movie{
				"Star Wars":                   {Title: "Star Wars", Year: 1977},
				"Star Wars The Force Awakens": {Title: "Star Wars The Force Awakens", Year: 2015},
				"Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb": {Title: "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb", Year: 1964},
				"Plan 9 from outer space": {Title: "Plan 9 from outer space", Year: 1959},
			},
		},
		{
			name:     "Error Path",
			input:    "",
			expected: map[string]model.Movie{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movie, err := ReadMoviesInfo(tt.input)
			if diff := cmp.Diff(movie, tt.expected); diff != "" {
				if err != nil {
					t.Log("ReadMoviesInfo error occured :\n", err)
				}
				t.Fatalf("ReadMoviesInfo failed mismatch: (-got +want)\n%s", diff)
			} else {
				if err != nil {
					t.Log("ReadMoviesInfo error occured :\n", err)
				}
			}
		})
	}
}

func Test_ReadReviews(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []model.Review
	}{
		{
			name:  "Success Path",
			input: "../assets/reviews.json",
			expected: []model.Review{
				{Title: "Star Wars", Review: "Great, this film was", Score: 77},
				{Title: "Star Wars The Force Awakens", Review: "A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some chap came along and basically made the same movie again", Score: 50},
				{Title: "Metropolis", Review: "Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening because the audio was too low", Score: 15},
				{Title: "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb", Review: "Hilarious Kubrick satire", Score: 100},
				{Title: "Plan 9 from outer space", Review: "So bad it's bad", Score: 7},
			},
		},
		{
			name:     "Error Path",
			input:    "",
			expected: []model.Review{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			review, err := ReadReviews(tt.input)
			if diff := cmp.Diff(review, tt.expected); diff != "" {
				if err != nil {
					t.Log("ReadReviews error occured :\n", err)
				}
				t.Fatalf("ReadReviews failed mismatch: (-got +want)\n%s", diff)
			} else {
				if err != nil {
					t.Log("ReadReviews error occured :\n", err)
				}
			}
		})
	}
}

func Test_createStars(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "success Path 100",
			input:    100,
			expected: "★★★★★",
		},
		{
			name:     "Success Path 77",
			input:    77,
			expected: "★★★★",
		},
		{
			name:     "Success Path 50",
			input:    50,
			expected: "★★½",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stars := createStars(tt.input)
			if diff := cmp.Diff(stars, tt.expected); diff != "" {
				t.Fatalf("createStars failed mismatch: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_createTweet(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Tweet
		expected string
	}{
		{
			name: "success_Path: with year",
			input: model.Tweet{
				Title:  "Star Wars",
				Review: "Great, this film was",
				Score:  77,
				Year:   1977,
			},
			expected: "Star Wars(1977): Great, this film was ★★★★",
		},
		{
			name: "success_Path: without year",
			input: model.Tweet{
				Title:  "Metropolis",
				Review: "Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening because the audio was too low",
				Score:  15,
				Year:   0,
			},
			expected: "Metropolis: Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening because the audio was too low ★",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tweet := createTweet(tt.input)
			if diff := cmp.Diff(tweet, tt.expected); diff != "" {
				t.Fatalf("createTweet failed mismatch: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_MovieReview(t *testing.T) {
	tests := []struct {
		name         string
		inputmovie   map[string]model.Movie
		inputreviews []model.Review
		expected     []string
	}{
		{
			name: "success_Path",
			inputmovie: map[string]model.Movie{
				"Star Wars":                   {Title: "Star Wars", Year: 1977},
				"Star Wars The Force Awakens": {Title: "Star Wars The Force Awakens", Year: 2015},
				"Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb": {Title: "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb", Year: 1964},
				"Plan 9 from outer space": {Title: "Plan 9 from outer space", Year: 1959},
			},
			inputreviews: []model.Review{
				{Title: "Star Wars", Review: "Great, this film was", Score: 77},
				{Title: "Star Wars The Force Awakens", Review: "A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some chap came along and basically made the same movie again", Score: 50},
				{Title: "Metropolis", Review: "Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening because the audio was too low", Score: 15},
				{Title: "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb", Review: "Hilarious Kubrick satire", Score: 100},
				{Title: "Plan 9 from outer space", Review: "So bad it's bad", Score: 7},
			},
			expected: []string{
				"Star Wars(1977): Great, this film was ★★★★",
				"Star Wars The Force Awake(2015): A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some  ★★½",
				"Metropolis: Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening beca ★",
				"Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb(1964): Hilarious Kubrick satire ★★★★★",
				"Plan 9 from outer space(1959): So bad it's bad ½",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tweets := MovieReview(tt.inputmovie, tt.inputreviews)
			if diff := cmp.Diff(tweets, tt.expected); diff != "" {
				t.Fatalf("MovieReview failed mismatch: (-got +want)\n%s", diff)
			}
		})
	}
}
