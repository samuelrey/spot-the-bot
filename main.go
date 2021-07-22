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

var (
	c   config
	ch  framework.CommandHandler
	pc  framework.PlaylistCreator
	uq  framework.UserQueue
	err error
)

func main() {
	c = loadConfigFromEnv()

	uq = framework.NewUserQueue([]framework.MessageUser{})

	ch = framework.NewCommandHandler()
	registerCommands(ch)

	pc, err = spotify.NewPlaylistCreator(c.SpotifyConfig)
	if err != nil {
		log.Println(err)
		return
	}

	discordSession, err := discordgo.New("Bot " + c.Token)
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

type config struct {
	*discord.DiscordConfig
	spotify.SpotifyConfig
	Prefix string
}

func loadConfigFromEnv() config {
	return config{
		DiscordConfig: discord.LoadConfigFromEnv(),
		SpotifyConfig: spotify.LoadConfigFromEnv(),
		Prefix:        os.Getenv("SPOT_PREFIX"),
	}
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

	if len(content) <= len(c.Prefix) {
		return
	}

	if content[:len(c.Prefix)] != c.Prefix {
		return
	}

	args := strings.Fields(content[len(c.Prefix):])
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
		UserQueue:       &uq,
		Actor: framework.MessageUser{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	c := *command
	c(&ctx)
}
