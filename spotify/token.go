package spotify

import "golang.org/x/oauth2"

type (
	TknMap       map[string]*oauth2.Token
	TokenHandler struct {
		tnks TknMap
	}
)

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
