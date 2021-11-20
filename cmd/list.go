package cmd

import (
	"fmt"
	"log"
)

const (
	StrListNoUsers  = "Nobody's joined yet, join with `!join`"
	StrListUsersFmt = "Here's the queue:\n%s\n"
)

func List(ctx *Context) {
	rotationRepository := ctx.RepositoryProvider.GetRotationRepository()
	rotation, err := rotationRepository.FindOne(ctx.ServerID)
	if err != nil {
		log.Println(err)
		return
	}

	if rotation.Length() == 0 {
		ctx.Messager.Reply(StrListNoUsers)
		return
	}

	content := fmt.Sprintf(StrListUsersFmt, rotation)
	ctx.Messager.Reply(content)
}
