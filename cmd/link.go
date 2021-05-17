package cmd

import (
	"fmt"

	"../framework"
)

func Link(ctx *framework.Context) {
	content := fmt.Sprintf("Here you go!\n%v\n", ctx.SpotifyLink)
	_, err := ctx.Reply(content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}
