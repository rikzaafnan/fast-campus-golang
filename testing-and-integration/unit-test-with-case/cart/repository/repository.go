package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cacher redis.Cmdable
}

func New(redis redis.Cmdable) Cache {
	return Cache{
		cacher: redis,
	}
}

type CartData struct {
	ProductID string
	Count     int
}

func (c Cache) AddToCart(ctx context.Context, userID string, productID string) error {
	_ = fmt.Sprintf("cart:%s", userID)
	_ = CartData{
		ProductID: productID,
	}

	// err := c.cacher.HSet(ctx, cacheKey, cacheValue).Err()
	// if err != nil {
	// 	return err
	// }

	return nil

}
