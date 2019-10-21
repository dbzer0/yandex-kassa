package main

import (
	"encoding/json"
	"fmt"

	"github.com/dbzer0/yandex-kassa/api"
)

func main() {
	kassa := api.New("MyAccountID", "MySecretKey")

	// формирование объекта платежа
	newPayment := kassa.NewPayment("2.00", "RUB").
		WithMethod("bank_card").
		WithConfirmationRedirect("http://example.com").
		WithDescription("test payment").
		WithCapture()

	// создание нового платежа
	p, err := newPayment.Create()
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(newPayment, "", "\t")
	fmt.Println(string(data))

	// получение информации о платеже
	p, err = p.Find()
	if err != nil {
		panic(err)
	}
	info, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(info))
}
