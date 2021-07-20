package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-the-bot/cmd"
	"github.com/samuelrey/spot-the-bot/discord"
	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/samuelrey/spot-the-bot/spotify"
)

// TODO make prefix configurable.
const prefix = "!"

var (
	ch  *framework.CommandHandler
	pc  framework.PlaylistCreator
	uq  framework.UserQueue
	err error
)

func main() {
	q := framework.NewSimpleUserQueue([]framework.MessageUser{})
	uq = &q

	ch = framework.NewCommandHandler()
	registerCommands(*ch)

	spotifyConfig := spotify.LoadConfigFromEnv()
	spotifyAuthorizer := spotify.NewSpotifyAuthorizer(spotifyConfig)
	pc, err = spotifyAuthorizer.AuthorizeUser()
	if err != nil {
		log.Println(err)
		return
	}

	discordConfig := discord.LoadConfigFromEnv()
	discordSession, err := discordgo.New("Bot " + discordConfig.Token)
	if err != nil {
		log.Println(err)
		return
	}

	discordSession.AddHandler(handleMessage)
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	log.Println("Discord session opening.")
	err = discordSession.Open()
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err := discordSession.Close()
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

// handleMessage reads messages sent in the guild and runs commands based on
// those messages.
func handleMessage(
	dg *discordgo.Session,
	message *discordgo.MessageCreate,
) {
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

	// TODO get playlist name from config
	// TODO optionally populate dependencies based on command
	ctx := framework.CommandContext{
		Messager: &discord.DiscordMessager{
			Session:   dg,
			ChannelID: message.ChannelID,
		},
		PlaylistCreator: pc,
		PlaylistName:    "Einstok",
		UserQueue:       uq,
		Actor: framework.MessageUser{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	c := *command
	c(&ctx)
}
