package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const (
	StrListNoUsers  = "Nobody's joined yet, join with `!join`"
	StrListUsersFmt = "Here's the queue:\n%s\n"
)

func List(ctx *framework.CommandContext) {
	if len(*ctx.EnrolledUsers) == 0 {
		ctx.Reply(StrListNoUsers)
		return
	}

	content := fmt.Sprintf(StrListUsersFmt, *ctx.EnrolledUsers)
	ctx.Reply(content)
}
