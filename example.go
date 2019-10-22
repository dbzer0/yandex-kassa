package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dbzer0/yandex-kassa/api"
)

func main() {
	secretKey := os.Getenv("YAK_SECRET")
	shopID := os.Getenv("YAK_SHOPID")

	kassa := api.New(shopID, secretKey)

	// формирование объекта платежа
	newPayment := kassa.NewPayment("2.00", "RUB").
		WithMethod("bank_card").
		WithConfirmationRedirect("http://example.com").
		WithDescription("test payment").
		WithCapture()

	// создание нового платежа
	ctx := context.Background()
	p, err := newPayment.Create(ctx, "uniq-idempotence-key-3")
	if err != nil {
		panic(err)
	}

	if p.Confirmation != nil && p.Confirmation.ConfirmationURL != nil {
		fmt.Printf("Confirmation URL: %s\n", *p.Confirmation.ConfirmationURL)
	}

	// получение информации о платеже
	p, err = kassa.Find(ctx, p.ID)
	if err != nil {
		panic(err)
	}
	info, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(info))
}
