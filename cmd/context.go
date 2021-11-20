package cmd

import (
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/samuelrey/spot-the-bot/repository"
)

type Context struct {
	Messager           message.Messager
	PlaylistCreator    playlist.Creator
	PlaylistName       string
	Actor              message.User
	RepositoryProvider repository.IProvider
	ServerID           string
}
