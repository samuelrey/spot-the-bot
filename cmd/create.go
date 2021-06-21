package cmd

import (
	"fmt"
	"log"

	"github.com/samuelrey/spot-discord/framework"
)

const StrPlaylistCreatedFmt = "Done! :tada: Now it's up to to you to " +
	"add a few tracks to set the vibe :performing_arts:. " +
	"Then share it in channel! :headphones:\n%s\n"

// TODO test
func Create(ctx *framework.CommandContext) {
	if len(*ctx.EnrolledUsers) < 1 {
		return
	}

	if ctx.Actor.ID != (*ctx.EnrolledUsers)[0].ID {
		return
	}

	playlist, err := ctx.CreatePlaylist("Einstok")
	if err != nil {
		log.Println(err)
		return
	}

	content := fmt.Sprintf(StrPlaylistCreatedFmt, playlist.URL)
	ctx.DirectMessage(ctx.Actor.ID, content)
}
