package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"
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

	spotifyClient := spotify.Client(ctx)
	if spotifyClient == nil {
		return
	}

	user, err := spotifyClient.CurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	playlist, err := spotifyClient.CreatePlaylistForUser(user.ID, "einstock", "", true)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := fmt.Sprintf(StrPlaylistCreatedFmt, playlist.ExternalURLs["spotify"])
	ctx.DirectMessage(ctx.Actor.ID, content)
}
