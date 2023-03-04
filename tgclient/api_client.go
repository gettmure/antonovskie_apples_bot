package tgclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiClient interface {
	GetMe(token string) (*GetMeResponse, error)                                                    // https://core.telegram.org/bots/api#getme
	GetUpdates(token string, offset int64) (*GetUpdatesResponse, error)                            // https://core.telegram.org/bots/api#getupdates
	SendMessage(token string, text string, chatId int64) (*MessageResponse, error)                 // https://core.telegram.org/bots/api#sendmessage
	SendAudio(token string, fileId string, chatId int64, caption string) (*MessageResponse, error) // https://core.telegram.org/bots/api#sendmessage
}

type apiClient struct {
	httpClient http.Client
}

func InitClient() ApiClient {
	return &apiClient{httpClient: http.Client{}}
}

func (c *apiClient) GetMe(token string) (*GetMeResponse, error) {
	response, err := c.getMe(token)
	if err != nil {
		return nil, err
	}

	handleGetMeResponse(response)
	return response, nil
}

func (c *apiClient) GetUpdates(token string, offset int64) (*GetUpdatesResponse, error) {
	response, err := c.getUpdates(token, offset)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *apiClient) SendMessage(token string, text string, chatId int64) (*MessageResponse, error) {
	response, err := c.sendMessage(token, text, chatId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *apiClient) SendAudio(token string, audioId string, chatId int64, caption string) (*MessageResponse, error) {
	response, err := c.sendAudio(token, audioId, chatId, caption)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *apiClient) getMe(token string) (*GetMeResponse, error) {
	url := getApiUrl(token, "getMe", nil)

	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getMe failed: code: %d, body: %s", response.StatusCode, body)
	}

	var data GetMeResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *apiClient) getUpdates(token string, offset int64) (*GetUpdatesResponse, error) {
	query := map[string]string{"offset": fmt.Sprint(offset)}
	url := getApiUrl(token, "getUpdates", &query)

	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getUpdates failed: code: %d, body: %s", response.StatusCode, body)
	}

	var data GetUpdatesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *apiClient) sendMessage(token string, text string, chatId int64) (*MessageResponse, error) {
	url := getApiUrl(token, "sendMessage", nil)

	sendMessageData := &SendMessageData{ChatId: chatId, Text: text}
	request, err := json.Marshal(sendMessageData)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sendMessage failed: code: %d, body: %s", response.StatusCode, body)
	}

	var data MessageResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *apiClient) sendAudio(token string, audioId string, chatId int64, caption string) (*MessageResponse, error) {
	url := getApiUrl(token, "sendAudio", nil)

	sendAudioData := &SendAudioData{ChatId: chatId, AudioId: audioId, Caption: caption}
	request, err := json.Marshal(sendAudioData)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sendAudio failed: code: %d, body: %s", response.StatusCode, body)
	}

	var data MessageResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
