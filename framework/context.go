package framework

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Discord       *discordgo.Session
	Channel       *discordgo.Channel
	EnrolledUsers *[]string
	SpotifyLink   string
	User          *discordgo.User
}

func NewContext(
	discord *discordgo.Session,
	channel *discordgo.Channel,
	enrolledUsers *[]string,
	spotifyLink string,
	user *discordgo.User,
) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Channel = channel
	ctx.EnrolledUsers = enrolledUsers
	ctx.SpotifyLink = spotifyLink
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
