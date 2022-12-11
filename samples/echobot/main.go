package main

import (
	"net/http"

	bot "github.com/KeralaBots/GoTGramBot"
	"github.com/KeralaBots/GoTGramBot/filters"
	"github.com/KeralaBots/GoTGramBot/types"
)

func start(b *bot.Bot, m *types.Message) error {
	replyButton := [][]types.InlineKeyboardButton{
		{
			{
				Text:         "Hi",
				CallbackData: "test",
			},
		},
	}

	_, err := b.SendMessage(
		m.Chat.Id,
		"Hi",
		&bot.SendMessageOpts{ReplyMarkup: types.InlineKeyboardMarkup{InlineKeyboard: replyButton}},
	)

	return err
}

func main() {
	tbot, _ := bot.CreateBot("token", &bot.ClientOpts{
		Client: http.Client{},
	})

	d := tbot.NewDispatcher()
	d.AddMessageHandler(start, filters.All)

	d.Run()
}
