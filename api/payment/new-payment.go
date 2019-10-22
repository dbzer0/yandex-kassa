package payment

import (
	"github.com/dbzer0/yandex-kassa/api/client"
)

type NewPayment struct {
	APIClient    *client.APIClient `json:"-"`
	Amount       Amount            `json:"amount"`                        // сумма платежа
	Description  *string           `json:"description,omitempty"`         // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient    *Recipient        `json:"recipient,omitempty"`           // получатель платежа
	MethodData   *MethodData       `json:"payment_method_data,omitempty"` // данные для оплаты конкретным способом  (payment_method)
	Confirmation *Confirmation     `json:"confirmation,omitempty"`        // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
	Capture      *bool             `json:"capture,omitempty"`             // автоматический прием  поступившего платежа
}

type Recipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

type Amount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

type Confirmation struct {
	Type      string  `json:"type"`                 // код сценария подтверждения
	ReturnURL *string `json:"return_url,omitempty"` // URL, на который вернется пользователь после подтверждения или отмены платежа
}

type MethodData struct {
	Type string `json:"type"` // код способа оплаты
}
