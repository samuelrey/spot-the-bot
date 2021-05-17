package framework

import "github.com/bwmarrin/discordgo"

type Context struct {
	Discord       *discordgo.Session
	User          *discordgo.User
	EnrolledUsers map[string]bool
}

func NewContext(discord *discordgo.Session, user *discordgo.User, enrolledUsers map[string]bool) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.User = user
	ctx.EnrolledUsers = enrolledUsers
	return ctx
}
