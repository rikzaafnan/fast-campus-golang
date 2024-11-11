package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"unit-test-case-cart/payment/repository"
)

func TestXenditPayment_SendPaymentRequest(t *testing.T) {
	httpClient := &http.Client{}
	host := "http://mock.server"
	xenditClient := repository.NewXenditClient(httpClient, host)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())

	assert.NoError(t, err, "it should not return error")

	assert.NotEmpty(t, paymentID, "it should return a valid payment ID")
}
