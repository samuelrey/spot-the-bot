package framework

import (
	"context"
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
	http.HandleFunc("/callback", authCallback)
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
	// Start the auth callback server before messaging the user.
	// TODO DM the url to the user directly.
	server := serveAuthCallback()
	authUrl := spotifyAuthenticator.AuthURL(state)
	fmt.Printf("%v, %v\n", userID, authUrl)

	token := <-tokenChan

	// Shutdown the HTTP server after receiving the Token.
	if err := server.Shutdown(context.Background()); err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("Error getting token from Spotify")
	}

	return token, nil
}

func serveAuthCallback() *http.Server {
	// Create HTTP server to receive callback from Spotify.
	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("Error starting server: ", err)
		}
	}()

	return server
}

func authCallback(w http.ResponseWriter, r *http.Request) {
	token, err := spotifyAuthenticator.Token(state, r)

	if err != nil {
		fmt.Println("Error getting token, ", err)
		tokenChan <- nil
		return
	}

	if s := r.FormValue("state"); s != state {
		fmt.Println("Error validating state")
		tokenChan <- nil
		return
	}

	fmt.Println("User authorized Spot.")
	tokenChan <- token
}
