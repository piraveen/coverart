package spotifyart_test

import (
	"fmt"
	"github.com/piraveen/go-coverart/spotifyart"
	"testing"
)

func TestAlbumCover(t *testing.T) {
	results, err := spotifyart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover https://i.scdn.co/image/c649d891ee6e0b86bf460bca264bd66715bd87f4
	}
}

func TestArtistCover(t *testing.T) {
	results, err := spotifyart.ArtistCover("ellie goulding", "pop", "metropopolis")
	if err == nil {
		fmt.Printf("ArtistCover %v\n", results.Default)
		// Output: ArtistCover https://i.scdn.co/image/b72e148adf8cec8bf91784bee05d836858546367
	}
}

func TestTrackCover(t *testing.T) {
	results, err := spotifyart.TrackCover("lights", "ellie goulding")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
		// Output: TrackCover https://i.scdn.co/image/5031974d62fe1ef31b7678343e3ad4d11d922131
	}
}

func ExampleAlbumCover() {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// c := os.Getenv("SPOTIFY_CLIENTID")
	// s := os.Getenv("SPOTIFY_CLIENTSECRET")

	// Note: Providing Spotify Client Id and Client Secret is optional, but it
	// would help you increase the Spotify rate limit for requests
	// if err := spotifyart.Configure(c, s); err != nil {
	// 	// Abort action
	// 	return
	// }

	results, err := spotifyart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
	}
}

func ExampleArtistCover() {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// c := os.Getenv("SPOTIFY_CLIENTID")
	// s := os.Getenv("SPOTIFY_CLIENTSECRET")

	// Note: Providing Spotify Client Id and Client Secret is optional, but it
	// would help you increase the Spotify rate limit for requests
	// if err := spotifyart.Configure(c, s); err != nil {
	// 	// Abort action
	// 	return
	// }

	results, err := spotifyart.ArtistCover("ellie goulding", "pop", "metropopolis")
	if err == nil {
		fmt.Printf("ArtistCover %v\n", results.Default)
	}
}

func ExampleTrackCover() {
	// The API Keys can be defined in your code itself, however I recommend
	// loading them through an environment variable like this:
	// c := os.Getenv("SPOTIFY_CLIENTID")
	// s := os.Getenv("SPOTIFY_CLIENTSECRET")

	// Note: Providing Spotify Client Id and Client Secret is optional, but it
	// would help you increase the Spotify rate limit for requests
	// if err := spotifyart.Configure(c, s); err != nil {
	// 	// Abort action
	// 	return
	// }
	//

	results, err := spotifyart.TrackCover("lights", "ellie goulding")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
	}
}
