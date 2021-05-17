package framework

import "github.com/bwmarrin/discordgo"

type Context struct {
	Discord       *discordgo.Session
	Channel       *discordgo.Channel
	EnrolledUsers map[string]bool
	SpotifyLink   string
	User          *discordgo.User
}

func NewContext(discord *discordgo.Session, channel *discordgo.Channel, enrolledUsers map[string]bool, spotifyLink string, user *discordgo.User) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Channel = channel
	ctx.EnrolledUsers = enrolledUsers
	ctx.SpotifyLink = spotifyLink
	ctx.User = user
	return ctx
}

func (ctx Context) Reply(content string) (*discordgo.Message, error) {
	return ctx.Discord.ChannelMessageSend(ctx.Channel.ID, content)
}
