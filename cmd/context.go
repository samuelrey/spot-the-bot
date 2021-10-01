package cmd

import (
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/samuelrey/spot-the-bot/rotation"
)

type Context struct {
	Messager        message.Messager
	PlaylistCreator playlist.Creator
	PlaylistName    string
	UserQueue       *rotation.Rotation
	Actor           message.User
}
