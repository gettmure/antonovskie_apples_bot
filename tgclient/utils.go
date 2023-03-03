package tgclient

import (
	"fmt"
	"log"
	"net/url"
)

func getApiUrl(token string, method string, params *map[string]string) string {
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
