package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/framework"
)

// NewContext takes discord specific data and produces a Context which can be
// used by commands.
func NewContext(
	session *discordgo.Session,
	channel *discordgo.Channel,
	enrolledUsers *[]framework.User,
	actor *discordgo.User,
) *framework.Context {
	ctx := new(framework.Context)
	discordReplyer := DiscordReplyer{
		Session: session,
		Channel: channel,
	}
	ctx.Replyer = discordReplyer
	ctx.EnrolledUsers = enrolledUsers
	ctx.Actor = framework.User{
		ID:       actor.ID,
		Username: actor.Username,
	}
	return ctx
}

type DiscordReplyer struct {
	Session *discordgo.Session
	Channel *discordgo.Channel
}

func (dr DiscordReplyer) Reply(content string) error {
	_, err := dr.Session.ChannelMessageSend(dr.Channel.ID, content)
	return err
}
