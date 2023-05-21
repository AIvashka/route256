package utils

import (
	"github.com/mailru/easyjson"
	"io"
	"net/http"
)

func DecodeInto(v easyjson.Unmarshaler) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		err := DecodeResponse(resp, v)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func DecodeResponse(resp *http.Response, v easyjson.Unmarshaler) error {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	return easyjson.UnmarshalFromReader(resp.Body, v)
}
