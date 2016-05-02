package coverart

import (
    "errors"
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

type Image struct {
    Size    string      `json:size`
    Url     string      `json:"#text"`
}

type Album struct {
    Name    string      `json:name`
    Image   []Image     `json:image`
}

type Track struct {
    Name    string      `json:name`
    Album   *Album      `json:album`
}

type Artist struct {
    Name    string      `json:name`
    Image   []Image     `json:image`
}

type HttpResponse struct {
    Album   *Album      `json:album`
    Artist  *Artist     `json:artist`
    Track   *Track      `json:track`
}

type HttpError struct {
    Error   *int        `json:error`
    Message *string     `json:message`
}

// Activate auto correct for spelling mistakes
// Note: Result may not be as expected
func AutoCorrect(act bool) {
    APICorrect = true
}

// Configure API Key and other required fields
func Configure(key string) {
    APIKey = key
    APICorrect = false
}

// Build all the images into size typed object for easy access
// { Result.SizeName }
// e.g: Result.Small would return the url for a small size cover art
func buildResult(images []Image) (Result, error) {
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

    return res, nil
}

// Parse http response and build results based on requested type
// parse { album, artist, track }
func parseResults(data []byte, parse string) (Result, error) {
    resp := HttpResponse{}

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
    resErr := HttpError{}
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

// Get Album Cover art
// Parametes: Album name, Artist name
// Returns: Result, error object
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

// Get Artist Cover art
// Parametes: Artist name
// Returns: Result, error object
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

// Get Track Cover art
// Parametes: Track name, Artist name
// Returns: Result, error object
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
