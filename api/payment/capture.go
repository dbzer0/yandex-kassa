package payment

func (p *NewPayment) WithCapture() *NewPayment {
	capture := true
	p.Capture = &capture
	return p
}
