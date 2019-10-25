# Yandex.Checkout API Go Client Library (Яндекс.Касса, Yandex.Checkout)

Клиент для работы с платежами по [API Яндекс.Кассы](https://kassa.yandex.ru/developers/api?lang=bash)
Подходит тем, у кого способ подключения к Яндекс.Кассе называется API.

## Пример создания платежа

```
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
np, err := newPayment.Create(context.Background(), id.String())
if err != nil {
    // обработка ошибки
}

// получение информации о платеже
p, err = kassa.Find(context.Background(), np.ID)
if err != nil {
    // обработка ошибки
}
```
