package service

import "context"

type PaymentGatewayProvider interface {
	SendPaymentRequest(ctx context.Context) (paymentID string, err error)
}

type Payment struct {
	xendit PaymentGatewayProvider
}

func (p *Payment) Pay() (err error) {

	//	create transaction ID
	// insert into postgres
	//	call third party API
	_, err = p.xendit.SendPaymentRequest()
	if err != nil {
		return err
	}

	return nil
}
