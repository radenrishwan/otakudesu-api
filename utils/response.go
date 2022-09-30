package utils

import (
	"fmt"
	"net/http"
)

type DefaultResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type HomePageResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
	Episode   string `json:"episode"`
}

type AnimeListResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type AnimeDetailResponse struct {
	Id            string               `json:"id"`
	Title         string               `json:"title"`
	Thumbnail     string               `json:"thumbnail"`
	JapaneseTitle string               `json:"japanese_title"`
	Score         string               `json:"score"`
	Producer      string               `json:"producer"`
	Type          string               `json:"type"`
	Status        string               `json:"status"`
	Duration      string               `json:"duration"`
	RealeaseDate  string               `json:"realease_date"`
	Studio        string               `json:"studio"`
	Genre         string               `json:"genre"`
	Synopsis      string               `json:"synopsis"`
	Episode       []AnimeDetailEpisode `json:"episode"`
}

type AnimeDetailEpisode struct {
	Id         string `json:"id"`
	Episode    string `json:"episode"`
	Url        string `json:"url"`
	UploadDate string `json:"upload_date"`
}

type EpisodeDetail struct {
	Id    string    `json:"id"`
	Title string    `json:"title"`
	Urls  []UrlInfo `json:"urls"`
}

type UrlInfo struct {
	Host       string
	Url        string
	Size       string
	Resolution string
}

type FindAnimeResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
}

func NewSuccessResponse(resp string, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	_, err := fmt.Fprint(w, resp)
	PanicIfError(err, w, r)
}

func NewCustomResponse(resp string, code int, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err := fmt.Fprint(w, resp)
	PanicIfError(err, w, r)
}
