package checkout

import (
	"context"
	"errors"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/repository"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/loms_client"
	ps "gitlab.ozon.dev/alexeyivashka/homework/libs/product_service_client"
)

var (
	ErrorProductOutOfStock = errors.New("product out of stock")
)

type ProductService interface {
	GetProduct(ctx context.Context, sku uint32) (*ps.GetProductResponse, error)
	ListSKUs(ctx context.Context, startAfterSKU, count uint32) (*ps.ListSKUsResponse, error)
}

type CheckoutService struct {
	lomsClient     *loms_client.Client
	repository     repository.CartRepository
	productService ProductService
}

func NewService(lomsClient *loms_client.Client,
	repository repository.CartRepository, productService ProductService) *CheckoutService {
	return &CheckoutService{
		lomsClient:     lomsClient,
		repository:     repository,
		productService: productService,
	}
}

func (s *CheckoutService) AddToCart(ctx context.Context, req *model.AddToCartRequest) (*model.AddToCartResponse, error) {
	// Check product availability via LOMS.stocks
	stocksReq := &loms_client.StocksRequest{
		SKU: req.SKU,
	}

	stocksResp, err := s.lomsClient.Stocks(ctx, stocksReq)
	if err != nil {
		return &model.AddToCartResponse{}, err
	}

	var totalNumberOfProduct uint64
	for _, stock := range stocksResp.Stocks {
		totalNumberOfProduct += stock.Count
	}

	if len(stocksResp.Stocks) == 0 || uint16(totalNumberOfProduct) < req.Count {
		return &model.AddToCartResponse{}, ErrorProductOutOfStock
	}

	err = s.repository.AddToCart(ctx, req)
	if err != nil {
		return &model.AddToCartResponse{}, err
	}

	return &model.AddToCartResponse{}, nil
}

func (s *CheckoutService) DeleteFromCart(ctx context.Context, req *model.DeleteFromCartRequest) (*model.DeleteFromCartResponse, error) {
	err := s.repository.RemoveFromCart(ctx, req.User, req.SKU, req.Count)
	if err != nil {
		return &model.DeleteFromCartResponse{}, err
	}
	return &model.DeleteFromCartResponse{}, nil
}

func (s *CheckoutService) Purchase(ctx context.Context, req *model.PurchaseRequest) (*model.PurchaseResponse, error) {
	// 1. Fetch the user's cart items
	cartItems, err := s.repository.ListCartItems(ctx, req.User)
	if err != nil {
		return nil, err
	}

	// 2. Prepare the CreateOrderRequest for LOMS
	createOrderReq := &loms_client.CreateOrderRequest{
		UserID:     int(req.User),
		OrderItems: make([]loms_client.OrderItem, len(cartItems)),
	}

	for i, item := range cartItems {
		createOrderReq.OrderItems[i] = loms_client.OrderItem{
			SKU:   int(item.SKU),
			Count: int(item.Count),
		}
	}

	// 3. Call LOMS to create the order
	orderResponse, err := s.lomsClient.CreateOrder(ctx, createOrderReq)
	if err != nil {
		return nil, err
	}

	// 4. Clear the user's cart
	err = s.repository.ClearCart(ctx, req.User)
	if err != nil {
		return nil, err
	}

	// 5. Return the PurchaseResponse with the orderID
	return &model.PurchaseResponse{
		OrderID: orderResponse.OrderID,
	}, nil
}

func (s *CheckoutService) ListCart(ctx context.Context, req *model.ListCartRequest) (*model.ListCartResponse, error) {
	cartItems, err := s.repository.ListCartItems(ctx, req.User)
	if err != nil {
		return nil, err
	}

	response := &model.ListCartResponse{}

	for _, item := range cartItems {
		product, err := s.productService.GetProduct(ctx, item.SKU)
		if err != nil {
			return nil, err
		}
		response.Items = append(response.Items, model.Item{
			SKU:   item.SKU,
			Count: item.Count,
			Name:  product.Name,
			Price: product.Price,
		})
		response.TotalPrice += uint16(product.Price) * item.Count

	}

	return response, nil

}
