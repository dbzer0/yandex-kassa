package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dbzer0/yandex-kassa/api"
	"github.com/dbzer0/yandex-kassa/api/currency"
	"github.com/google/uuid"
)

func main() {
	secretKey := os.Getenv("YAK_SECRET")
	shopID := os.Getenv("YAK_SHOPID")

	kassa := api.New(shopID, secretKey)

	// формирование объекта платежа
	newPayment := kassa.NewPayment("2.00", currency.RUB).
		WithMethod("bank_card").
		WithConfirmationRedirect("http://example.com").
		WithDescription("test payment").
		WithCapture()

	// генерация ключа идемпотентности
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	// создание нового платежа
	ctx := context.Background()
	p, err := newPayment.Create(ctx, id.String())
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
