package spotify

import (
	"net/http"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const STATE = "spot-the-bot"

var tokenChan = make(chan *oauth2.Token)
var errChan = make(chan error)

type authenticator struct {
	spotify.Authenticator
	authURL string
}

func newAuthenticator(conf SpotifyConfig) authenticator {
	a := spotify.NewAuthenticator(
		conf.RedirectURL,
		spotify.ScopePlaylistModifyPrivate,
	)
	a.SetAuthInfo(conf.ClientID, conf.Secret)
	authURL := a.AuthURL(STATE)

	return authenticator{
		Authenticator: a,
		authURL:       authURL,
	}
}

// startAuthServer creates HTTP server to handle callback request from Spotify.
// authCallback waits for the request from Spotify containing the oauth2 token
// for using their API.
func (a authenticator) startAuthServer() *http.Server {
	server := &http.Server{Addr: ":8080"}

	authCallback := func(w http.ResponseWriter, r *http.Request) {
		token, err := a.Token(STATE, r)
		if err != nil {
			errChan <- err
			return
		}

		tokenChan <- token
	}

	http.HandleFunc("/callback", authCallback)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return server
}

func getToken() (*oauth2.Token, error) {
	select {
	case token := <-tokenChan:
		return token, nil
	case err := <-errChan:
		return nil, err
	}
}
