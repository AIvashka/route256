package model

type OrderPaidRequest struct {
	ID int64 `json:"orderID"`
}

func (r OrderPaidRequest) Validate() error {
	if r.ID == 0 {
		return ErrorOrderID
	}
	return nil
}
