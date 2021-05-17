package cmd

import (
	"fmt"

	"../framework"
)

func Leave(ctx *framework.Context) {
	found := -1
	for i, id := range *ctx.UserIDs {
		if id == ctx.User.ID {
			found = i
			break
		}
	}
	if found != -1 {
		*ctx.UserIDs = append((*ctx.UserIDs)[:found], (*ctx.UserIDs)[found+1:]...)
		fmt.Println("remove user: ", ctx.User.ID)
	}
}
