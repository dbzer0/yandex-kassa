package payment

type Requestor struct {
	Merchant
	ThirdParty
}

// RMerchant структура описывающая инциатора - магазин.
type Merchant struct {
	Type      string  `json:"type"`                 // значение — merchant. Тип инициатора
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
}

// RThirdParty структура описывающая инциатора - приложение.
type ThirdParty struct {
	Type       string  `json:"type"`                  // значение — merchant. Тип инициатора
	ClientID   *string `json:"client_id,omitempty"`   // идентификатор приложения
	ClientName *string `json:"client_name,omitempty"` // название приложения
}
