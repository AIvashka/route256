package product_service_client

//go:generate easyjson -all

type GetProductRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type ListSKUsRequest struct {
	Token         string `json:"token"`
	StartAfterSKU uint32 `json:"startAfterSku"`
	Count         uint32 `json:"count"`
}

type ListSKUsResponse struct {
	SKUs []uint32 `json:"skus"`
}
