package spotify

import (
	"fmt"
	"net/url"

	"more-than-2/internal/stats"

	"golang.org/x/exp/maps"
)

func (c *HttpClient) CountTrackByArtist() ([]stats.ArtistCount, error) {
	artistTrackCount := map[string]uint{}

	url, err := url.JoinPath(c.BaseUrl, "/me/playlists")
	if err != nil {
		return []stats.ArtistCount{}, err
	}

	fmt.Println("Processing playlists from owner...")

	nextUrl := url
	for nextUrl != "" { // todo gather urls and use goroutines
		playlists, err := GetApiStruct[Playlists](c, nextUrl)
		if err != nil {
			return []stats.ArtistCount{}, err
		}

		playlistCount, err := countForPlaylist(c, playlists.Items)
		if err != nil {
			return []stats.ArtistCount{}, err
		}
		maps.Copy(artistTrackCount, playlistCount)

		nextUrl = playlists.Next
	}

	return stats.SortByCount(artistTrackCount), nil
}

func countForPlaylist(client *HttpClient, playlist []PlaylistItem) (map[string]uint, error) {
	artistTrackCount := map[string]uint{}

	for _, p := range playlist {
		if p.Owner.Id != client.Owner {
			continue // we skip playlists from others
		}

		playlist, err := GetApiStruct[Playlist](client, p.Tracks.Href)
		if err != nil {
			return artistTrackCount, err
		}

		for _, track := range playlist.Items {
			for _, artist := range track.Track.Artists {
				artistTrackCount[artist.Name]++
			}
		}
	}

	return artistTrackCount, nil
}
