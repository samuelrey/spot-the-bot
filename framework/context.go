package framework

import (
	"fmt"
)

type Context struct {
	Replyer       Replyer
	EnrolledUsers *[]User
	User          User
}

func (ctx Context) Reply(content string) {
	err := ctx.Replyer.Reply(content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}
