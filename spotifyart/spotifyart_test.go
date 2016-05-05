package spotifyart_test

import (
	"fmt"
	"github.com/piraveen/go-coverart/spotifyart"
	"os"
	"testing"
)

func TestAlbumCover(t *testing.T) {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	c := os.Getenv("SPOTIFY_CLIENTID")
	s := os.Getenv("SPOTIFY_CLIENTSECRET")


	if err := spotifyart.Configure(c, s); err != nil {
		// Abort action
		fmt.Println(err)
		fmt.Printf("No API Key or incorrectly set\n")
		// Output: No API Key or incorrectly set
		return
	}

	results, err := spotifyart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// : AlbumCover http://img2-ak.lst.fm/i/u/arQ/e0a131728ae7438d8b7adf87ae323b46.png
	} else {
		fmt.Println(err)
	}
}

// func TestArtistCover(t *testing.T) {
// 	// The API Keys can be defined in your code itself, however I recommend
// 	// loading them through an environment variable like this:
// 	c := os.Getenv("SPOTIFY_CLIENTID")
// s := os.Getenv("SPOTIFY_CLIENTSECRET")
// 	spotifyart.Configure(c, s)
//
// 	if err := spotifyart.CheckAPIKey(); err != nil {
// 		// Abort action
// 		fmt.Printf("No API Key or incorrectly set\n")
// 		// Output: No API Key or incorrectly set
// 		return
// 	}
//
// 	results, err := spotifyart.ArtistCover("ellie goulding")
// 	if err == nil {
// 		fmt.Printf("ArtistCover %v\n", results.Default)
// 		// Output: ArtistCover http://img2-ak.lst.fm/i/u/arQ/eb410194931c9427e2240023426be62b.png
// 	}
// }
//
// func TestTrackCover(t *testing.T) {
// 	// The API Keys can be defined in your code itself, however I recommend
// 	// loading them through an environment variable like this:
// c := os.Getenv("SPOTIFY_CLIENTID")
// s := os.Getenv("SPOTIFY_CLIENTSECRET")
// 	spotifyart.Configure(c, s)
//
// 	if err := spotifyart.CheckAPIKey(); err != nil {
// 		// Abort action
// 		fmt.Printf("No API Key or incorrectly set\n")
// 		// Output: No API Key or incorrectly set
// 		return
// 	}
//
// 	results, err := spotifyart.TrackCover("lights", "ellie goulding")
// 	if err == nil {
// 		fmt.Printf("TrackCover %v\n", results.Default)
// 		// Output: TrackCover http://img2-ak.lst.fm/i/u/34s/24029dde6b1345dea4aadfcfe4126b9c.png
// 	}
// }
//
// func ExampleAlbumCover() {
// 	// The API Keys can be defined in your code itself, however I recommend
// 	// loading them through an environment variable like this:
// 	// c := os.Getenv("SPOTIFY_CLIENTID"); s := os.Getenv("SPOTIFY_CLIENTSECRET")
// 	spotifyart.Configure("SPOTIFY_CLIENTID", "SPOTIFY_CLIENTSECRET")
//
// 	if err := spotifyart.CheckAPIKey(); err != nil {
// 		// Abort action
// 		return
// 	}
//
// 	results, err := spotifyart.AlbumCover("halcyon days", "ellie goulding")
// 	if err == nil {
// 		fmt.Printf("AlbumCover %v\n", results.Default)
// 	}
// }
//
// func ExampleArtistCover() {
// 	// The API Keys can be defined in your code itself, however I recommend
// 	// loading them through an environment variable like this:
// 	// c := os.Getenv("SPOTIFY_CLIENTID"); s := os.Getenv("SPOTIFY_CLIENTSECRET")
// 	spotifyart.Configure("SPOTIFY_CLIENTID", "SPOTIFY_CLIENTSECRET")
//
// 	if err := spotifyart.CheckAPIKey(); err != nil {
// 		// Abort action
// 		return
// 	}
//
// 	results, err := spotifyart.ArtistCover("ellie goulding")
// 	if err == nil {
// 		fmt.Printf("ArtistCover %v\n", results.Default)
// 	}
// }
//
// func ExampleTrackCover() {
// 	// The API Keys can be defined in your code itself, however I recommend
// 	// loading them through an environment variable like this:
// 	// c := os.Getenv("SPOTIFY_CLIENTID"); s := os.Getenv("SPOTIFY_CLIENTSECRET")
// 	spotifyart.Configure("SPOTIFY_CLIENTID", "SPOTIFY_CLIENTSECRET")
//
// 	if err := spotifyart.CheckAPIKey(); err != nil {
// 		// Abort action
// 		return
// 	}
//
// 	results, err := spotifyart.TrackCover("lights", "ellie goulding")
// 	if err == nil {
// 		fmt.Printf("TrackCover %v\n", results.Default)
// 	}
// }
