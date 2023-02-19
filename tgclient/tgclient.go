package tgclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface {
	GetMe(ch chan interface{}) // https://core.telegram.org/bots/api#getme
}

type TelegramBot struct {
	Token string // https://core.telegram.org/bots/api#authorizing-your-bot
}

func (client TelegramBot) GetMe(ch chan interface{}) {
	go func() {
		response, err := sendRequest[GetMeResponse](client.Token, "getMe")
		if err != nil {
			log.Fatalln(fmt.Sprintf("Error fetching response for getMe method: %s", err))
		}

		ch <- response
	}()
}

func sendRequest[T Fetchable](token string, method ApiMethod) (*Response[T], error) {
	url := getApiUrl(token, method)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data Response[T]
	json.Unmarshal(body, &data)

	return &data, nil
}

func getApiUrl(token string, method ApiMethod) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method)
}
