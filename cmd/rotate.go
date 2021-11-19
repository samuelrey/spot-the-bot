package cmd

import (
	"fmt"
	"log"
)

const (
	StrNextUser = "%s, see you next time around!\n%s, you're up next!"
)

func Rotate(ctx *Context) {
	rotationRepository := ctx.RepositoryProvider.GetRotationRepository()
	rotation, err := rotationRepository.FindOne(ctx.ServerID)
	if err != nil {
		log.Println(err)
		return
	}

	nextUser, err := rotation.Rotate()
	if err != nil {
		log.Println(err)
		return
	}

	err = rotationRepository.Upsert(*rotation)
	if err != nil {
		log.Println(err)
		return
	}
	content := fmt.Sprintf(StrNextUser, ctx.Actor, nextUser)
	ctx.Messager.Reply(content)
}
