package cmd

import (
	"fmt"

	"../framework"
)

func Leave(ctx *framework.Context) {
	ctx.EnrolledUsers[ctx.User.ID] = false
	fmt.Println("remove user: ", ctx.User.ID)
}
