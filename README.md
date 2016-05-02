# coverart
Go package to get a track / album covert art from external APIs like
[Last.fm](http://www.last.fm), [Spotify](https://www.spotify.com),
[Itunes](https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/),
etc...

<strong>Important: This package is strictly for a non-commercial use.</strong>

## Installing package
```
go get github.com/piraveen/covertart
```

## Commands
### General
All the response for the requests below will be sent in this format:
```go
Result{
    Small: "http://url"
    Medium: "http://url"
    Large: "http://url"
    ExtraLarge: "http://url"
    Mega: "http://url"
    Default: "http://url"
}
```
Note: Only the `Default` field will be always present. And if no images were found
a go `error` object will be returned.

Please do read the EXAMPLE.md for further details.

### Last.fm API
Read more about the [last.fm](http://last.fm) API [here](http://www.last.fm/api)

- Configuration
```
lastfmart.Configure("LASTFM_API_KEY")
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
