package framework

import (
	"fmt"
)

type Context struct {
	Messager      Messager
	EnrolledUsers *[]User
	Actor         User
}

func (ctx Context) Reply(content string) {
	err := ctx.Messager.Reply(content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}

func (ctx Context) DirectMessage(recipient, content string) {
	err := ctx.Messager.DirectMessage(recipient, content)
	if err != nil {
		fmt.Println("Error sending direct message, ", err)
	}
}

type User struct {
	ID       string
	Username string
}

func (u User) String() string {
	return u.Username
}
