// Package spotifyart provides few helper methods to get album, artist or track
// artworks from the Spotify API
package spotifyart

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"errors"
	"fmt"
	"os"
)

var clId, clSecret string

const apiUrl = "https://api.spotify.com/v1"
const apiUrlTrack = apiUrl + "/search?type=track&limit=1&q="
const apiUrlAlbum = apiUrl + "/search?type=album&limit=1&q="
const apiUrlArtist = apiUrl + "/search?type=artist&limit=1&q="
const apiUrlToken = "https://accounts.spotify.com/api/token"

// The Result represents the specific size of artworks and contains the url of
// each size of artwork returned by the Spotify API
type Result struct {
	Large   string
	Medium  string
	Small   string
	Default string
}

type image struct {
	Width  *int `json:width`
	Height *int `json:height`
	Url    string `json:url`
}

type item struct {
	Type string `json:type`
	Name string `json:name`
	Images []image `json:images`
}

type items struct {
	Items	[]item `json:items`
}

type httpSearch struct {
	Albums  *items  `json:albums`
	Tracks  *items  `json:tracks`
	Artists *items `json:artists`
}

type httpErrorDetails struct {
	Status  int    `json:status`
	Message string `json:message`
}

type httpError struct {
	Error *httpErrorDetails `json:error`
}

type httpTokenError struct {
	Error       *string `json:error`
	Description *string `json:"errorerror_description,omitempty"`
}

type httpToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// SetCredentials provides a method to set the  the Spotify Client Id and
// Client Secret. This action will result in with a call to the Spotify API
// to get an access token. The access token will allow you to have a higher
// limit rate than unauthorized requests
func SetCredentials(i string, s string) error {
	clId, clSecret = i, s
	return CheckCredentials()
}

// Configure is optional, you can use it to set the Spotify Client Id and
// Client Secret. This action will result in with a call to the Spotify API
// to get an access token. The access token will allow you to have a higher
// limit rate than unauthorized requests
func Configure(i string, s string) error {
	return SetCredentials(i, s)
}

// CheckCredentials provides a simple method to verify if the API Key has been set and
// if it is valid
func CheckCredentials() error {
	if len(clId) == 0 || len(clSecret) == 0 {
		return nil
	}

	byteCreds := []byte(clId + ":" + clSecret)
	encodedCres := base64.StdEncoding.EncodeToString(byteCreds)
	return getAccessToken(encodedCres)
}

// Used to get an access token from the Spotify API
func getAccessToken(ec string) error {
	data := url.Values{"grant_type": {"client_credentials"}}
	req, _ := http.NewRequest("POST", apiUrlToken, bytes.NewBufferString(data.Encode()))

	req.Header.Add("Authorization", "Basic "+ec)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resToken, err := requestToken(req)
	if err != nil {
		return err
	}

	setToken(resToken.AccessToken)
	return nil
}

func setToken(t string) {
	os.Setenv("SPOTIFY_ACCESSTOKEN", t)
}

func getToken() string {
	return os.Getenv("SPOTIFY_ACCESSTOKEN")
}

// Executes a manually created request with http.Client service
func requestToken(req *http.Request) (*httpToken, error) {
	resErr := httpTokenError{}
	resToken := httpToken{}
	client := &http.Client{}

	resp, err := client.Do(req)
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
		return nil, errors.New(*resErr.Description)
	}

	err = json.Unmarshal(body, &resToken)
	if err != nil {
		return nil, err
	}

	return &resToken, nil
}

// Build all the artwork into size typed object for easy access
// { Result.SizeName }
// e.g: Result.Small would return the url for a small size artwork
func buildResult(sItem item) (Result, error) {
	res := Result{}
	sizes := []string{ "large", "medium", "small", }

	if len(sItem.Images) == 0 {
		return res, errors.New("No image was found")
	}

	for key, value := range sItem.Images {
		switch sizes[key] {
		default:
			res.Default = value.Url
		case "small":
			res.Small = value.Url
		case "medium":
			res.Medium = value.Url
		case "large":
			res.Large = value.Url
			res.Default = value.Url
		}
	}

	return res, nil
}

// Parse http response and build results based on requested type
// parse { album, artist, track }
func parseResults(data []byte, parse string) (Result, error) {
	resp := httpSearch{}

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Result{}, err
	}

	switch parse {
	default:
		return Result{}, errors.New("No image was found")
	case "album":
		if resp.Albums != nil && len(resp.Albums.Items) > 0 {
			return buildResult(resp.Albums.Items[0])
		}
	case "track":
		if resp.Tracks != nil && len(resp.Tracks.Items) > 0 {
			return buildResult(resp.Tracks.Items[0])
		}
	case "artist":
		if resp.Artists != nil && len(resp.Artists.Items) > 0 {
			return buildResult(resp.Artists.Items[0])
		}
	}

	return Result{}, errors.New("No image was found")
}

// Executes an http request and returns error or response body
func request(url string) ([]byte, error) {
	resErr := httpError{}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	if len(getToken()) > 0 {
		req.Header.Add("Authorization", "Bearer " + getToken())
	}

	resp, err := client.Do(req)
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
		err = errors.New(resErr.Error.Message)
	}

	return body, err
}

// AlbumCover gets the album artwork from the Spotify database through out it's
// dedicated API.
func AlbumCover(album string, artist string) (Result, error) {
	Url := apiUrlAlbum + "album:" + url.QueryEscape(album + " ")
	Url += "artist:" + url.QueryEscape(artist)

	data, err := request(Url)
	if err != nil {
		return Result{}, err
	}

	return parseResults(data, "album")
}

// ArtistCover gets the artist artwork from the Spotify database through out it's
// dedicated API.
func ArtistCover(artist string) (Result, error) {
	Url := apiUrlArtist + url.QueryEscape(artist)

	data, err := request(Url)
	if err != nil {
		return Result{}, err
	}

	return parseResults(data, "artist")
}
//
// // TrackCover gets the track artwork from the Spotify database through out it's
// // dedicated API.
// func TrackCover(track string, artist string) (Result, error) {
// 	Url := apiUrl + "track.getinfo&api_key=" + apiKey + "&artist="
// 	Url += url.QueryEscape(artist) + "&track=" + url.QueryEscape(track)
//
// 	// if apiCorrect {
// 	// 	Url += "&autocorrect=1"
// 	// }
//
// 	data, err := request(Url)
// 	if err != nil {
// 		return Result{}, err
// 	}
//
// 	return parseResults(data, "track")
// }

func testParse() {
	item := []byte(`{
  "artists": {
    "href": "https://api.spotify.com/v1/search?query=ellie+goulding&offset=0&limit=1&type=artist",
    "items": [
      {
        "external_urls": {
          "spotify": "https://open.spotify.com/artist/0X2BH1fck6amBIoJhDVmmJ"
        },
        "followers": {
          "href": null,
          "total": 2694010
        },
        "genres": [
          "metropopolis",
          "pop",
          "synthpop"
        ],
        "href": "https://api.spotify.com/v1/artists/0X2BH1fck6amBIoJhDVmmJ",
        "id": "0X2BH1fck6amBIoJhDVmmJ",
        "images": [
          {
            "height": 1000,
            "url": "https://i.scdn.co/image/cdfa418a53726ce2255cd543d4be873af49b5499",
            "width": 1000
          },
          {
            "height": 640,
            "url": "https://i.scdn.co/image/3820c08b5be577f9dd746bedf80d6a85640d7a30",
            "width": 640
          },
          {
            "height": 200,
            "url": "https://i.scdn.co/image/5987a65ac367986e41a7ec650c8d0eecb3963309",
            "width": 200
          },
          {
            "height": 64,
            "url": "https://i.scdn.co/image/b72e148adf8cec8bf91784bee05d836858546367",
            "width": 64
          }
        ],
        "name": "Ellie Goulding",
        "popularity": 86,
        "type": "artist",
        "uri": "spotify:artist:0X2BH1fck6amBIoJhDVmmJ"
      }
    ],
    "limit": 1,
    "next": "https://api.spotify.com/v1/search?query=ellie+goulding&offset=1&limit=1&type=artist",
    "offset": 0,
    "previous": null,
    "total": 5
  }
}`)

	fmt.Println(item)
}
