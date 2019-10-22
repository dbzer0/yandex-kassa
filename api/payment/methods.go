package payment

import (
	"context"
	"encoding/json"

	"github.com/dbzer0/yandex-kassa/api/info"
)

// Create создает платеж. Он содержит всю необходимую информацию для проведения о
// платы (сумму, валюту и статус).
func (p *NewPayment) Create(ctx context.Context, idempKey string) (*info.Payment, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	reply, err := p.APIClient.Create(ctx, idempKey, &body)
	if err != nil {
		return nil, err
	}
	defer reply.Close()

	var payment info.Payment
	if err := json.NewDecoder(reply).Decode(&payment); err != nil {
		return nil, err
	}

	return &payment, nil
}

func (p *NewPayment) WithMethod(method string) *NewPayment {
	p.MethodData = &MethodData{
		Type: method,
	}
	return p
}

func (p *NewPayment) WithConfirmationRedirect(url string) *NewPayment {
	p.Confirmation = &Confirmation{
		Type:      "redirect",
		ReturnURL: &url,
	}
	return p
}

func (p *NewPayment) WithCapture() *NewPayment {
	capture := true
	p.Capture = &capture
	return p
}

func (p *NewPayment) WithDescription(desc string) *NewPayment {
	p.Description = &desc
	return p
}
