package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"../cmd"
	"../framework"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	CmdHandler           *framework.CommandHandler
	config               *framework.Config
	TknHandler           *framework.TokenHandler
	ch                   = make(chan *oauth2.Token)
	enrolledUsers        = make(map[string]bool)
	spotifyAuthenticator = spotify.NewAuthenticator(
		"http://localhost:8080/callback", spotify.ScopePlaylistModifyPublic)
	state = ""
)

func init() {
	config = framework.LoadConfig("secrets.json")
	spotifyAuthenticator.SetAuthInfo(
		config.SpotifyClientID, config.SpotifyClientSecret)
}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	TknHandler = framework.NewTokenHandler()

	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(commandHandler)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	defer discord.Close()

	members, err := discord.GuildMembers(config.ServerID, "", 1000)
	if err != nil {
		fmt.Println("Error retrieving guild members, ", err)
		return
	}
	for _, member := range members {
		if !member.User.Bot {
			enrolledUsers[member.User.ID] = false
		}
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		tok, err := spotifyAuthenticator.Token(state, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			fmt.Println("Error getting token, ", err)
		}
		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			fmt.Println("Error validating state")
		}
		fmt.Println("User authorized Spot.")
		ch <- tok
	})
	go http.ListenAndServe(":8080", nil)

	fmt.Println("Spot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.Bot {
		return
	}

	// For now only check messages in a specific channel.
	if message.ChannelID != config.ChannelID {
		return
	}

	content := message.Content
	// TODO use a command prefix to filter through messages.
	if len(content) == 0 {
		return
	}

	if content == "auth" {
		if _, found := TknHandler.Get(user.ID); found {
			fmt.Println("already authenticated")
			return
		}
		// The user must visit this URL to authorize Spot.
		// TODO DM the url to the user directly.
		url := spotifyAuthenticator.AuthURL(state)
		fmt.Println(url)

		token := <-ch

		TknHandler.Register(user.ID, token)

		return
	}

	// TODO split content into command name and arguments.
	command, found := CmdHandler.Get(content)
	if !found {
		return
	}

	channel, err := discord.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error retrieving channel, ", err)
		return
	}

	ctx := framework.NewContext(discord, channel, enrolledUsers, config.PlaylistLink, user)
	c := *command
	c(ctx)
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
	CmdHandler.Register("link", cmd.Link)
}
