package srvwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Wrapper[Req Validator, Res any] struct {
	fn func(ctx context.Context, req Req) (Res, error)
	//logger *zap.Logger
}

type Validator interface {
	Validate() error
}

func NewWrapper[Req Validator, Res any](fn func(ctx context.Context, req Req) (Res, error)) *Wrapper[Req, Res] {
	return &Wrapper[Req, Res]{
		fn: fn,
		//logger: logger,
	}
}

func (w *Wrapper[Req, Res]) ServeHTTP(resWriter http.ResponseWriter, httpReq *http.Request) {

	var request Req
	//w.logger.Info("got request", zap.Any("request", httpReq.Body))
	defer httpReq.Body.Close()
	err := json.NewDecoder(httpReq.Body).Decode(&request)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		writeErrorText(resWriter, "parse request", err)
		return
	}

	errValidation := request.Validate()
	if errValidation != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		writeErrorText(resWriter, "bad request", errValidation)
		return
	}

	resp, err := w.fn(httpReq.Context(), request)
	if err != nil {
		//w.logger.Error("executor fail", zap.Error(err))
		resWriter.WriteHeader(http.StatusInternalServerError)
		writeErrorText(resWriter, "exec handler", err)
		return
	}

	rawData, err := json.Marshal(&resp)
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		writeErrorText(resWriter, "decode response", err)
		return
	}

	_, _ = resWriter.Write(rawData)
}

func writeErrorText(w http.ResponseWriter, text string, err error) {
	buf := bytes.NewBufferString(text)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteByte('\n')
	w.Write(buf.Bytes())
}
