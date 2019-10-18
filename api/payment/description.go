package payment

func (p *NewPayment) WithDescription(desc string) *NewPayment {
	p.Description = &desc
	return p
}
