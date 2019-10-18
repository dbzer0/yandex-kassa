package payment

type Payment struct {
	ID          *string    `json:"id,omitempty"`             // идентификатор платежа в Яндекс.Кассе
	Status      *string    `json:"status,omitempty"`         // статус платежа. Возможные значения: pending, waiting_for_capture, succeeded и canceled
	Amount      Amount     `json:"amount"`                   // сумма платежа
	Description *string    `json:"description,omitempty"`    // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient   *Recipient `json:"recipient,omitempty"`      // получатель платежа
	Requestor   *Requestor `json:"requestor,omitempty"`      // инициатор платежа или возврата
	Method      *Method    `json:"payment_method,omitempty"` // способ оплаты, который был использован для платежа
}

type Amount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

type Recipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}
