package payment

type MethodData struct {
	Type string `json:"type"` // код способа оплаты
}

func (p *NewPayment) WithMethod(method string) *NewPayment {
	p.MethodData = &MethodData{
		Type: method,
	}
	return p
}
