package payment

type Payment struct {
	ID          string           `json:"id"`                    // идентификатор платежа в Яндекс.Кассе
	Status      string           `json:"status"`                // статус платежа. Возможные значения: pending, waiting_for_capture, succeeded и canceled
	Payment     PaymentAmount    `json:"amount"`                // сумма платежа
	Description *string          `json:"description,omitempty"` // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient   PaymentRecipient `json:"recipient"`             // получатель платежа
	Requestor   PRequestor       `json:"requestor"`             // инициатор платежа или возврата
}

type PaymentAmount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

type PaymentRecipient struct {
	AccountID string `json:"account_id"` // идентификатор магазина в Яндекс.Кассе
	GatewayID string `json:"gateway_id"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}
