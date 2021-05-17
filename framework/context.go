package framework

import "github.com/bwmarrin/discordgo"

type Context struct {
	Discord *discordgo.Session
	User    *discordgo.User
	UserIDs *[]string
}

func NewContext(discord *discordgo.Session, user *discordgo.User, userIDs *[]string) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.User = user
	ctx.UserIDs = userIDs
	return ctx
}
