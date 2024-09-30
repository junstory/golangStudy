package main

import (
	"errors"
	"io"
	"net/http"
)

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(string(body))
	}
	return string(body), nil
}

const APIKEY = "193ef3a"

// omdbapi에서 반환된 json 구조체
// 예시의 간략화를 위해 일부 값은 매핑하지 않음
type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}
