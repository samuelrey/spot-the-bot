package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuelrey/spot-the-bot/cmd"
	"github.com/samuelrey/spot-the-bot/discord"
	"github.com/samuelrey/spot-the-bot/framework"
	"github.com/samuelrey/spot-the-bot/spotify"
)

func main() {
	enrolledUsers := make([]framework.MessageUser, 0)

	cmdHandler := framework.NewCommandHandler()
	registerCommands(*cmdHandler)

	spotifyConfig := spotify.LoadConfigFromEnv()
	sa := spotify.NewSpotifyAuthorizer(spotifyConfig)
	sp, err := sa.AuthorizeUser()
	if err != nil {
		log.Println(err)
		return
	}

	discordConfig := discord.LoadConfigFromEnv()
	d, err := discord.NewDiscordBuilder(discordConfig, cmdHandler, &enrolledUsers, sp)
	if err != nil {
		log.Println(err)
		return
	}

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
