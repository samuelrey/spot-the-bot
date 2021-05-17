package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"../cmd"
	"../framework"

	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler    *framework.CommandHandler
	config        *framework.Config
	enrolledUsers map[string]bool
)

func init() {
	config = framework.LoadConfig("secrets.json")
}

func main() {
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

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

	enrolledUsers = make(map[string]bool)
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

	discord.AddHandler(commandHandler)
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	defer discord.Close()

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

	// TODO split content into command name and arguments.
	command, found := CmdHandler.Get(content)
	if !found {
		return
	}

	ctx := framework.NewContext(discord, user, enrolledUsers)
	c := *command
	c(ctx)
}

func registerCommands() {
	CmdHandler.Register("join", cmd.Join)
	CmdHandler.Register("leave", cmd.Leave)
}
