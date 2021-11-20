package cmd

import (
	"fmt"
	"log"
)

const StrLeaveFmt = "No hard feelings, %s!\n"

func Leave(ctx *Context) {
	rotationRepository := ctx.RepositoryProvider.GetRotationRepository()
	rotation, err := rotationRepository.FindOne(ctx.ServerID)
	if err != nil {
		log.Println(err)
		return
	}

	if rotation.Leave(ctx.Actor) {
		err = rotationRepository.Upsert(*rotation)
		if err != nil {
			log.Println(err)
			return
		}
		content := fmt.Sprintf(StrLeaveFmt, ctx.Actor)
		ctx.Messager.Reply(content)
	}
}
