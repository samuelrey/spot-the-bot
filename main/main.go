package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuelrey/spot-discord/cmd"
	"github.com/samuelrey/spot-discord/discord"
	"github.com/samuelrey/spot-discord/framework"
	"github.com/samuelrey/spot-discord/spotify"
)

var (
	CmdHandler *framework.CommandHandler
)

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	// Open Discord session.
	fmt.Println("Discord session opening.")
	discordSession, err := discord.DiscordSession(CmdHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Start server to handle Spotify OAuth callback.
	authServer := spotify.StartAuthServer()

	// Cleanup
	defer func() {
		if err := discordSession.Close(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Discord session closed.")
		}

		if err := authServer.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Authentication server shutdown.")
		}
	}()

	fmt.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println()
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
	CmdHandler.Register("list", cmd.List)
	CmdHandler.Register("next", cmd.Next)
}
