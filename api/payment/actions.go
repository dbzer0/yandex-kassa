package payment

func (p *NewPayment) Create() (*Payment, error) {
	return &Payment{
		ID:          "",
		Status:      "",
		Amount:      Amount{},
		Description: nil,
		Recipient:   Recipient{},
		Requestor:   Requestor{},
		Method:      nil,
		CreatedAt:   "",
		Test:        false,
		Paid:        false,
		Refundable:  false,
	}, nil
}

func (p *Payment) Info() error {
	return nil
}
