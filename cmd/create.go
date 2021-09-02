package cmd

import (
	"fmt"
	"log"

	"github.com/samuelrey/spot-the-bot/framework"
)

const StrPlaylistCreatedFmt = "Done! :tada: Now it's up to to you to " +
	"add a few tracks to set the vibe :performing_arts:. " +
	"Then share it in channel! :headphones:\n%s\n"

func Create(ctx *framework.CommandContext) {
	head := ctx.UserQueue.Head()
	if head == nil || ctx.Actor.ID != head.ID {
		return
	}

	playlist, err := ctx.CreatePlaylist(ctx.PlaylistName)
	if err != nil {
		log.Println(err)
		return
	}

	content := fmt.Sprintf(StrPlaylistCreatedFmt, playlist.URL)
	ctx.DirectMessage(ctx.Actor.ID, content)
}
