package main

import (
	"encoding/json"
	"fmt"

	"github.com/dbzer0/yandex-kassa/api/payment"
)

func main() {
	//kassa := api.New("", "")
	payment := payment.Payment{
		ID:     "1",
		Status: "ok",
		Amount: payment.PAmount{
			Value:    "2",
			Currency: "RUB",
		},
		Description: nil,
		Recipient: payment.PRecipient{
			AccountID: "123",
			GatewayID: "345",
		},
		//Requestor: payment.Requestor().Merchant("account123"),
		Requestor: payment.Requestor().ThirdParty("client123", "client name"),
	}

	p, _ := json.MarshalIndent(payment, "", "\t")
	fmt.Println(string(p))
}
