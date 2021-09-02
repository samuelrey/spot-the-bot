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

	commandRegistry := framework.NewCommandRegistry()
	registerCommands(*commandRegistry)

	spotifyConfig := spotify.LoadConfigFromEnv()
	sa := spotify.NewSpotifyAuthorizer(spotifyConfig)
	sp, err := sa.AuthorizeUser()
	if err != nil {
		log.Println(err)
		return
	}

	discordConfig := discord.LoadConfigFromEnv()
	d, err := discord.NewDiscordBuilder(discordConfig, commandRegistry, &enrolledUsers, sp)
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

func registerCommands(commandRegistry framework.CommandRegistry) {
	commandRegistry.Register("join", cmd.Join, "helloWorld")
	commandRegistry.Register("leave", cmd.Leave, "helloWorld")
	commandRegistry.Register("list", cmd.List, "helloWold")
	commandRegistry.Register("next", cmd.Next, "helloWorld")
	commandRegistry.Register("create", cmd.Create, "helloWorld")
}
