package payment

import (
	"context"
	"encoding/json"

	"github.com/dbzer0/yandex-kassa/api/info"
)

const (
	methodBankCard    = "bank_card"
	methodApplePay    = "apple_pay"
	methodGooglePay   = "google_pay"
	methodYooMoney    = "yoo_money"
	methodSberBank    = "sberbank"
	methodAlfaBank    = "alfabank"
	methodTinkoffBank = "tinkoff_bank"
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

func (p *NewPayment) WithMethodBankCard() *NewPayment {
	return p.WithMethod(methodBankCard)
}

func (p *NewPayment) WithMethodApplePay() *NewPayment {
	return p.WithMethod(methodApplePay)
}

func (p *NewPayment) WithMethodGooglePay() *NewPayment {
	return p.WithMethod(methodGooglePay)
}

func (p *NewPayment) WithMethodYandexMoney() *NewPayment {
	return p.WithMethod(methodYooMoney)
}

func (p *NewPayment) WithMethodSberBank() *NewPayment {
	return p.WithMethod(methodSberBank)
}

func (p *NewPayment) WithMethodAlfaBank() *NewPayment {
	return p.WithMethod(methodAlfaBank)
}

func (p *NewPayment) WithMethodTinkoffBank() *NewPayment {
	return p.WithMethod(methodTinkoffBank)
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
