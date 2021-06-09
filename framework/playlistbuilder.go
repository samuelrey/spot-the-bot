package framework

type PlaylistBuilder struct {
	MusicServiceBuilder
}

type Playlist struct {
	ID  string
	URL string
}

type MusicServiceBuilder interface {
	AuthorizeUser(msgUserID string, sendAuthURL func(string)) MusicServiceClient
}

type MusicServiceClient interface {
	CreatePlaylistForUser(userID, playlistName string) (*Playlist, error)
}
