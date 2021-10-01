package discord

import (
	"os"
)

type DiscordConfig struct {
	Token string
}

func LoadConfig() DiscordConfig {
	return DiscordConfig{
		Token: os.Getenv("DISCORD_TOKEN"),
	}
}
