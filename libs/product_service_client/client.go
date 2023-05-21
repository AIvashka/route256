package product_service_client

import (
	"bytes"
	"context"
	"github.com/mailru/easyjson"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/utils"
	"net/http"
)

type Client struct {
	BaseURL    string
	Token      string
	httpClient http.Client
}

func NewClient(baseURL string, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		httpClient: http.Client{},
	}
}

func (c *Client) Request(ctx context.Context, method, path string, body easyjson.Marshaler, handleResponse func(*http.Response) (interface{}, error)) (interface{}, error) {
	jsonBody, _ := easyjson.Marshal(body)

	req, err := http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(jsonBody))
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

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*GetProductResponse, error) {
	req := GetProductRequest{
		Token: c.Token,
		SKU:   sku,
	}
	response, err := c.Request(ctx, http.MethodGet, "/get_product", req, utils.DecodeInto(&GetProductResponse{}))
	if err != nil {
		return &GetProductResponse{}, nil
	}
	return response.(*GetProductResponse), nil
}

func (c *Client) ListSKUs(ctx context.Context, startAfterSku uint32, count uint32) (*ListSKUsResponse, error) {
	req := ListSKUsRequest{
		Token:         c.Token,
		StartAfterSKU: startAfterSku,
		Count:         count,
	}
	response, err := c.Request(ctx, http.MethodGet, "/list_skus", req, utils.DecodeInto(&ListSKUsResponse{}))
	if err != nil {
		return &ListSKUsResponse{}, nil
	}
	return response.(*ListSKUsResponse), nil
}
