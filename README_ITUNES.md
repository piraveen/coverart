# go-coverart/itunesart
[![Build Status](https://travis-ci.org/piraveen/go-coverart.svg?branch=master)](https://travis-ci.org/piraveen/go-coverart)
[![GoDoc](https://godoc.org/github.com/piraveen/go-coverart?status.svg)](https://godoc.org/github.com/piraveen/go-coverart/itunesart)

A simple Go package to get a track or album artwork artwork from [Itunes Search](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/).

Read more about the [Itunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/).

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
```bash
go get -u github.com/piraveen/go-coverart/itunesart
```

### Commands
- Importing
```go
import "github.com/piraveen/go-coverart/itunesart"
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
You can get some sample code for testing from [this](https://github.com/piraveen/go-coverart/blob/master/itunesart/itunesart_test.go) file.

## Documentation
You can read the package [documentation](https://godoc.org/github.com/piraveen/go-coverart/itunesart) details in [Godoc](godoc.org).

## Feedback
If you have any suggestions or improvements, please do open an issue [here](https://github.com/piraveen/go-coverart/issues).

Cheers :)
