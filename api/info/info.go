package info

import (
	"github.com/dbzer0/yandex-kassa/api/client"
)

type Payment struct {
	APIClient            *client.APIClient     `json:"-"`
	ID                   string                `json:"id"`                              // идентификатор платежа в Яндекс.Кассе
	Status               *string               `json:"status,omitempty"`                // статус платежа. Возможные значения: pending, waiting_for_capture, succeeded и canceled
	Amount               *Amount               `json:"amount,omitempty"`                // сумма платежа
	RefundedAmount       *RefundedAmount       `json:"refunded_amount,omitempty"`       // сумма, которая вернулась пользователю
	Description          *string               `json:"description,omitempty"`           // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient            *Recipient            `json:"recipient,omitempty"`             // получатель платежа
	Requestor            *Requestor            `json:"requestor,omitempty"`             // инициатор платежа или возврата
	Method               *Method               `json:"payment_method,omitempty"`        // способ оплаты, который был использован для платежа
	CreatedAt            *string               `json:"created_at,omitempty"`            // время создания заказа в формате ISO 8601. Пример: 2017-11-03T11:52:31.827Z
	CapturedAt           *string               `json:"captured_at,omitempty"`           // время подтверждения платежа
	Confirmation         *Confirmation         `json:"confirmation,omitempty"`          // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"` // данные об авторизации платежа
	Refundable           *bool                 `json:"refundable,omitempty"`            // возможность провести возврат по API
	Paid                 *bool                 `json:"paid,omitempty"`                  // признак оплаты заказа
	Test                 *bool                 `json:"test,omitempty"`                  // признак тестовой операции

	Code      *string `json:"code,omitempty"`      // содержит ID ошибки, например invalid_request
	Parameter *string `json:"parameter,omitempty"` // содержит параметр в котором произошла ошибка, например: payment_id
	Type      *string `json:"type,omitempty"`      // содержит признак ошибки, например: error
}

type Amount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

type Recipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

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

type Method struct {
	Type  string  `json:"type"`            // код способа оплаты
	ID    string  `json:"id"`              // идентификатор способа оплаты
	Saved bool    `json:"saved"`           // с помощью сохраненного способа оплаты можно проводить безакцептные списания
	Title *string `json:"title,omitempty"` // название способа оплаты
	BankCard
	YandexMoney
	ApplePay
	GooglePay
}

// BankCard описывает способ оплаты банковской картой.
type BankCard struct {
	Card *Card `json:"card,omitempty"` // данные банковской карты
}

type Card struct {
	First6        string `json:"first6"`                   // первые 6 цифр номера карты (BIN)
	Last4         string `json:"last4"`                    // последние 4 цифры номера карты
	ExpiryYear    string `json:"expiry_year"`              // срок действия, год, YYYY
	ExpiryMonth   string `json:"expiry_month"`             // срок действия, месяц, MM
	CardType      string `json:"card_type"`                // тип банковской карты: MasterCard (для карт Mastercard и Maestro), Visa (для карт Visa и Visa Electron), Mir, UnionPay, JCB, AmericanExpress, DinersClub и Unknown
	IssuerCountry string `json:"issuer_country,omitempty"` // код страны, в которой выпущена карта в формате ISO-3166 alpha-2
	IssuerName    string `json:"issuer_name,omitempty"`    // наименование банка, выпустившего карту
	Source        string `json:"source,omitempty"`         // источник данных банковской карты. Возможное значение: google_pay
}

// YandexMoney описывает способ оплаты Яндекс.Деньгами.
type YandexMoney struct {
	AccountNumber string `json:"account_number,omitempty"` // номер кошелька в Яндекс.Деньгах, из которого заплатил пользователь
}

// ApplePay описывает способ оплаты Apple Pay.
type ApplePay struct{}

// GooglePay описывает способ оплаты Google Pay.
type GooglePay struct{}

type Confirmation struct {
	Type            string  `json:"type"`                       // код сценария подтверждения
	ReturnURL       *string `json:"return_url,omitempty"`       // URL, на который вернется пользователь после подтверждения или отмены платежа
	ConfirmationURL *string `json:"confirmation_url,omitempty"` // url на который необходимо перенаправить пользователя для подтверждения оплаты
}

type AuthorizationDetails struct {
	RRN      string `json:"rrn,omitempty"`       // уникальный идентификатор транзакции в системе эмитента
	AuthCode string `json:"auth_code,omitempty"` // код авторизации банковской карты
}

type RefundedAmount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}
