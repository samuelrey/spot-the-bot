package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-discord/framework"
)

func Link(ctx *framework.Context) {
	content := fmt.Sprintf("Here you go!\n%v\n", ctx.SpotifyLink)
	ctx.Reply(content)
}
