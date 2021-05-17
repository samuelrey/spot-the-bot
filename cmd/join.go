package cmd

import (
	"fmt"

	"../framework"
)

func Join(ctx *framework.Context) {
	ctx.EnrolledUsers[ctx.User.ID] = true
	fmt.Println("new user: ", ctx.User.ID)
}
