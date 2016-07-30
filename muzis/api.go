package muzis

import (
	"github.com/go-resty/resty"
	"encoding/json"
	_"fmt"
)

const (
	SEACH_API = "http://muzis.ru/api/search.api"
	SEACH_PERFORMER_API = "http://muzis.ru/api/get_songs_by_performer.api"
)

type SearchResponse struct {
	Songs     *[]Song        `json:"songs"`
	Performer *[]Performer        `json:"performers"`
}

type Song struct {
	Type          int        `json:"type"`
	TrackId       int           `json:"id"`
	TrackName     string        `json:"track_name"`
	Lyrics        string        `json:"lyrics"`
	PerformerName string        `json:"performer"`
	PerformerId   int           `json:"performer_id"`
	Poster        string        `json:"poster"`
	Url           string        `json:"file_mp3"`
	Duration      int        `json:"timestudy"`
}

type Performer struct {
	PerformerId   int        `json:"id"`
	PerformerName string        `json:"title"`
	Poster        string        `json:"poster"`
}

type MuzisApi struct {

}

func (api *MuzisApi) GetSongsByPerformerName(name string) []Song {
	performer := getPerformerByName(name)
	return api.GetSongsByPerformerId(performer[0].PerformerId)

}

func (api *MuzisApi) GetSongsByPerformerId(performerId int) []Song {
	response, err := resty.R().
		SetFormData(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).
		SetFormData(map[string]string{
		"performer_id": string(performerId),
	}).
		Post(SEACH_PERFORMER_API)

	if err != nil {
		panic(err)
	}

	responseAsBytes := []byte(response.String())

	searchResponse := &SearchResponse{}
	if err := json.Unmarshal(responseAsBytes, searchResponse); err != nil {
		panic(err)
	}

	return *searchResponse.Songs
}

func getPerformerByName(name string) []Performer {
	response, err := resty.R().
		SetFormData(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).
		SetFormData(map[string]string{
		"q_performer": name,
	}).
		Post(SEACH_API)

	if err != nil {
		panic(err)
	}

	responseAsBytes := []byte(response.String())
	searchResponse := &SearchResponse{}
	if err := json.Unmarshal(responseAsBytes, searchResponse); err != nil {
		panic(err)
	}

	return *searchResponse.Performer
}