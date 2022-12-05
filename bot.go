package bot

import "net/http"

type Bot struct {
	Token  string      `json:"token"`
	Client http.Client `json:"-"`
}

type ClientOpts struct {
	Client http.Client
}

func CreateBot(token string, clientOpts *ClientOpts) (*Bot, error) {
	b := &Bot{
		Token:  token,
		Client: http.Client{},
	}

	return b, nil
}
