package main

import (
	"encoding/json"
	"fmt"

	"github.com/dbzer0/yandex-kassa/api/payment"
)

func main() {
	np := payment.New("2.00", "RUB").
		WithMethod("bank_card").
		WithConfirmationRedirect("http://example.com").
		WithDescription("test payment")

	p, err := np.Create()
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(np, "", "\t")
	fmt.Println(string(data))

	fmt.Println(p.Info())
}
