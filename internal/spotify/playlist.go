package spotify

type Artist struct {
	Id   string
	Name string
}

type Track struct {
	Id      string
	Name    string
	Artists []Artist
}

type Playlist struct {
	Items []PlaylistItem
}

type PlaylistOwner struct {
	Id string
}

type PlaylistTracks struct {
	Total uint
	Href  string
}

type PlaylistItem struct {
	Track  Track
	Owner  *PlaylistOwner
	Tracks *PlaylistTracks
}

type Playlists struct {
	Items []PlaylistItem
	Total uint
	Limit uint
	Next  string
}
