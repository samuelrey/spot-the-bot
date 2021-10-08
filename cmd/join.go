package cmd

import (
	"fmt"
	"log"
)

const StrJoinFmt = "Welcome to the club, %s!\n"

func Join(ctx *Context) {
	err := ctx.UserQueue.Join(ctx.Actor)
	if err != nil {
		log.Println(err)
	} else {
		content := fmt.Sprintf(StrJoinFmt, ctx.Actor)
		ctx.Messager.Reply(content)
	}
}
