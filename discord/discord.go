package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var config *Config

func init() {
	config = LoadConfig("secrets_discord.json")
}

func DiscordClient(commandHandler func(discord *discordgo.Session, message *discordgo.MessageCreate)) *discordgo.Session {
	client, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = client.Open()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	client.AddHandler(commandHandler)
	client.Identify.Intents = discordgo.IntentsGuildMessages
	return client
}
