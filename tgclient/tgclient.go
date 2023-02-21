package tgclient

import (
	"log"
	"sync"
)

type Client interface {
	GetMe(ch chan any)                          // https://core.telegram.org/bots/api#getme
	GetUpdates(ch chan any, wg *sync.WaitGroup) // https://core.telegram.org/bots/api#getupdates
}

type TelegramBot struct {
	Token        string // https://core.telegram.org/bots/api#authorizing-your-bot
	LastUpdateId *int64
}

func (client TelegramBot) GetMe(ch chan any) {
	go func() {
		response, err := sendRequest[GetMeResponse](client.Token, "getMe", nil)
		if err != nil {
			log.Println(err)
		}

		ch <- response
	}()
}

func (client TelegramBot) GetUpdates(ch chan any, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		query := map[string]string{"offset": "-1"}
		response, err := sendRequest[GetUpdatesResponse](client.Token, "getUpdates", &query)
		if err != nil {
			log.Println(err)
		}

		ch <- response
	}()
}
