package bot

import (
	"fmt"

	"github.com/KeralaBots/GoTGramBot/types"
)

func (d *Dispatcher) handleWorkers() error {
	for d.IsRunning {
		updates, err := d.Bot.GetUpdates(&GetUpdatesOpts{
			Offset:  d.Offset,
			Timeout: int64(TIMEOUT),
		})

		if err != nil {
			return fmt.Errorf("failed to get update : %w", err)
		}

		if len(updates) > 0 {
			//d.Offset = update[len(update)-1].UpdateId + 1
			for _, update := range updates {

				if update.Message != nil {
					handleMessageWorkers(d.MessageHandlers, d.Bot, update.Message)
				}

				if update.CallbackQuery != nil {
					handleCallbackWorkers(d.CallbackHandlers, d.Bot, update.CallbackQuery)
				}

				d.Offset = update.UpdateId + 1
			}
		}
	}

	return nil
}

func handleMessageWorkers(handlers []MessageHandlers, b *Bot, m *types.Message) {
	for _, handler := range handlers {
		check := handler.Filter.CheckMessage(m)
		if check {
			go handler.Function(b, m)
		}
	}
}

func handleCallbackWorkers(handlers []CallbackHandlers, b *Bot, m *types.CallbackQuery) {
	for _, handler := range handlers {
		check := handler.Filter.CheckCallback(m)
		if check {
			go handler.Function(b, m)
		}
	}
}
