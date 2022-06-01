# Yandex.Checkout API Go Client Library (Яндекс.Касса, Yandex.Checkout)

Клиент для работы с платежами по [API Yoo Money](https://yoomoney.ru/docs/wallet) (ранее [API Яндекс.Кассы](https://kassa.yandex.ru/developers/api?lang=bash))

## Пример создания платежа

```go
kassa := api.New("myShopID", "mySecretKey")

// формирование объекта платежа
newPayment := kassa.NewPayment("2.00", currency.RUB).
	WithMethodBankCard().
	WithConfirmationRedirect("http://example.com").
	WithDescription("test payment").
	WithCapture()

// генерация ключа идемпотентности
// import "github.com/google/uuid"
id, _ := uuid.NewUUID()

// создание платежа в Яндекс.Касса
p, err := newPayment.Create(context.Background(), id.String())
if err != nil {
	// обработка ошибки
}

// получение URL для подтверждения или отмены платежа пользователем
fmt.Printf("Confirmation URL: %s\n", *p.ConfirmationURL())

// получение информации о платеже
p, err = kassa.Find(context.Background(), p.ID)
if err != nil {
	// обработка ошибки
}
```
