// Package coverart provides a helper that imports the itunesart and lastfmart packages.
// Then returns an access interface for each individual services to get album,
// artist or track artworks.
//
// Note: This is a lazy package to load all the sub-service packages concurrently.
//
// Concerned packages:
//
// "github.com/piraveen/coverart/itunesart"
// "github.com/piraveen/coverart/lastfmart"
package coverart

import (
    "errors"
    "github.com/piraveen/coverartbk/itunesart"
    "github.com/piraveen/coverart/lastfmart"
)

// The ItunesArt represents the specific helper methods of the itunesart package
type ItunesArt struct {
    TrackCover  func(track string, artist string) (itunesart.Result, error)
    AlbumCover  func(album string, artist string) (itunesart.Result, error)
}

// The LastFmArt represents the specific helper methods of the lastfmart package
type LastFmArt struct {
    CheckAPIKey func() bool
    AutoCorrect func(v bool)
    TrackCover  func(track string, artist string) (lastfmart.Result, error)
    AlbumCover  func(album string, artist string) (lastfmart.Result, error)
    ArtistCover func(artist string) (lastfmart.Result, error)
}


// LastFm configures and returns all the exported methods of the package lastfmart
func LastFm(apiKey string) (LastFmArt, error) {
    lastfmart.Configure(apiKey)

    if !lastfmart.CheckAPIKey() {
        return LastFmArt{}, errors.New("API Key is not set")
    }

    return LastFmArt{
        lastfmart.CheckAPIKey,
        lastfmart.AutoCorrect,
        lastfmart.TrackCover,
        lastfmart.AlbumCover,
        lastfmart.ArtistCover,
    }, nil
}

// Itunes configures and returns all the exported methods of the package itunesart
func Itunes() ItunesArt {
    return ItunesArt{
        itunesart.TrackCover,
        itunesart.AlbumCover,
    }
}
