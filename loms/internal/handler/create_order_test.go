package handler_test

import (
	"bytes"
	"encoding/json"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/handler"
	"net/http"
	"testing"
)

var URL = "http://localhost:8080"

func TestAPI(t *testing.T) {
	server := handler.StartTestServer()
	defer func(server *http.Server) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	client := &http.Client{}

	t.Run("Test CreateOrder", func(t *testing.T) {
		createOrderReqBody, _ := json.Marshal(map[string]interface{}{
			"user": 123,
			"items": []map[string]int{
				{
					"sku":   111,
					"count": 2,
				},
				{
					"sku":   222,
					"count": 3,
				},
			},
		})

		resp, err := client.Post(URL+"/createOrder", "application/json", bytes.NewBuffer(createOrderReqBody))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK, got %v", resp.Status)
		}
	})

	t.Run("Test OrderPaid", func(t *testing.T) {
		orderPaidReqBody, _ := json.Marshal(map[string]interface{}{
			"orderID": 12345,
		})

		resp, err := client.Post(URL+"/orderPaid", "application/json", bytes.NewBuffer(orderPaidReqBody))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK, got %v", resp.Status)
		}
	})

	t.Run("Test CancelOrder", func(t *testing.T) {
		cancelOrderReqBody, _ := json.Marshal(map[string]interface{}{
			"orderID": 12345,
		})

		resp, err := client.Post(URL+"/cancelOrder", "application/json", bytes.NewBuffer(cancelOrderReqBody))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK, got %v", resp.Status)
		}
	})

	t.Run("Test ListOrder", func(t *testing.T) {
		listOrderReqBody, _ := json.Marshal(map[string]interface{}{
			"userID": 123,
		})

		resp, err := client.Post(URL+"/listOrder", "application/json", bytes.NewBuffer(listOrderReqBody))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK, got %v", resp.Status)
		}
	})

	t.Run("Test Stocks", func(t *testing.T) {
		stocksReqBody, _ := json.Marshal(map[string]interface{}{
			"SKU": 123,
		})

		resp, err := client.Post(URL+"/stocks", "application/json", bytes.NewBuffer(stocksReqBody))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status OK, got %v", resp.Status)
		}
	})
}
