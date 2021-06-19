package framework

type Playlist struct {
	ID  string
	URL string
}

type PlaylistCreator interface {
	CreatePlaylist(playlistName string) (*Playlist, error)
}

type MusicAuthorizer interface {
	AuthorizeUser() error
}
