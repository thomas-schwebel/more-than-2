package stats

import (
	"sort"
)

func SortByCount(artistTrackCount map[string]uint) []ArtistCount {
	var sorted []ArtistCount
	for artist, count := range artistTrackCount {
		sorted = append(sorted, ArtistCount{artist, count})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count > sorted[j].Count
	})

	return sorted
}
