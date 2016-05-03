// Package itunesart provides few helper methods to get album, artist or track
// artworks from the Itunes API
package itunesart

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

const apiUrlTrack = "https://itunes.apple.com/search?media=music&entity=musicTrack&limit=1&term="
const apiUrlAlbum = "https://itunes.apple.com/search?media=music&entity=album&limit=1&term="

// The Result represents the specific size of artworks returned by the Itunes API
type Result struct {
	Tiny    string
	Small   string
	Medium  string
	Default string
}

type httpResult struct {
	Tiny    string `json:"artworkUrl30"`
	Small   string `json:"artworkUrl60"`
	Medium  string `json:"artworkUrl100"`
	Default string
}

type httpResponse struct {
	ResultCount int          `json:resultCount`
	Results     []httpResult `json:results`
}

// Build all the artworks into size typed object for easy access
// { Result.SizeName }
// e.g: Result.Small would return the url for a small size artwork
func buildResult(result httpResult) (Result, error) {
	v := reflect.ValueOf(result)
	min := false
	res := Result{
		result.Tiny,
		result.Small,
		result.Medium,
		result.Default,
	}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).String()

		if len(value) > 0 {
			min = true
			res.Default = value
		}
	}

	if !min {
		return res, errors.New("No artwork was found")
	}

	return res, nil
}

// Parse http response and build the result
func parseResults(data []byte) (Result, error) {
	resp := httpResponse{}

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Result{}, err
	}

	if resp.ResultCount == 0 {
		return Result{}, errors.New("No match was found")
	}

	return buildResult(resp.Results[0])
}

// Executes an http request and returns error or response body
func request(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}

// AlbumCover gets the album artworks art from the Itunes database through out it's
// dedicated API.
func AlbumCover(album string, artist string) (Result, error) {
	url := apiUrlAlbum + url.QueryEscape(album+" "+artist)

	data, err := request(url)
	if err != nil {
		return Result{}, err
	}

	return parseResults(data)
}

// TrackCover gets the track artworks from the Itunes database through out it's
// dedicated API.
func TrackCover(track string, artist string) (Result, error) {
	url := apiUrlTrack + url.QueryEscape(track+" "+artist)

	data, err := request(url)
	if err != nil {
		return Result{}, err
	}

	return parseResults(data)
}
