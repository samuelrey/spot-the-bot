package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuelrey/spot-discord/cmd"
	"github.com/samuelrey/spot-discord/discord"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"
)

func main() {
	cmdHandler := framework.NewCommandHandler()
	registerCommands(*cmdHandler)

	spotifyConfig := spotify.LoadConfig("secrets_spotify.json")
	_ = spotify.CreateSpotifyBuilder(spotifyConfig)

	// Open Discord session.
	log.Println("Discord session opening.")
	discordSession, err := discord.DiscordSession(cmdHandler)
	if err != nil {
		log.Println(err)
		return
	}

	// Start server to handle Spotify OAuth callback.
	authServer := spotify.StartAuthServer()

	// Cleanup
	defer func() {
		if err := discordSession.Close(); err != nil {
			log.Println(err)
		} else {
			log.Println("Discord session closed.")
		}

		if err := authServer.Shutdown(context.Background()); err != nil {
			log.Println(err)
		} else {
			log.Println("Authentication server shutdown.")
		}
	}()

	log.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
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
