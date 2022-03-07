package crawler_test

import (
	"encoding/json"
	"fmt"
	"golang-training/crawler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBillFetch(t *testing.T) {
	t.Run("GiveServerURLWhenFetchIsCalledShouldReturnBill", func(t *testing.T) {
		url := "https://mocki.io/v1/82f434ce-4124-45a2-901c-374ca88a4429"

		client := crawler.NewClient(url)

		userID := 1234
		bill, err := client.FetchBill(userID)

		assert.Nil(t, err)
		assert.NotNil(t, bill)
		assert.Equal(t, userID, bill.ID)

		fmt.Println(bill)
	})

	t.Run("GivenServerRespondsWithStatusOKAndBillThenShouldReturnBillWithoutError", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bill := crawler.Bill{
				ID:     1234,
				Name:   "Amal",
				Month:  "Feb",
				Amount: 1000,
			}
			bytes, err := json.Marshal(bill)

			assert.Nil(t, err)

			w.Write(bytes)
		}))

		defer server.Close()

		client := crawler.NewClient(server.URL)

		userID := 1234
		bill, err := client.FetchBill(userID)

		assert.Nil(t, err)
		assert.NotNil(t, bill)
		assert.Equal(t, userID, bill.ID)

		fmt.Println(bill)
	})
}
