package cart_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	cart "unit-test-case-cart/cart/service"
	"unit-test-case-cart/cart/service/mock"
	"unit-test-case-cart/cart/service/stub"
)

func TestShoppingCart_AddItemToCart_ErrorOnRredis(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-1", "product-a").Return(errors.New("failing on cache system"))

	shoppingCartSvc := cart.NewShoppingCart(repositoryMock)

	err := shoppingCartSvc.AddItemToCart(context.Background(), "user-1", "product-a")
	assert.Error(t, err, "it should return an error")
}

func TestShoppingCart_AddItemToCart_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-2", "product-b").Return(nil)

	shoppingCartSvc := cart.NewShoppingCart(repositoryMock)

	err := shoppingCartSvc.AddItemToCart(context.Background(), "user-2", "product-b")
	assert.NoError(t, err, "it should not return any error")
}

func TestShoppingCart_AddItemToCart_Success_WithStub(t *testing.T) {

	repositoryStub := stub.NewRepositoryStub()
	shoppingCartSvc := cart.NewShoppingCart(repositoryStub)

	err := shoppingCartSvc.AddItemToCart(context.Background(), "user-2", "product-b")
	assert.NoError(t, err, "it should not return any error")
}
