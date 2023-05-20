package model

type CancelOrderRequest struct {
	ID int64 `json:"orderID"`
}

func (r CancelOrderRequest) Validate() error {
	if r.ID == 0 {
		return ErrorOrderID
	}
	return nil
}
