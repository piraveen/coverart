# CoverArt
[![GoDoc](https://godoc.org/github.com/piraveen/coverart?status.svg)](https://godoc.org/github.com/piraveen/coverart)

A simple Go package to get a track, artist or album artwork art from external APIs like
[Last.fm](http://www.last.fm), [Spotify](https://www.spotify.com),
[Itunes Search](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/),
etc...

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
Full package
```bash
go get github.com/piraveen/coverart
```
Last.fm package only
```bash
go get github.com/piraveen/coverart/lastfmart
```
Itunes package only
```bash
go get github.com/piraveen/coverart/itunesart
```

## Commands
### Using All
A full Go documentation is available for this package [here](https://godoc.org/github.com/piraveen/coverart).

- Importing
```go
import "github.com/piraveen/coverart"
```

- Setup Last.fm
```go
lastfm, err := coverart.LastFm("LASTFM_APIKEY")
```
Then follow the [Last.fm Helper methods](#using-lastfm-api)

- Setup Itunes
```go
itunes := coverart.Itunes()
```
Then follow the [Itunes Helper methods](#using-itunes-search-api)

#### Examples
You can get some sample code for testing from [this](https://github.com/piraveen/coverart/blob/master/coverart_test.go) file.

### Using Last.fm API
Read more about the [last.fm API](http://last.fm).
A full Go documentation is available for this package [here](https://godoc.org/github.com/piraveen/coverart/lastfmart).

- Importing
```go
import "github.com/piraveen/coverart/lastfmart"
```
- Configuration
```go
lastfmart.Configure("LASTFM_APIKEY")
```
- Checking if API Key is set
```go
lastfmart.CheckAPIKey()
```
- Enable Auto correction
```go
lastfmart.AutoCorrect(true)
```
- Get Album Artwork
```go
result, err = lastfmart.AlbumCover("album name", "artist name")
```
- Get Artist Artwork
```go
result, err = lastfmart.ArtistCover("artist name")
```
- Get Track Artwork
```go
result, err = lastfmart.TrackArt("track name", "artist name")
```
#### Examples
You can get some sample code for testing from [this](https://github.com/piraveen/coverart/blob/master/lastfmart/lastfmart_test.go) file.

### Using Itunes Search API
Read more about the [Itunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/).
A full Go documentation is available for this package [here](https://godoc.org/github.com/piraveen/coverart/itunesart).

- Importing
```go
import "github.com/piraveen/coverart/itunesart"
```
- Get Album Artwork
```go
result, err = itunesart.AlbumCover("album name", "artist name")
```
- Get Track Artwork
```go
result, err = itunesart.TrackArt("track name", "artist name")
```
#### Examples
You can get some sample code for testing from [this](https://github.com/piraveen/coverart/blob/master/itunesart/itunesart_test.go) file.

## Documentation
You can read the package in [documentation](https://godoc.org/github.com/piraveen/coverart) details in [Godoc](godoc.org).
