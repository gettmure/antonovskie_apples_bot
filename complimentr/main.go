package complimentr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://complimentr.com/
type ComplimentrClient interface {
	GetCompliment() (*ComplimentResponse, error)
}

type complimentrClient struct {
	httpClient http.Client
}

func InitClient() ComplimentrClient {
	return &complimentrClient{httpClient: http.Client{}}
}

func (c *complimentrClient) GetCompliment() (*ComplimentResponse, error) {
	response, err := c.getComplimentResponse()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *complimentrClient) getComplimentResponse() (*ComplimentResponse, error) {
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
