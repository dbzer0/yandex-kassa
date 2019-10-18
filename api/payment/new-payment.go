package payment

type NewPayment struct {
	Amount       NewAmount     `json:"amount"`                        // сумма платежа
	Description  *string       `json:"description,omitempty"`         // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient    *NewRecipient `json:"recipient,omitempty"`           // получатель платежа
	MethodData   *MethodData   `json:"payment_method_data,omitempty"` // данные для оплаты конкретным способом  (payment_method)
	Confirmation *Confirmation `json:"confirmation,omitempty"`        // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
}

func New(value, currency string) *NewPayment {
	return &NewPayment{
		Amount: NewAmount{
			Value:    value,
			Currency: currency,
		},
	}
}

type NewRecipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

type NewAmount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}
