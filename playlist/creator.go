package playlist

type Playlist struct {
	ID  string
	URL string
}

type Creator interface {
	Create(playlistName string) (*Playlist, error)
}
