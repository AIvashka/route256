package loms_client

import (
	"bytes"
	"context"
	"github.com/mailru/easyjson"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/utils"
	"net/http"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) Request(ctx context.Context, method, path string, body easyjson.Marshaler, handleResponse func(*http.Response) (interface{}, error)) (interface{}, error) {
	jsonBody, err := easyjson.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return handleResponse(resp)
}

func (c *Client) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	response, err := c.Request(ctx, http.MethodPost, "/createOrder", req, utils.DecodeInto(&CreateOrderResponse{}))
	if err != nil {
		return &CreateOrderResponse{}, nil
	}
	return response.(*CreateOrderResponse), nil
}

func (c *Client) OrderPaid(ctx context.Context, req *OrderPaidRequest) (*OrderPaidResponse, error) {
	response, err := c.Request(ctx, http.MethodPost, "/orderPaid", req, utils.DecodeInto(&OrderPaidResponse{}))
	if err != nil {
		return &OrderPaidResponse{}, nil
	}
	return response.(*OrderPaidResponse), nil
}

func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	response, err := c.Request(ctx, http.MethodPost, "/cancelOrder", req, utils.DecodeInto(&CancelOrderResponse{}))
	if err != nil {
		return &CancelOrderResponse{}, nil
	}
	return response.(*CancelOrderResponse), nil
}

func (c *Client) ListOrder(ctx context.Context, req *ListOrderRequest) (*ListOrderResponse, error) {
	response, err := c.Request(ctx, http.MethodPost, "/listOrder", req, utils.DecodeInto(&ListOrderResponse{}))
	if err != nil {
		return &ListOrderResponse{}, nil
	}
	return response.(*ListOrderResponse), nil
}

func (c *Client) Stocks(ctx context.Context, req *StocksRequest) (*StocksResponse, error) {
	response, err := c.Request(ctx, http.MethodPost, "/stocks", req, utils.DecodeInto(&StocksResponse{}))
	if err != nil {
		return &StocksResponse{}, nil
	}
	return response.(*StocksResponse), nil

}
