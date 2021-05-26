package framework

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type (
	TknMap       map[string]*oauth2.Token
	TokenHandler struct {
		tnks TknMap
	}
)

var (
	redirUrl             = "http://localhost:8080/callback"
	spotifyAuthenticator = spotify.NewAuthenticator(
		redirUrl, spotify.ScopePlaylistModifyPublic)
	state     = ""
	tokenChan = make(chan *oauth2.Token)
)

func init() {
	spotifyAuthenticator.SetAuthInfo("", "")
}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{make(TknMap)}
}

func (tknHandler TokenHandler) Get(key string) (*oauth2.Token, bool) {
	tkn, found := tknHandler.tnks[key]
	return tkn, found
}

func (tknHandler TokenHandler) Register(key string, token *oauth2.Token) {
	tknHandler.tnks[key] = token
}

func SpotifyClient(token *oauth2.Token) *spotify.Client {
	client := spotifyAuthenticator.NewClient(token)
	return &client
}

func AuthorizeSpotForUser(userID string) (*oauth2.Token, error) {
	// TODO DM the url to the user directly.
	authUrl := spotifyAuthenticator.AuthURL(state)
	fmt.Printf("%v, %v\n", userID, authUrl)

	token := <-tokenChan

	if token == nil {
		return nil, fmt.Errorf("didn't get token from Spotify")
	}

	return token, nil
}

// StartAuthServer creates HTTP server to handle callback request from Spotify.
// This function is meant to be called only once.
func StartAuthServer() *http.Server {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/callback", authCallback)

	go func() {
		fmt.Println("Authentication server starting")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return server
}

func authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyAuthenticator.Token(state, r)

	if err != nil {
		select {
		case tokenChan <- nil:
			fmt.Println("Error getting token,", err)
		default:
			// Protect against blocking the goroutine.
			fmt.Println("Error endpoint accessed directly")
		}
		return
	}

	fmt.Println("User authorized Spot.")
	tokenChan <- token
}
