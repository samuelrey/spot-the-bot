package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/slack-go/slack"
)

type Secret struct {
	Token        string `json:"SLACK_TOKEN"`
	ChannelID    string `json:"SLACK_CHANNEL_ID"`
	PlaylistLink string `json:"SPOTIFY_PLAYLIST"`
}

func main() {
	f, err := os.Open("secrets.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(f)
	var secret Secret
	json.Unmarshal(byteValue, &secret)

	api := slack.New(secret.Token)
	attachment := slack.Attachment{
		Pretext: "Playlist link",
		Text:    secret.PlaylistLink,
	}

	channelID, timestamp, err := api.PostMessage(
		secret.ChannelID,
		slack.MsgOptionText("This week's playlist", false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
