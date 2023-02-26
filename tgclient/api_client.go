package tgclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiClient interface {
	GetMe(token string) (*GetMeResponse, error)                       // https://core.telegram.org/bots/api#getme
	GetUpdates(token string, offset int) (*GetUpdatesResponse, error) // https://core.telegram.org/bots/api#getupdates
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
		return nil, fmt.Errorf("Failed to fetch response: code: %d, body: %s", response.StatusCode, body)
	}

	var data GetMeResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *apiClient) GetUpdates(token string, offset int) (*GetUpdatesResponse, error) {
	resp, err := c.getUpdates(token, offset)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *apiClient) getUpdates(token string, offset int) (*GetUpdatesResponse, error) {
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
		return nil, fmt.Errorf("Failed to fetch response: code: %d, body: %s", response.StatusCode, body)
	}

	var data GetUpdatesResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
