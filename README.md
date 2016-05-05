# go-coverart
[![Build Status](https://travis-ci.org/piraveen/go-coverart.svg?branch=master)](https://travis-ci.org/piraveen/go-coverart)
[![GoDoc](https://godoc.org/github.com/piraveen/go-coverart?status.svg)](https://godoc.org/github.com/piraveen/go-coverart)

A simple Go package to get a track, artist or album artwork art from external APIs like
[Last.fm](http://www.last.fm), [Spotify](https://www.spotify.com),
[Itunes Search](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/),
etc...

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
Full package
```bash
go get -u github.com/piraveen/go-coverart
```

## Commands
A full Go documentation is available for this package [here](https://godoc.org/github.com/piraveen/go-coverart).

- Importing
```go
import "github.com/piraveen/go-coverart"
```

- Setup Last.fm
```go
lastfm, err := coverart.LastFm("LASTFM_APIKEY")
```
Then follow the [Last.fm Helper methods](https://github.com/piraveen/go-coverart/blob/master/README_LASTFM.md)

- Setup Itunes
```go
itunes := coverart.Itunes()
```
Then follow the [Itunes Helper methods](https://github.com/piraveen/go-coverart/blob/master/README_ITUNES.md)

- Setup Spotify
```go
spotify := coverart.Spotify()
```
Then follow the [Spotify Helper methods](https://github.com/piraveen/go-coverart/blob/master/README_SPOTIFY.md)

#### Examples
You can get some sample code for testing from [this](https://github.com/piraveen/go-coverart/blob/master/coverart_test.go) file.

## Documentation
You can read the package in [documentation](https://godoc.org/github.com/piraveen/go-coverart) details in [Godoc](godoc.org).

## Feedback
If you have any suggestions or improvements, please do open an issue [here](https://github.com/piraveen/go-coverart/issues).

Cheers :)
