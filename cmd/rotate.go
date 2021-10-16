package cmd

import (
	"fmt"
	"log"
)

const (
	StrNextUser = "%s, see you next time around!\n%s, you're up next!"
)

func Rotate(ctx *Context) {
	nextUser, err := ctx.UserQueue.Rotate()
	if err != nil {
		log.Println(err)
	} else {
		content := fmt.Sprintf(StrNextUser, ctx.Actor, nextUser)
		ctx.Messager.Reply(content)
	}
}
