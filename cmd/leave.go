package cmd

import (
	"fmt"
)

const StrLeaveFmt = "No hard feelings, %s!\n"

func Leave(ctx *Context) {
	if ctx.UserQueue.Leave(ctx.Actor) {
		content := fmt.Sprintf(StrLeaveFmt, ctx.Actor)
		ctx.Messager.Reply(content)
	}
}
