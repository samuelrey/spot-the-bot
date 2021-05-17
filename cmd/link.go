package cmd

import (
	"fmt"

	"../framework"
)

func Link(ctx *framework.Context) {
	content := fmt.Sprintf("Here you go!\n%v\n", ctx.SpotifyLink)
	ctx.Reply(content)
}
