package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/KeralaBots/GoTGramBot/filters"
	"github.com/KeralaBots/GoTGramBot/types"
)

type MessageDispatch func(b *Bot, m *types.Message) error
type CallbackDispatch func(b *Bot, m *types.CallbackQuery) error

type Dispatcher struct {
	Bot              *Bot
	IsRunning        bool
	Offset           int64
	MessageHandlers  []MessageHandlers
	CallbackHandlers []CallbackHandlers
}

type MessageHandlers struct {
	Function MessageDispatch
	Filter   filters.FilterResponse
}

type CallbackHandlers struct {
	Function CallbackDispatch
	Filter   filters.FilterResponse
}

func sigHandler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		fmt.Print("SIGTERM signal recieved. Exiting....")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		fmt.Print("SIGINT signal recieved. Exiting....")
		os.Exit(0)
	} else if signal == syscall.SIGABRT {
		fmt.Print("SIGABRT signal recieved. Exiting....")
		os.Exit(0)
	}
}

func (d *Dispatcher) Idle() {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)
	exitchnl := make(chan int)
	fmt.Println("Idling.....\nPress CTRL + C to exit")

	go func() {
		for {
			s := <-sigchnl
			sigHandler(s)
		}
	}()

	exitcode := <-exitchnl
	os.Exit(exitcode)
}

func (d *Dispatcher) Start() {
	d.IsRunning = true
	go d.handleWorkers()

}

func (d *Dispatcher) Stop() {
	d.IsRunning = false
}

func (b *Bot) NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Bot:       b,
		IsRunning: false,
		Offset:    -1,
	}
}

func (d *Dispatcher) AddMessageHandler(fn MessageDispatch, filter filters.FilterResponse) error {
	if fn != nil {
		res := MessageHandlers{
			Function: fn,
			Filter:   filter,
		}

		d.MessageHandlers = append(d.MessageHandlers, res)
		return nil
	} else {
		return fmt.Errorf("failed to add messagehandler")
	}
}

func (d *Dispatcher) AddCallbackHandler(fn CallbackDispatch, filter filters.FilterResponse) error {
	if fn != nil {
		res := CallbackHandlers{
			Function: fn,
			Filter:   filter,
		}

		d.CallbackHandlers = append(d.CallbackHandlers, res)
		return nil
	} else {
		return fmt.Errorf("failed to add callbackhandler")
	}
}

func (d *Dispatcher) Run() {
	d.Start()
	d.Idle()
	d.Stop()
}
