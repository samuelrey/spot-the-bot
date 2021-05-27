package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ClientID    string `json:"CLIENT_ID"`
	RedirectURL string `json:"REDIRECT_URL"`
	Secret      string `json:"SECRET"`
	State       string `json:"STATE"`
}

func LoadConfig(filename string) *Config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config, ", err)
		return nil
	}
	var config Config
	json.Unmarshal(body, &config)
	return &config
}
