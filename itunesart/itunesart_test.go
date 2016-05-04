package itunesart_test

import (
	"fmt"
	"github.com/piraveen/go-coverart/itunesart"
	"testing"
)

func TestAlbumCover(t *testing.T) {
	results, err := itunesart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is3.mzstatic.com/image/thumb/Music4/v4/38/42/2b/38422b5a-d597-c4ac-5287-be05cd05dc9e/source/100x100bb.jpg
	}
}

func TestTrackCover(t *testing.T) {
	results, err := itunesart.TrackCover("lights", "ellie goulding")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
		// Output: TrackCover http://is3.mzstatic.com/image/thumb/Music/v4/d7/67/6a/d7676ac4-c1a2-d159-70be-08ff99ee99e4/source/100x100bb.jpg
	}
}

func ExampleAlbumCover() {
	results, err := itunesart.AlbumCover("halcyon days", "ellie goulding")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is3.mzstatic.com/image/thumb/Music4/v4/38/42/2b/38422b5a-d597-c4ac-5287-be05cd05dc9e/source/100x100bb.jpg
	}
}

func ExampleTrackCover() {
	results, err := itunesart.TrackCover("lights", "ellie goulding")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
		// Output: TrackCover http://is3.mzstatic.com/image/thumb/Music/v4/d7/67/6a/d7676ac4-c1a2-d159-70be-08ff99ee99e4/source/100x100bb.jpg
	}
}
