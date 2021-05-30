package framework

import (
	"fmt"
)

type Context struct {
	Replyer       Replyer
	EnrolledUsers *[]User
	Actor         User
}

func (ctx Context) Reply(content string) {
	err := ctx.Replyer.Reply(content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}

type User struct {
	ID       string
	Username string
}

func (u User) String() string {
	return u.Username
}
