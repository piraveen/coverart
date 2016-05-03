# CoverArt
[![GoDoc](https://godoc.org/github.com/piraveen/coverart?status.svg)](https://godoc.org/github.com/piraveen/coverart)

A simple Go package to get a track, artist or album artwork art from external APIs like
[Last.fm](http://www.last.fm), [Spotify](https://www.spotify.com),
[Itunes Search](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/),
etc...

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
Full package
```
go get github.com/piraveen/coverart
```
Last.fm package only
```
go get github.com/piraveen/coverart/lastfmart
```
Itunes package only
```
go get github.com/piraveen/coverart/itunesart
```

## Commands
### Using Last.fm API
Read more about the [last.fm API](http://last.fm).

- Importing
```
import "github.com/piraveen/coverart/lastfmart"
```
- Configuration
```
lastfmart.Configure("LASTFM_APIKEY")
```
- Checking if API Key is set
```
lastfmart.CheckAPIKey()
```
- Enable Auto correction
```
lastfmart.AutoCorrect(true)
```
- Get Album Artwork
```
lastfmart.AlbumCover("album name", "artist name")
```
- Get Artist Artwork
```
lastfmart.ArtistCover("artist name")
```
- Get Track Artwork
```
lastfmart.TrackArt("track name", "artist name")
```

Sample code for testing Last.fm API:
```go
package main

import (
    "fmt"
    "github.com/piraveen/coverart/lastfmart"
)

func main() {
    // The API Keys can be defined in your code itself, however I recommend
    // loading them through an environment variable like this:
    apiKey := os.Getenv("LASTFM_APIKEY")
    lastfmart.Configure(apiKey)

    if !lastfmart.CheckAPIKey() {
        // Abort action
        fmt.Printf("No API Key or incorrectly set\n")
        // Output: No API Key or incorrectly set
        return;
    }

    results, err := lastfmart.ArtistCover("ellie goulding")
    if err == nil {
        fmt.Printf("ArtistCover %v\n", results.Default)
        // Output: ArtistCover http://img2-ak.lst.fm/i/u/arQ/eb410194931c9427e2240023426be62b.png
    }
}
```

### Using Itunes Search API
Read more about the [Itunes Search API](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/).

- Importing
```
import "github.com/piraveen/coverart/itunesart"
```
- Get Album Artwork
```
itunesart.AlbumCover("album name", "artist name")
```
- Get Track Artwork
```
itunesart.TrackArt("track name", "artist name")
```

Sample code for testing Itunes Search API:
```go
package main

import (
    "fmt"
    "github.com/piraveen/coverart/itunesart"
)

func main() {
    results, err := itunesart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is3.mzstatic.com/image/thumb/Music4/v4/38/42/2b/38422b5a-d597-c4ac-5287-be05cd05dc9e/source/100x100bb.jpg
	}
}
```

## Documentation
Please read the package [documentation](https://godoc.org/github.com/piraveen/coverart) details here.
