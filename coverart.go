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
	"github.com/piraveen/coverart/lastfmart"
	"github.com/piraveen/coverart/itunesart"
)

// The ItunesArt represents the specific helper methods of the itunesart package
type ItunesArt struct {
	TrackCover func(track string, artist string) (itunesart.Result, error)
	AlbumCover func(album string, artist string) (itunesart.Result, error)
}

// The LastFmArt represents the specific helper methods of the lastfmart package
type LastFmArt struct {
	CheckAPIKey func() error
	AutoCorrect func(v bool)
	SetAPIKey	func(k string)
	TrackCover  func(track string, artist string) (lastfmart.Result, error)
	AlbumCover  func(album string, artist string) (lastfmart.Result, error)
	ArtistCover func(artist string) (lastfmart.Result, error)
}

// LastFm configures and returns all the exported methods of the package lastfmart
func LastFm(apiKey string) (LastFmArt, error) {
	lastfmart.Configure(apiKey)

	if err := lastfmart.CheckAPIKey(); err != nil {
		return LastFmArt{}, err
	}

	return LastFmArt{
		lastfmart.CheckAPIKey,
		lastfmart.AutoCorrect,
		lastfmart.SetAPIKey,
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
