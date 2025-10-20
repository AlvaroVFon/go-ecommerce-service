package carts

import "context"

type (
	Repository interface {
		FindOrCreateActiveCart(ctx context.Context, userID int64) (*Cart, error)
		UpsertItem(ctx context.Context, cartID int64, productID int64, quantity int) error
		GetItems(ctx context.Context, cartID int64) ([]CartItem, error)
		ClearCart(ctx context.Context, cartID int64) error
		SetCompleted(ctx context.Context, cartID int64) error
	}

	CartService struct {
		cartRepo Repository
	}
)

func NewCartService(cartRepo Repository) *CartService {
	return &CartService{cartRepo: cartRepo}
}

func (s *CartService) GetCart(ctx context.Context, userID int64) (*Cart, error) {
	return s.cartRepo.FindOrCreateActiveCart(ctx, userID)
}

func (s *CartService) AddItemToCart(ctx context.Context, userID int64, productID int64, quantity int) (*Cart, error) {
	cart, err := s.cartRepo.FindOrCreateActiveCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	err = s.cartRepo.UpsertItem(ctx, cart.ID, productID, quantity)
	if err != nil {
		return nil, err
	}

	items, err := s.cartRepo.GetItems(ctx, cart.ID)
	if err != nil {
		return nil, err
	}
	cart.CartItems = items
	return cart, nil
}

func (s *CartService) ClearCart(ctx context.Context, userID int64) error {
	cart, err := s.cartRepo.FindOrCreateActiveCart(ctx, userID)
	if err != nil {
		return err
	}

	return s.cartRepo.ClearCart(ctx, cart.ID)
}

func (s *CartService) CompleteCart(ctx context.Context, userID int64) error {
	cart, err := s.cartRepo.FindOrCreateActiveCart(ctx, userID)
	if err != nil {
		return err
	}

	return s.cartRepo.SetCompleted(ctx, cart.ID)
}
