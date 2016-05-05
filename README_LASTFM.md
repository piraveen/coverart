# go-coverart/lastfmart
[![Build Status](https://travis-ci.org/piraveen/go-coverart.svg?branch=master)](https://travis-ci.org/piraveen/go-coverart)
[![GoDoc](https://godoc.org/github.com/piraveen/go-coverart?status.svg)](https://godoc.org/github.com/piraveen/go-coverart/lastfmart)

A simple Go package to get a track, artist or album artwork from [Last.fm](http://www.last.fm)

Read more about the [last.fm API](http://last.fm).

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
```bash
go get -u github.com/piraveen/go-coverart/lastfmart
```

### Commands
- Importing
```go
import "github.com/piraveen/go-coverart/lastfmart"
```
- Configuration
```go
lastfmart.Configure("LASTFM_APIKEY")
```
- Checking if API Key is set and it's validity
```go
err := lastfmart.CheckAPIKey()
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
You can get some sample code for testing from [this](https://github.com/piraveen/go-coverart/blob/master/lastfmart/lastfmart_test.go) file.

## Documentation
You can read the package [documentation](https://godoc.org/github.com/piraveen/go-coverart/lastfmart) details in [Godoc](godoc.org).

## Feedback
If you have any suggestions or improvements, please do open an issue [here](https://github.com/piraveen/go-coverart/issues).

Cheers :)
