package itunesart_test

import (
	"fmt"
	"github.com/piraveen/go-coverart/itunesart"
	"testing"
)

func TestAlbumCover(t *testing.T) {
	results, err := itunesart.AlbumCover("unapologetic", "rihanna")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is4.mzstatic.com/image/thumb/Music/v4/7a/d3/8d/7ad38df1-c8da-f278-af55-e346a073451a/source/100x100bb.jpg
	}
}

func TestTrackCover(t *testing.T) {
	results, err := itunesart.TrackCover("stay", "rihanna")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
		// Output: TrackCover http://is4.mzstatic.com/image/thumb/Music/v4/7a/d3/8d/7ad38df1-c8da-f278-af55-e346a073451a/source/100x100bb.jpg
	}
}

func ExampleAlbumCover() {
	results, err := itunesart.AlbumCover("unapologetic", "rihanna")
	if err == nil {
		fmt.Printf("AlbumCover %v\n", results.Default)
		// Output: AlbumCover http://is4.mzstatic.com/image/thumb/Music/v4/7a/d3/8d/7ad38df1-c8da-f278-af55-e346a073451a/source/100x100bb.jpg
	} else {
		fmt.Println("error", err)
	}
}

func ExampleTrackCover() {
	results, err := itunesart.TrackCover("stay", "rihanna")
	if err == nil {
		fmt.Printf("TrackCover %v\n", results.Default)
		// Output: TrackCover http://is4.mzstatic.com/image/thumb/Music/v4/7a/d3/8d/7ad38df1-c8da-f278-af55-e346a073451a/source/100x100bb.jpg
	}
}
