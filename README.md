# coverart
Go package to get a track / album covert art from external APIs like
[Last.fm](http://www.last.fm), [Spotify](https://www.spotify.com),
[Itunes](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/),
etc...

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Install
Full package
```
go get -d github.com/piraveen/coverart
```
Last.fm package only
```
go get github.com/piraveen/coverart/lastfmart
```

## Commands
### Using Last.fm API
Read more about the [last.fm](http://last.fm) API [here](http://www.last.fm/api).

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
- Get Album Art
```
lastfmart.AlbumCover("album name", "artist name")
```
- Get Artist Art
```
lastfmart.ArtistCover("artist name")
```
- Get Track Art
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

## Documentation
Please read the package [documentation](https://godoc.org/github.com/piraveen/coverart) details here.
