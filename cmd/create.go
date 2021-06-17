package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const StrPlaylistCreatedFmt = "Done! :tada: Now it's up to to you to " +
	"update the title :sa:, change the cover photo :frame_photo: and " +
	"add a few tracks to set the vibe :performing_arts:. " +
	"Then share it in channel! :headphones:\n%v\n"

func Create(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) < 1 {
		return
	}

	if ctx.Actor.ID != (*ctx.EnrolledUsers)[0].ID {
		return
	}

	// TODO create playlist.
	content := fmt.Sprintf(StrPlaylistCreatedFmt, "URL")
	ctx.DirectMessage(ctx.Actor.ID, content)
}
