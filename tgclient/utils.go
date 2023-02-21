package tgclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func sendRequest[T Fetchable](
	token string,
	method ApiMethod,
	query *map[string]string,
) (*Response[T], error) {
	url := getApiUrl(token, method, query)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data Response[T]
	json.Unmarshal(body, &data)

	return &data, nil
}

func getApiUrl(token string, method ApiMethod, params *map[string]string) string {
	url, err := url.Parse(fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method))
	if err != nil {
		log.Println(err)
	}

	if params == nil {
		return url.String()
	}

	query := url.Query()
	for key, value := range *params {
		query.Set(key, value)
	}
	url.RawQuery = query.Encode()

	return url.String()
}
