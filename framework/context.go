package framework

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Replyer       Replyer
	Discord       *discordgo.Session
	Channel       *discordgo.Channel
	EnrolledUsers *[]string
	User          *discordgo.User
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
	ctx.Replyer = discordReplyer
	ctx.Discord = session
	ctx.Channel = channel
	ctx.EnrolledUsers = enrolledUsers
	ctx.User = user
	return ctx
}

func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.Channel.ID, content)
	if err != nil {
		fmt.Println("Error sending message, ", err)
		return nil
	}
	return msg
}

type DiscordReplyer struct {
	Session *discordgo.Session
	Channel *discordgo.Channel
}

func (dr DiscordReplyer) Reply(content string) error {
	_, err := dr.Session.ChannelMessageSend(dr.Channel.ID, content)
	return err
}
