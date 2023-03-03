package complimentr

import (
	"fmt"
	"log"
	"strings"
)

type Complimentr interface {
	GetCompliment() (*string, error)
}

type complimentr struct {
	client ApiClient
}

func InitComplimentr(client ApiClient) Complimentr {
	return &complimentr{client: client}
}

func (c *complimentr) GetCompliment() (*string, error) {
	response, err := c.client.GetCompliment()
	if err != nil {
		log.Println(err)

		return nil, err
	}

	compliment := fmt.Sprintf("%s! 🍎🍏❤️", strings.Title(response.Compliment))

	return &compliment, nil
}
