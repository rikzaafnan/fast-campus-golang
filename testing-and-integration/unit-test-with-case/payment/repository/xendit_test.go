package repository_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
	"unit-test-case-cart/payment/repository"
	"unit-test-case-cart/payment/repository/mock"
)

// NEGATIVE TEST CASEs
// - sending request with empty auth header
// - sending request with invalid auth header
// - sending request with valid auth header
//   - sending request with empty request body
//   - sending request with broken request body
//   - sending request with incomplete/invalid request body
//	 = sending request with duplicate reference_id
//	 = sending request with inactive channel code : BCA
//
// POSITIVE TEST CASEs

// = sending request with complete and valid request body
// EDGE Cases
// - TODO : find the edge cases

func TestXenditPayment_SendPaymentRequest_WithEmptyAuthHeader(t *testing.T) {
	httpClient := &http.Client{}
	hostName := "https://api.xendit.co"
	xenditClient := repository.NewXenditClient(httpClient, hostName)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())
	if err != nil {
		t.Fatalf("it should not return any error, but got: %s", err.Error())
	}

	if paymentID == "" {
		t.Errorf("it should not return empty paymentID, but got: %s", paymentID)
	}
}

func TestXenditPayment_SendPaymentRequest(t *testing.T) {
	httpClientMock := mock.NewMockHttpConnector(gomock.NewController(t))
	host := "http://mock.server"
	httpClientMock.EXPECT().Do(gomock.Any()).Return(nil, errors.New("something error on xendit end"))
	xenditClient := repository.NewXenditClient(httpClientMock, host)
	paymentID, err := xenditClient.SendPaymentRequest(context.Background())

	assert.Error(t, err, "it should not return error")

	assert.Empty(t, paymentID, "it should return a valid payment ID")
}
