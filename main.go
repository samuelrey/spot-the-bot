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
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/samuelrey/spot-the-bot/rotation"
	"github.com/samuelrey/spot-the-bot/spotify"
)

var (
	c   config
	cr  *cmd.Registry
	pc  playlist.Creator
	uq  rotation.Rotation
	err error
)

func main() {
	c = loadConfigFromEnv()

	cr = cmd.NewRegistry()
	registerCommands(*cr)
	uq = rotation.NewRotation([]message.User{})

	pc, err = spotify.NewCreator(c.SpotifyConfig)
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

// handleMessage reads messages sent in the guild and runs commands based on
// those messages.
func handleMessage(
	dg *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	user := m.Author
	if user.Bot {
		return
	}

	content := m.Content

	if len(content) <= len(c.Prefix) {
		return
	}

	if content[:len(c.Prefix)] != c.Prefix {
		return
	}

	args := strings.Fields(content[len(c.Prefix):])
	name := strings.ToLower(args[0])

	command, found := cr.Get(name)
	if !found {
		return
	}

	// TODO get playlist name from config
	// TODO optionally populate dependencies based on command
	ctx := cmd.Context{
		Messager: &discord.Messager{
			Session:   dg,
			ChannelID: m.ChannelID,
		},
		PlaylistCreator: pc,
		PlaylistName:    "Einstok",
		UserQueue:       &uq,
		Actor: message.User{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	command.RunWithContext(&ctx)
}
