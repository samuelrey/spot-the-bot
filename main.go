package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-discord/cmd"
	"github.com/samuelrey/spot-discord/discord"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"
)

const prefix = "!"

var (
	ch            *framework.CommandHandler
	pc            framework.PlaylistCreator
	enrolledUsers []framework.MessageUser
	err           error
)

func main() {
	enrolledUsers = make([]framework.MessageUser, 0)

	ch = framework.NewCommandHandler()
	registerCommands(*ch)

	spotifyConfig := spotify.LoadConfig("secrets_spotify.json")
	sa := spotify.NewSpotifyAuthorizer(spotifyConfig)
	pc, err = sa.AuthorizeUser()
	if err != nil {
		log.Println(err)
		return
	}

	discordConfig := discord.LoadConfig("secrets_discord.json")
	d, err := discordgo.New("Bot " + discordConfig.Token)
	if err != nil {
		log.Println(err)
		return
	}

	d.AddHandler(handleMessage)
	d.Identify.Intents = discordgo.IntentsGuildMessages

	log.Println("Discord session opening.")
	err = d.Open()
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err := d.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	log.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	fmt.Println()
}

func registerCommands(cmdHandler framework.CommandHandler) {
	cmdHandler.Register("join", cmd.Join)
	cmdHandler.Register("leave", cmd.Leave)
	cmdHandler.Register("list", cmd.List)
	cmdHandler.Register("next", cmd.Next)
	cmdHandler.Register("create", cmd.Create)
}

func handleMessage(dg *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.Bot {
		return
	}

	content := message.Content

	if len(content) <= len(prefix) {
		return
	}

	if content[:len(prefix)] != prefix {
		return
	}

	args := strings.Fields(content[len(prefix):])
	name := strings.ToLower(args[0])

	command, found := ch.Get(name)
	if !found {
		return
	}

	ctx := framework.CommandContext{
		Messager: &discord.DiscordMessager{
			Session:   dg,
			ChannelID: message.ChannelID,
		},
		PlaylistCreator: pc,
		PlaylistName:    "Einstok",
		EnrolledUsers:   &enrolledUsers,
		Actor: framework.MessageUser{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	c := *command
	c(&ctx)
}
