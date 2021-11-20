package cmd

import (
	"fmt"
	"log"
)

const StrPlaylistCreatedFmt = "Done! :tada: Now it's up to to you to " +
	"add a few tracks to set the vibe :performing_arts:. " +
	"Then share it in channel! :headphones:\n%s\n"

func Create(ctx *Context) {
	rotationRepository := ctx.RepositoryProvider.GetRotationRepository()
	rotation, err := rotationRepository.FindOne(ctx.ServerID)
	if err != nil {
		log.Println(err)
		return
	}

	head := rotation.Current()
	if head == nil || ctx.Actor.ID != head.ID {
		return
	}

	playlist, err := ctx.PlaylistCreator.Create(ctx.PlaylistName)
	if err != nil {
		log.Println(err)
		return
	}

	content := fmt.Sprintf(StrPlaylistCreatedFmt, playlist.URL)
	ctx.Messager.DirectMessage(ctx.Actor.ID, content)
}
