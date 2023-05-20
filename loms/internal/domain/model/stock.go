package model

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

func (r StocksRequest) Validate() error {
	if r.SKU == 0 {
		return ErrorMissingSKU
	}
	return nil
}

type Stock struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []Stock `json:"stocks"`
}
