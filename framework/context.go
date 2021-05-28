package framework

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	replyer       Replyer
	Discord       *discordgo.Session
	EnrolledUsers *[]string
	User          User
}

func NewContext(
	session *discordgo.Session,
	channel *discordgo.Channel,
	enrolledUsers *[]string,
	user *discordgo.User,
) *Context {
	ctx := new(Context)
	discordReplyer := DiscordReplyer{
		Session: session,
		Channel: channel,
	}
	ctx.replyer = discordReplyer
	ctx.Discord = session
	ctx.EnrolledUsers = enrolledUsers
	ctx.User = User{
		ID:       user.ID,
		Username: user.Username,
	}
	return ctx
}

func (ctx Context) Reply(content string) {
	err := ctx.replyer.Reply(content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
	}
}

type DiscordReplyer struct {
	Session *discordgo.Session
	Channel *discordgo.Channel
}

func (dr DiscordReplyer) Reply(content string) error {
	_, err := dr.Session.ChannelMessageSend(dr.Channel.ID, content)
	return err
}
