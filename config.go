package main

import (
	"os"

	"github.com/samuelrey/spot-the-bot/discord"
	"github.com/samuelrey/spot-the-bot/spotify"
)

type config struct {
	discord.DiscordConfig
	spotify.SpotifyConfig
	Prefix string
}

func loadConfigFromEnv() config {
	return config{
		DiscordConfig: discord.LoadConfig(),
		SpotifyConfig: spotify.LoadConfig(),
		Prefix:        os.Getenv("SPOT_PREFIX"),
	}
}
