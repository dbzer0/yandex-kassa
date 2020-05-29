package payment

import (
	"github.com/GiddeonWyeth/yandex-kassa/api/client"
)

type NewPayment struct {
	APIClient    *client.APIClient `json:"-" bson:"-"`
	Amount       Amount            `json:"amount" bson:"amount"`                                               // сумма платежа
	Description  *string           `json:"description,omitempty" bson:"description,omitempty"`                 // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient    *Recipient        `json:"recipient,omitempty" bson:"recipient,omitempty"`                     // получатель платежа
	MethodData   *MethodData       `json:"payment_method_data,omitempty" bson:"payment_method_data,omitempty"` // данные для оплаты конкретным способом  (payment_method)
	Confirmation *Confirmation     `json:"confirmation,omitempty" bson:"confirmation,omitempty"`               // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
	Capture      *bool             `json:"capture,omitempty" bson:"capture,omitempty"`                         // автоматический прием  поступившего платежа
	Receipt      *Receipt          `json:"receipt,omitempty" bson:"receipt,omitempty"`
}

type Recipient struct {
	AccountID *string `json:"account_id,omitempty" bson:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty" bson:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

type Amount struct {
	Value    string `json:"value" bson:"value"`       // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency" bson:"currency"` // код валюты в формате ISO-4217
}

type Confirmation struct {
	Type      string  `json:"type" bson:"type"`                                 // код сценария подтверждения
	ReturnURL *string `json:"return_url,omitempty" bson:"return_url,omitempty"` // URL, на который вернется пользователь после подтверждения или отмены платежа
}

type MethodData struct {
	Type string `json:"type" bson:"type"` // код способа оплаты
}

type Receipt struct {
	Customer ReceiptCustomer `json:"customer"`
	Items    [1]ReceiptItem  `json:"items"`
}

type ReceiptCustomer struct {
	Email string `json:"email"`
}

type ReceiptItem struct {
	Description    string `json:"description"`
	Quantity       string `json:"quantity"`
	Amount         Amount `json:"amount"`
	VatCode        string `json:"vat_code"`
	PaymentMode    string `json:"payment_mode"`
	PaymentSubject string `json:"payment_subject"`
}
