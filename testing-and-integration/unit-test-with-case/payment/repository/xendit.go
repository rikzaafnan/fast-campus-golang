package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type XenditPayment struct {
	host          string
	httpConnector HttpConnector
}

func NewXenditClient(httpConnector HttpConnector, host string) XenditPayment {
	return XenditPayment{
		host:          host,
		httpConnector: httpConnector,
	}
}

type (
	XenditPaymentReqBody struct {
		Currency      string
		Amount        float64
		PaymentMethod PaymentMethod
		Metadata      Metadata
	}

	Metadata struct {
		SKU string
	}

	VirtualAccount struct {
		ChannelCode       string
		ChannelProperties struct {
			CustomerName string
		}
	}

	PaymentMethod struct {
		TypePaymentMethod string
		Reusability       string
		ReferenceID       string
		VirtualAccount    VirtualAccount
	}
)

func (x XenditPayment) SendPaymentRequest(ctx context.Context) (paymentID string, err error) {
	// TODO : inject  the xendit http client
	// call xendit paymentRequest API
	// construct request body
	// handle error response
	// handle success response

	requestBody := XenditPaymentReqBody{}
	_ = requestBody
	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(requestBody)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf("%s%s", x.host, "/payment_requests")
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, reqBody)
	if err != nil {
		return "", err
	}

	res, err := x.httpConnector.Do(httpReq)
	if err != nil {
		return "", err
	}

	rawResponseBody, err := io.ReadAll(res.Body)
	stringResponseBody := string(rawResponseBody)

	_ = stringResponseBody

	return "", nil
}
