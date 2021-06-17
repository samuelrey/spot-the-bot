package framework

type Playlist struct {
	ID  string
	URL string
}
type PlaylistBuilder interface {
	CreatePlaylist(playlistName string) (*Playlist, error)
}

type MusicServiceAuthorizer interface {
	AuthorizeUser() error
}
