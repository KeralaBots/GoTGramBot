package main

import (
	"net/http"

	bot "github.com/KeralaBots/GoTGramBot"
	"github.com/KeralaBots/GoTGramBot/filters"
	"github.com/KeralaBots/GoTGramBot/types"
)

func test1(b *bot.Bot, m *types.CallbackQuery) error {
	_, err := b.SendPhoto(
		m.Message.Chat.Id,
		"test.jpg",
		&bot.SendPhotoOpts{},
	)

	return err
}

func main() {
	tbot, _ := bot.CreateBot("token", &bot.ClientOpts{
		Client: http.Client{},
	})

	d := tbot.NewDispatcher()
	d.AddCallbackHandler(test1, filters.CallbackData("test"))

	d.Run()
}
