package payment

type Confirmation struct {
	Type            string  `json:"type"`                       // код сценария подтверждения
	ReturnURL       *string `json:"return_url,omitempty"`       // URL, на который вернется пользователь после подтверждения или отмены платежа
	ConfirmationURL *string `json:"confirmation_url,omitempty"` // url на который необходимо перенаправить пользователя для подтверждения оплаты
}

func (p *NewPayment) WithConfirmationRedirect(url string) *NewPayment {
	p.Confirmation = &Confirmation{
		Type:      "redirect",
		ReturnURL: &url,
	}
	return p
}
