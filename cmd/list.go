package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

const (
	StrListNoUsers  = "Nobody's joined yet, join with `!join`"
	StrListUsersFmt = "Here's the queue:\n%s\n"
)

func List(ctx *framework.Context) {
	if len(*ctx.EnrolledUsers) == 0 {
		ctx.Reply(ctx.ChannelID, StrListNoUsers)
		return
	}

	content := fmt.Sprintf(StrListUsersFmt, *ctx.EnrolledUsers)
	ctx.Reply(ctx.ChannelID, content)
}
