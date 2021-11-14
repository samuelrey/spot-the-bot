package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/samuelrey/spot-the-bot/cmd"
	"github.com/samuelrey/spot-the-bot/discord"
	"github.com/samuelrey/spot-the-bot/message"
	"github.com/samuelrey/spot-the-bot/playlist"
	"github.com/samuelrey/spot-the-bot/spotify"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	conf            config
	commandRegistry *cmd.Registry
	playlistCreator playlist.Creator
	rotationRepo    *message.RotationRepository
	err             error
)

func main() {
	conf = loadConfigFromEnv()

	commandRegistry = cmd.NewRegistry()
	registerCommands(*commandRegistry)

	playlistCreator, err = spotify.NewCreator(conf.SpotifyConfig)
	if err != nil {
		log.Println(err)
		return
	}

	rotationRepo, err = getRotationRepository()
	if err != nil {
		log.Println(err)
		return
	}

	discordSession, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.Println(err)
		return
	}

	discordSession.AddHandler(handleMessage)
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages

	log.Println("Discord session opening.")
	err = discordSession.Open()
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		err := discordSession.Close()
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

// handleMessage reads messages sent in the guild and runs commands based on
// those messages.
func handleMessage(
	dg *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	user := m.Author
	if user.Bot {
		return
	}

	content := m.Content

	if len(content) <= len(conf.Prefix) {
		return
	}

	if content[:len(conf.Prefix)] != conf.Prefix {
		return
	}

	args := strings.Fields(content[len(conf.Prefix):])
	name := strings.ToLower(args[0])

	command, found := commandRegistry.Get(name)
	if !found {
		return
	}

	// TODO get playlist name from config
	// TODO optionally populate dependencies based on command
	ctx := cmd.Context{
		Messager: &discord.Messager{
			Session:   dg,
			ChannelID: m.ChannelID,
		},
		PlaylistCreator:    playlistCreator,
		PlaylistName:       "Einstok",
		RotationRepository: rotationRepo,
		ServerID:           m.GuildID,
		Actor: message.User{
			ID:       user.ID,
			Username: user.Username,
		},
	}
	command.RunWithContext(&ctx)
}

type config struct {
	discord.DiscordConfig
	spotify.SpotifyConfig
	Prefix   string
	MongoURI string
}

func loadConfigFromEnv() config {
	return config{
		DiscordConfig: discord.LoadConfig(),
		SpotifyConfig: spotify.LoadConfig(),
		Prefix:        os.Getenv("SPOT_PREFIX"),
		MongoURI:      os.Getenv("MONGO_URI"),
	}
}

func registerCommands(cr cmd.Registry) {
	cr.Register("join", cmd.Join, "Join the rotation of people to start a playlist.")
	cr.Register("leave", cmd.Leave, "Leave the rotation. Remember you could still listen and add to playlists!")
	cr.Register("list", cmd.List, "See the rotation.")
	cr.Register("rotate", cmd.Rotate, "Move to the next person in the rotation.")
	cr.Register("create", cmd.Create, "Create a playlist.")
}

type iRepositoryProvider interface {
	getRotationRepository() message.IRotationRepository
}

type repositoryProvider struct {
	database           *mongo.Database
	rotationRepository message.IRotationRepository
}

func newRepositoryProvider() (iRepositoryProvider, error) {
	dbOpt := options.Client().ApplyURI(conf.MongoURI)
	mongoClient, err := mongo.Connect(context.TODO(), dbOpt)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := mongoClient.Database("spot-the-bot")
	rotationCollection := database.Collection("rotations")
	provider := repositoryProvider{
		database:           database,
		rotationRepository: message.NewRotationRepository(rotationCollection),
	}
	return &provider, nil
}

func (rp *repositoryProvider) getRotationRepository() message.IRotationRepository {
	return rp.rotationRepository
}

func getRotationRepository() (*message.RotationRepository, error) {
	dbOpt := options.Client().ApplyURI(conf.MongoURI)
	mongoClient, err := mongo.Connect(context.TODO(), dbOpt)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	mongoDB := mongoClient.Database("spot-the-bot")
	rotationCollection := mongoDB.Collection("rotations")

	return message.NewRotationRepository(rotationCollection), nil
}
