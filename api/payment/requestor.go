package payment

type PRequestor struct {
	RMerchant
	RThirdParty
}

// RMerchant структура описывающая инциатора - магазин.
type RMerchant struct {
	Type      string  `json:"type"`                 // значение — merchant. Тип инициатора
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
}

// RThirdParty структура описывающая инциатора - приложение.
type RThirdParty struct {
	Type       string  `json:"type"`                  // значение — merchant. Тип инициатора
	ClientID   *string `json:"client_id,omitempty"`   // идентификатор приложения
	ClientName *string `json:"client_name,omitempty"` // название приложения
}

// Requestor конструирует запрос PRequestor'а.
// Достаточно в чейне вызвать функции `Merchant()` или `ThirdParty()`.
func Requestor() *PRequestor {
	return &PRequestor{}
}

// Merchant формирует структуру инициатора платежа merchant
func (r *PRequestor) Merchant(accountID string) PRequestor {
	return PRequestor{
		RMerchant{
			Type:      "merchant",
			AccountID: &accountID,
		},
		RThirdParty{},
	}
}

// ThirdParty формирует структуру инициатора платежа third_party_client
func (r *PRequestor) ThirdParty(clientID, clientName string) PRequestor {
	return PRequestor{
		RMerchant{},
		RThirdParty{
			Type:       "merchant",
			ClientID:   &clientID,
			ClientName: &clientName,
		},
	}
}
