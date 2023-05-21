package loms_client

import (
	"bytes"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

func (c *Client) Request(method, path string, body easyjson.Marshaler, handleResponse func(*http.Response) (interface{}, error)) (interface{}, error) {
	jsonBody, _ := easyjson.Marshal(body)

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

func (c *Client) CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	response, err := c.Request(http.MethodPost, "/createOrder", req, decodeInto(&CreateOrderResponse{}))
	if err != nil {
		return &CreateOrderResponse{}, nil
	}
	return response.(*CreateOrderResponse), nil
}

func (c *Client) OrderPaid(req *OrderPaidRequest) (*OrderPaidResponse, error) {
	response, err := c.Request(http.MethodPost, "/orderPaid", req, decodeInto(&OrderPaidResponse{}))
	if err != nil {
		return &OrderPaidResponse{}, nil
	}
	return response.(*OrderPaidResponse), nil
}

func (c *Client) CancelOrder(req *CancelOrderRequest) (*CancelOrderResponse, error) {
	response, err := c.Request(http.MethodPost, "/cancelOrder", req, decodeInto(&CancelOrderResponse{}))
	if err != nil {
		return &CancelOrderResponse{}, nil
	}
	return response.(*CancelOrderResponse), nil
}

func (c *Client) ListOrder(req *ListOrderRequest) (*ListOrderResponse, error) {
	response, err := c.Request(http.MethodPost, "/listOrder", req, decodeInto(&ListOrderResponse{}))
	if err != nil {
		return &ListOrderResponse{}, nil
	}
	return response.(*ListOrderResponse), nil
}

func (c *Client) Stocks(req *StocksRequest) (*StocksResponse, error) {
	response, err := c.Request(http.MethodPost, "/stocks", req, decodeInto(&StocksResponse{}))
	if err != nil {
		return &StocksResponse{}, nil
	}
	return response.(*StocksResponse), nil

}

func decodeInto(v easyjson.Unmarshaler) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		err := decodeResponse(resp, v)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func decodeResponse(resp *http.Response, v easyjson.Unmarshaler) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	return easyjson.UnmarshalFromReader(resp.Body, v)
}
