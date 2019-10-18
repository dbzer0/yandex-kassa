package payment

func (p *NewPayment) Create() (*Payment, error) {
	return &Payment{
		ID:          nil,
		Status:      nil,
		Amount:      Amount{},
		Description: nil,
		Recipient:   nil,
		Requestor:   nil,
		Method:      nil,
	}, nil
}

func (p *Payment) Info() error {
	return nil
}
