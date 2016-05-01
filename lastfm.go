package coverart

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

var verbose bool
var APIKey string
var APIUrl = "http://ws.audioscrobbler.com/2.0/?format=json&method="

func SetApiToken(key string) {
    APIKey = key
}

type IResults struct {
    Size    string
    Url     string
}

func Verbose(v bool) {
    verbose = v
}

func Configure(key string) {
    APIKey = key
}

type IAlbumResponse struct {
    Name    string `json:name`
}

type IHttpResponse struct {
    Album   IAlbumResponse  `json:album`
}

func buildResults(data []byte) []IResults {
    var Results []IResults

    m := map[string]IHttpResponse{}
    err := json.Unmarshal(data, &m)
    if err != nil {
        panic(err)
    }
    fmt.Println(m)
    return Results
}

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

    // fmt.Printf("Results %v\n", string(body))

    buildResults(body)
    return body, err
}

func AlbumCover(album string, artist string) ([]IResults, error) {
    var Results []IResults

    url := APIUrl + "album.getinfo&api_key=" + APIKey
    url += "&album=" + album
    url += "&artist=" + artist

    if verbose {
        fmt.Printf("Fetching album cover for %s...\n", album)
    }

    _, err := request(url)
    return Results, err
}
