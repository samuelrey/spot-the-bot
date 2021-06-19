package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

// NewContext takes discord specific data and produces a Context which can be
// used by commands.
func NewContext(
	session *discordgo.Session,
	channelID string,
	enrolledUsers *[]framework.User,
	actor *discordgo.User,
) *framework.CommandContext {
	ctx := new(framework.CommandContext)
	ctx.EnrolledUsers = enrolledUsers
	ctx.Actor = framework.User{
		ID:       actor.ID,
		Username: actor.Username,
	}
	ctx.Messager = &DiscordMessager{
		session:   session,
		channelID: channelID,
	}
	return ctx
}
