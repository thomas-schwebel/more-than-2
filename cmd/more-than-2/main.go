package main

import (
	"fmt"
	"log"
	"more-than-2/internal/spotify"
	"net/http"
)

const (
	accessToken       = "REPLACE_WITH_VALID_SPOTIFY_ACCESS_TOKEN"
	owner             = "REPLACE_WITH_SPOTIFY_ID"
	spotifyApiBaseUrl = "https://api.spotify.com/v1/"
)

func main() {
	// Todo: treat likes as a playlist to process as well

	client := spotify.HttpClient{
		Client:      *http.DefaultClient,
		BaseUrl:     spotifyApiBaseUrl,
		AccessToken: accessToken,
		Owner:       owner,
	}

	sorted, err := client.CountTrackByArtist()
	if err != nil {
		log.Fatal(err)
	}

	maxCount := sorted[0].Count

	for _, ac := range sorted {
		if ac.Count == 1 {
			break // more-than-2 :) single tracks are irrelevant
		}

		if maxCount == sorted[0].Count || maxCount > ac.Count {
			maxCount = ac.Count
			fmt.Printf("\n******* %d MATCHES *******\n", maxCount)
		}

		fmt.Printf("- %s\n", ac.Artist)
	}
}
