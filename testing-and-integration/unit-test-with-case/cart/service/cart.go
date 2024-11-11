package cart

import "context"

//go:generate mockgen -build_flags=--mod=mod -destination=mock/repository_mock.go -package=mock . RepositoryManager
type RepositoryManager interface {
	AddToCart(ctx context.Context, userID string, productID string) (err error)
}
type ShoppingCart struct {
	repo RepositoryManager
}

func NewShoppingCart(repo RepositoryManager) ShoppingCart {
	return ShoppingCart{
		repo: repo,
	}
}

func (s ShoppingCart) AddItemToCart(ctx context.Context, userID string, productID string) (err error) {
	err = s.repo.AddToCart(ctx, userID, productID)
	if err != nil {
		return err
	}

	return nil

}
