package spotify

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

var tokenChan = make(chan *oauth2.Token)

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
	token, err := spotifyAuthenticator.Token(config.State, r)

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

	select {
	case tokenChan <- token:
		fmt.Println("User authorized Spot.")
	default:
		// Protect against blocking the goroutine.
		fmt.Println("User already authorized Spot.")
	}
}

func getToken() (*oauth2.Token, error) {
	token := <-tokenChan
	if token == nil {
		return nil, fmt.Errorf("didn't get token from Spotify")
	}

	return token, nil
}
