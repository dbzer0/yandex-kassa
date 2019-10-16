package api

import "net/http"

const apiURL = "https://payment.yandex.net/api/v3"

type Kassa struct {
	AccountID   string
	SecretKey   string
	MaxAttempts int
	client      *Client
}

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
