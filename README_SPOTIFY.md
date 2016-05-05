# go-coverart/spotifyart
[![Build Status](https://travis-ci.org/piraveen/go-coverart.svg?branch=master)](https://travis-ci.org/piraveen/go-coverart)
[![GoDoc](https://godoc.org/github.com/piraveen/go-coverart?status.svg)](https://godoc.org/github.com/piraveen/go-coverart/spotifyart)

A simple Go package to get a track, artist or album artwork artwork from [Spotify](https://www.spotify.com).

Read more about the [Spotify API](https://developer.spotify.com/web-api/).

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
```bash
go get -u github.com/piraveen/go-coverart/spotifyart
```

### Commands
- Importing
```go
import "github.com/piraveen/go-coverart/spotifyart"
```
- Configuration (optional)

    Setting up a Spotify Client Id and Client Secret will allow you to make more
    API requests by increasing the requests rate limit with an access token.

    - Configure
    ```go
    err := spotifyart.Configure("SPOTIFY_CLIENTID", "SPOTIFY_CLIENTSECRET")
    ```
    - Checking if Client Id and Client Secret are set
    ```go
    status := spotifyart.CheckCredentials()
    ```
    - Get or refresh Access (Note: Client Id and Client Secret must be set)
    ```go
    err := spotifyart.GetAccessToken()
    ```

- Get Album Artwork
```go
result, err = spotifyart.AlbumCover("album name", "optional name")
```
- Get Artist Artwork
```go
result, err = spotifyart.ArtistCover("artist name", "optional genre")
```
- Get Track Artwork
```go
result, err = spotifyart.TrackArt("track name", "optional name")
```

#### Examples
You can get some sample code for testing from [this](https://github.com/piraveen/go-coverart/blob/master/spotifyart/spotifyart_test.go) file.

## Documentation
You can read the package [documentation](https://godoc.org/github.com/piraveen/go-coverart/spotifyart) details in [Godoc](godoc.org).

## Feedback
If you have any suggestions or improvements, please do open an issue [here](https://github.com/piraveen/go-coverart/issues).

Cheers :)
