package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bill struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Month  string `json:"month"`
	Amount int    `json:"amount"`
}

type Crawler struct {
	url string
}

func NewCrawler(url string) *Crawler {
	return &Crawler{url}
}

func (c *Crawler) FetchBill(id int) (*Bill, error) {
	client := http.Client{}
	res, err := client.Get(c.url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		bill := &Bill{}
		err = json.Unmarshal(bytes, bill)
		if err != nil {
			return nil, err
		}

		return bill, nil
	}

	return &Bill{}, fmt.Errorf("unexpected response %v", res.StatusCode)
}
