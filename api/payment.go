package api

type Payment struct {
	ID      string        `json:"id"`     // идентификатор платежа в Яндекс.Кассе
	Status  string        `json:"status"` // статус платежа. Возможные значения: pending, waiting_for_capture, succeeded и canceled
	Payment PaymentAmount `json:"amount"` // сумма платежа
}

type PaymentAmount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}
