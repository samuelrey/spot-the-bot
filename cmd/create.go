package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"
)

func Create(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) < 1 {
		return
	}

	if ctx.Actor.ID != (*ctx.EnrolledUsers)[0].ID {
		return
	}

	spotifyClient := spotify.Client(ctx.Actor.ID)
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

	fmt.Println(playlist)
}
