// Package coverart provides a helper that imports the spotifyart, itunesart and
// lastfmart packages. Then returns an access interface for each individual
// services to get album, artist or track artworks.
//
// Note: This is a lazy package to load all the sub-service packages concurrently.
//
// Concerned packages:
//
// "github.com/piraveen/go-coverart/itunesart"
// "github.com/piraveen/go-coverart/lastfmart"
// "github.com/piraveen/go-coverart/spotifyart"
package coverart

import (
	"github.com/piraveen/go-coverart/itunesart"
	"github.com/piraveen/go-coverart/lastfmart"
	"github.com/piraveen/go-coverart/spotifyart"
)

// The ItunesArt represents the specific helper methods of the itunesart package
type ItunesArt struct {
	Result     itunesart.Result
	TrackCover func(track string, artist string) (itunesart.Result, error)
	AlbumCover func(album string, artist string) (itunesart.Result, error)
}

// The LastFmArt represents the specific helper methods of the lastfmart package
type LastFmArt struct {
	Result      lastfmart.Result
	CheckAPIKey func() error
	AutoCorrect func(v bool)
	SetAPIKey   func(k string)
	TrackCover  func(track string, artist string) (lastfmart.Result, error)
	AlbumCover  func(album string, artist string) (lastfmart.Result, error)
	ArtistCover func(artist string) (lastfmart.Result, error)
}

// The SpotifyArt represents the specific helper methods of the spotifyart package
type SpotifyArt struct {
	Result           spotifyart.Result
	CheckCredentials func() bool
	GetAccessToken   func() error
	Configure        func(clientId string, clientSecret string) error
	TrackCover       func(track string, artists ...string) (spotifyart.Result, error)
	AlbumCover       func(album string, artists ...string) (spotifyart.Result, error)
	ArtistCover      func(artist string, genres ...string) (spotifyart.Result, error)
}

// LastFm configures and returns all the exported methods of the package lastfmart
func LastFm(apiKey string) (LastFmArt, error) {
	lastfmart.Configure(apiKey)

	if err := lastfmart.CheckAPIKey(); err != nil {
		return LastFmArt{}, err
	}

	return LastFmArt{
		lastfmart.Result{},
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
		itunesart.Result{},
		itunesart.TrackCover,
		itunesart.AlbumCover,
	}
}

// Spotify configures and returns all the exported methods of the package spotifyart
func Spotify() SpotifyArt {
	return SpotifyArt{
		spotifyart.Result{},
		spotifyart.CheckCredentials,
		spotifyart.GetAccessToken,
		spotifyart.Configure,
		spotifyart.TrackCover,
		spotifyart.AlbumCover,
		spotifyart.ArtistCover,
	}
}
