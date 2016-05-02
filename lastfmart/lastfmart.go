// Package lastfmart provides few helper methods to get album, artist or track
// cover art from the Last.fm API
package lastfmart

import (
    "errors"
    "reflect"
    "net/http"
    "io/ioutil"
    "unicode/utf8"
    "encoding/json"
)

var APIKey string
var APICorrect bool
const APIUrl = "http://ws.audioscrobbler.com/2.0/?format=json&method="

type Result struct {
    Small       string
    Medium      string
    Large       string
    ExtraLarge  string
    Mega        string
    Default     string
}

type image struct {
    Size    string      `json:size`
    Url     string      `json:"#text"`
}

type album struct {
    Name    string      `json:name`
    Image   []image     `json:image`
}

type track struct {
    Name    string      `json:name`
    Album   *album      `json:album`
}

type artist struct {
    Name    string      `json:name`
    Image   []image     `json:image`
}

type httpResponse struct {
    Album   *album      `json:album`
    Artist  *artist     `json:artist`
    Track   *track      `json:track`
}

type httpError struct {
    Error   *int        `json:error`
    Message *string     `json:message`
}

// AutoCorrect activates the autocorrect parameter in the Last.fm query url to
// notify the Last.fm API to fix spelling mistakes
// Note: Result may not be as expected
func AutoCorrect(act bool) {
    APICorrect = true
}

// Configure must be called before calling any other requests to set the Last.fm API Key
func Configure(key string) {
    APIKey = key
    APICorrect = false
}

func setDefaultCover(res Result) Result {
    if utf8.RuneCountInString(res.Default) > 0 {
        return res
    }

    v := reflect.ValueOf(res)
    res.Default = v.Field(0).String()
    return res
}

// Build all the images into size typed object for easy access
// { Result.SizeName }
// e.g: Result.Small would return the url for a small size cover art
func buildResult(images []image) (Result, error) {
    res := Result{}
    min := false

    for _, value := range images {
        if utf8.RuneCountInString(value.Url) > 0 {
            min = true

            switch value.Size {
            default:
                res.Default = value.Url
            case "small":
                res.Small = value.Url
            case "medium":
                res.Medium = value.Url
            case "large":
                res.Large = value.Url
            case "extralarge":
                res.ExtraLarge = value.Url
            case "mega":
                res.Mega = value.Url
            }
        }
    }

    if !min {
        return res, errors.New("No image was found")
    }

    return setDefaultCover(res), nil
}

// Parse http response and build results based on requested type
// parse { album, artist, track }
func parseResults(data []byte, parse string) (Result, error) {
    resp := httpResponse{}

    err := json.Unmarshal(data, &resp)
    if err != nil {
        return Result{}, err
    }

    switch parse {
    default:
        return Result{}, errors.New("No image was found")
    case "album":
        if resp.Album != nil {
            return buildResult(resp.Album.Image)
        }
    case "artist":
        if resp.Artist != nil {
            return buildResult(resp.Artist.Image)
        }
    case "track":
        if resp.Track != nil && resp.Track.Album != nil {
            return buildResult(resp.Track.Album.Image)
        }
    }

    return Result{}, errors.New("No image was found")
}

// Executes an http request and returns error or response body
func request(url string) ([]byte, error) {
    resErr := httpError{}
    resp, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(body, &resErr)
    if err != nil {
        return nil, err
    }

    if resErr.Error != nil {
        err = errors.New(*resErr.Message)
    }

    return body, err
}

// AlbumCover gets the album cover art from the Last.fm database throught it's
// dedicated API.
func AlbumCover(album string, artist string) (Result, error) {
    url := APIUrl + "album.getinfo&api_key=" + APIKey + "&album=" + album + "&artist=" + artist

    if APICorrect {
        url += "&autocorrect=1"
    }

    data, err := request(url)
    if err != nil {
        return Result{}, err
    }

    return parseResults(data, "album")
}

// ArtistCover gets the artist cover art from the Last.fm database throught it's
// dedicated API.
func ArtistCover(artist string) (Result, error) {
    url := APIUrl + "artist.getinfo&api_key=" + APIKey + "&artist=" + artist

    if APICorrect {
        url += "&autocorrect=1"
    }

    data, err := request(url)
    if err != nil {
        return Result{}, err
    }

    return parseResults(data, "artist")
}

// TrackCover gets the track cover art from the Last.fm database throught it's
// dedicated API.
func TrackCover(track string, artist string) (Result, error) {
    url := APIUrl + "track.getinfo&api_key=" + APIKey + "&artist=" + artist + "&track=" + track

    if APICorrect {
        url += "&autocorrect=1"
    }

    data, err := request(url)
    if err != nil {
        return Result{}, err
    }

    return parseResults(data, "track")
}
