// Package itunesart provides few helper methods to get album, artist or track
// artworks from the Itunes API
package itunesart

import (
    "errors"
    "reflect"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

const apiUrlTrack = "https://itunes.apple.com/search?media=music&entity=musicTrack&limit=1&term="
const apiUrlAlbum = "https://itunes.apple.com/search?media=music&entity=album&limit=1&term="

type Result struct {
    Tiny    string      `json:"artworkUrl30,omitempty"`
    Small    string      `json:"artworkUrl60,omitempty"`
    Medium   string      `json:"artworkUrl100,omitempty"`
    Default         string
}

type httpResponse struct {
    ResultCount int             `json:resultCount`
    Results     []Result        `json:results`
}

// Build all the artworks into size typed object for easy access
// { Result.SizeName }
// e.g: Result.Small would return the url for a small size artwork
func buildResult(result Result) (Result, error) {
    v := reflect.ValueOf(result)
    res := Result{}
    min := false

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
    url := apiUrlAlbum + url.QueryEscape(album + " " + artist)

    data, err := request(url)
    if err != nil {
        return Result{}, err
    }

    return parseResults(data)
}

// TrackCover gets the track artworks from the Itunes database through out it's
// dedicated API.
func TrackCover(track string, artist string) (Result, error) {
    url := apiUrlTrack + url.QueryEscape(track + " " + artist)

    data, err := request(url)
    if err != nil {
        return Result{}, err
    }

    return parseResults(data)
}
