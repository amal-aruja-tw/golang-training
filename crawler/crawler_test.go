package crawler_test

import (
	"fmt"
	"golang-training/crawler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBillFetch(t *testing.T) {
	url := "https://mocki.io/v1/82f434ce-4124-45a2-901c-374ca88a4429"

	crwler := crawler.NewCrawler(url)

	userID := 1234
	bill, err := crwler.FetchBill(userID)

	assert.Nil(t, err)
	assert.NotNil(t, bill)
	assert.Equal(t, userID, bill.ID)

	fmt.Println(bill)
}
