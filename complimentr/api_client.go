package complimentr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://complimentr.com/
type ApiClient interface {
	GetCompliment() (*ComplimentResponse, error)
}

type apiClient struct {
	httpClient http.Client
}

func InitClient() ApiClient {
	return &apiClient{httpClient: http.Client{}}
}

func (c *apiClient) GetCompliment() (*ComplimentResponse, error) {
	response, err := c.getComplimentResponse()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *apiClient) getComplimentResponse() (*ComplimentResponse, error) {
	response, err := c.httpClient.Get("https://complimentr.com/api")
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

	var data ComplimentResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
