package cmd

import (
	"fmt"
	"log"

	"github.com/samuelrey/spot-the-bot/message"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *Context) {
	rotationRepository := ctx.RepositoryProvider.GetRotationRepository()
	rotation, err := rotationRepository.FindOne(ctx.ServerID)
	if err != nil {
		r := message.NewRotation([]message.User{}, ctx.ServerID)
		rotation = &r
	}

	err = rotation.Join(ctx.Actor)
	if err != nil {
		log.Println(err)
		return
	}

	err = rotationRepository.Upsert(*rotation)
	if err != nil {
		log.Println(err)
		return
	}

	content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
	ctx.Messager.Reply(content)
}
