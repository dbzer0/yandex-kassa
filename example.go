package main

import (
	"encoding/json"
	"fmt"

	"github.com/dbzer0/yandex-kassa/api/payment"
)

func main() {
	//kassa := api.New("", "")
	payment := payment.Payment{
		ID:          "1",
		Status:      "ok",
		Amount:      payment.Amount("2", "RUB"),
		Description: nil,
		Recipient:   payment.Recipient("123", "345"),
		Requestor:   payment.Requestor().Merchant("account123"),
		PaymentMethod: payment.Method().BankCard(&payment.Card{
			First6:      "123456",
			Last4:       "7890",
			ExpiryYear:  "2020",
			ExpiryMonth: "01",
			CardType:    "MasterCard",
		}, false),
	}

	p, _ := json.MarshalIndent(payment, "", "\t")
	fmt.Println(string(p))
}
