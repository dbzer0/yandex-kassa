package api

import (
	"net/http"

	"github.com/dbzer0/yandex-kassa/api/payment"
)

const apiURL = "https://payment.yandex.net/api/v3"

type Kassa struct {
	AccountID   string
	SecretKey   string
	MaxAttempts int
	client      *Client
}

// New создает объект для работы с API Яндекс Кассы.
func New(accountID, secretKey string) *Kassa {
	return &Kassa{
		AccountID: accountID,
		SecretKey: secretKey,
		client: &Client{
			HTTP:   http.DefaultClient,
			APIURL: apiURL,
		},
	}
}

// NewHTTPClient перезаписывает http клиент, определенный по-умолчанию.
func (k *Kassa) NewHTTPClient(client *http.Client) {
	k.client.HTTP = client
}

// NewPayment создает объект NewPayment. Используется для создания платежа.
func (k *Kassa) NewPayment(value, currency string) *payment.NewPayment {
	return &payment.NewPayment{
		Amount: payment.NewAmount{
			Value:    value,
			Currency: currency,
		},
	}
}

// Payment создает создает объект Payment по которому доступны операции:
//   * получения информации о платеже;
//   * подтверждение платежа;
//   * отмена платежа;
func (k *Kassa) Payment(paymentID string) *payment.Payment {
	return &payment.Payment{
		ID: paymentID,
	}
}

// Find позволяет получить информацию о текущем состоянии платежа по
// его уникальному идентификатору.
func (k *Kassa) Find(paymentID string) (*payment.Payment, error) {
	return payment.New(paymentID).Find()
}
